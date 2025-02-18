package oauth

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/syahriarreza/valorx-intv-task-01/internal/user"
	"github.com/syahriarreza/valorx-intv-task-01/pkg/models"
)

var (
	googleOauthConfig *oauth2.Config
	oauthStateString  = "random" // Replace with a secure random string
)

// InitializeOAuthConfig initializes the OAuth configuration
func InitializeOAuthConfig() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  viper.GetString("OAUTH_REDIRECT_URL"),
		ClientID:     viper.GetString("OAUTH_CLIENT_ID"),
		ClientSecret: viper.GetString("OAUTH_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleGoogleCallback(w http.ResponseWriter, r *http.Request, userUsecase user.Usecase) {
	if r.FormValue("state") != oauthStateString {
		http.Error(w, "State is not valid", http.StatusBadRequest)
		return
	}

	code := r.FormValue("code")
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("Could not get token: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	client := googleOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.Printf("Could not create request: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	defer resp.Body.Close()

	var userInfo struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		log.Printf("Could not parse response: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// Check if user already exists
	existingUser, err := userUsecase.GetUserByEmail(userInfo.Email)
	if err != nil {
		// If user does not exist, create a new user
		newUser := models.User{
			ID:    uuid.New(),
			Email: userInfo.Email,
			Name:  "Google User", // You might want to get the name from Google as well
		}
		if err := userUsecase.CreateUser(&newUser); err != nil {
			log.Printf("Could not create user: %s\n", err.Error())
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		existingUser = &newUser
	}

	fmt.Println("existingUser:", existingUser)

	// Redirect to the home page or wherever you want
	http.Redirect(w, r, fmt.Sprintf("/users/%s", existingUser.ID.String()), http.StatusTemporaryRedirect)
}
