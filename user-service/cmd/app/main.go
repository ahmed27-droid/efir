package main

import (
	"log"
	"user/config"
	"user/internal/auth"
	"user/internal/models"
	"user/internal/repository"
	"user/internal/services"
	"user/internal/transport"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.Load()

	db := config.SetUpDatabaseConnection()

	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("не удалось выполнить миграции: %v", err)
	}

	userRepo := repository.NewUserRepository(db)

	jwtManager := auth.NewJWTManager(cfg.JWTSecret)

	userService := services.NewUserService(userRepo, jwtManager)
	userHandler := transport.NewUserHandler(userService)
	authHandler := transport.NewAuthHandler(userService)

	r := gin.Default()

	transport.RegisterRoutes(r, userHandler, authHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
