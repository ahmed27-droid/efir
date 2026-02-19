package main

import (
	"comment-Service/internal/config"
	"comment-Service/internal/models"

	"github.com/gin-gonic/gin"
)

func main() {

	db := config.DatabaseConnect()

	db.AutoMigrate(&models.Comment{}, &models.Reaction{})

	//rdb := config.NewRedis()

	// broadcastClient := client.NewBroadcastClient("http://localhost:8081")
	// cache := cache.NewRedisCache(rdb)
	// commentRepo := repository.NewCommentRepository(db)
	// commentService := services.NewCommentServices(commentRepo, cache, *broadcastClient)

	router := gin.Default()

	router.Run()
}
