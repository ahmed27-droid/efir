package main

import (
	"log"
	"user/config"
	"user/internal/auth"
	"user/internal/models"
	"user/internal/repository"
	"user/internal/services"

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

	r := gin.Default()

	
}
