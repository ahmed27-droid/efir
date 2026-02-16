package main

import (
	"commen-sService/internal/config"
	"commen-sService/internal/models"

	"github.com/gin-gonic/gin"
)

func main() {

	db := config.DatabaseConnect()

	db.AutoMigrate(&models.Comment{}, &models.Reaction{})


	router := gin.Default()

	router.Run()
}