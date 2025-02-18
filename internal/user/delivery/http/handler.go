package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/syahriarreza/valorx-intv-task-01/internal/oauth"
	"github.com/syahriarreza/valorx-intv-task-01/internal/user"
	"github.com/syahriarreza/valorx-intv-task-01/pkg/models"
)

type UserHandler struct {
	UserUsecase user.Usecase
}

func NewUserHandler(router *gin.Engine, us user.Usecase) {
	handler := &UserHandler{
		UserUsecase: us,
	}

	router.POST("/users", handler.CreateUser)
	router.GET("/users/:id", handler.GetUserByID)
	router.PUT("/users/:id", handler.UpdateUser)
	router.DELETE("/users/:id", handler.DeleteUser)
	router.POST("/login", handler.Login)

	// OAuth routes
	router.GET("/auth/google", handler.GoogleLogin)
	router.GET("/callback", handler.GoogleCallback)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var request struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	user := models.User{
		ID:           uuid.New(),
		Name:         request.Name,
		Email:        request.Email,
		PasswordHash: string(hashedPassword),
	}

	if err := h.UserUsecase.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.UserUsecase.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID, _ = uuid.Parse(id)
	if err := h.UserUsecase.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := h.UserUsecase.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *UserHandler) Login(c *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UserUsecase.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GoogleLogin(c *gin.Context) {
	oauth.HandleGoogleLogin(c.Writer, c.Request)
}

func (h *UserHandler) GoogleCallback(c *gin.Context) {
	oauth.HandleGoogleCallback(c.Writer, c.Request, h.UserUsecase)
}
