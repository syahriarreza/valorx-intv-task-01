package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/syahriarreza/valorx-intv-task-01/config"
	"github.com/syahriarreza/valorx-intv-task-01/internal/oauth"
	userHttp "github.com/syahriarreza/valorx-intv-task-01/internal/user/delivery/http"
	userRepo "github.com/syahriarreza/valorx-intv-task-01/internal/user/repository/postgres"
	userUsecase "github.com/syahriarreza/valorx-intv-task-01/internal/user/usecase"
)

func main() {
	cfg := config.LoadConfig()

	// Initialize OAuth configuration
	oauth.InitializeOAuthConfig()

	db, err := gorm.Open(postgres.Open(cfg.DatabaseDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	userRepository := userRepo.NewUserRepository(db)
	userService := userUsecase.NewUserUsecase(userRepository)

	router := gin.Default()
	userHttp.NewUserHandler(router, userService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
