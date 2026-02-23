package main

import (
	"comment-Service/internal/cache"
	"comment-Service/internal/client"
	"comment-Service/internal/config"
	"comment-Service/internal/models"
	"comment-Service/internal/repository"
	"comment-Service/internal/services"
	"comment-Service/internal/transport"

	"github.com/gin-gonic/gin"
)

func main() {

	db := config.DatabaseConnect()

	db.AutoMigrate(&models.Comment{}, &models.Reaction{})

	rdb := config.NewRedis()

	broadcastClient := client.NewBroadcastClient("http://localhost:8081")
	cache := cache.NewRedisCache(rdb)

	commentRepo := repository.NewCommentRepository(db)
	commentService := services.NewCommentServices(commentRepo, cache, *broadcastClient)
	commentTransport := transport.NewCommentTransport(commentService)

	reactionRepo := repository.NewReactionRepository(db)
	reactionService := services.NewReactionService(reactionRepo, cache, *broadcastClient)
	reactionTransport := transport.NewReactionTransport(reactionService)

	router := gin.Default()

	transport.RegisterRoutes(
		router,
		commentTransport,
		reactionTransport,
	)

	router.Run()
}
