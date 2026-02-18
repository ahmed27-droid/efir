package main

import (
	"user/config"
	"user/internal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DatabaseConnect()

	err := db.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.Run(":8080")
}
