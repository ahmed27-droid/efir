package main

import (
	"comment-service/internal/cache"
	"comment-service/internal/client"
	"comment-service/internal/config"
	"comment-service/internal/kafka"
	"comment-service/internal/models"
	"comment-service/internal/repository"
	"comment-service/internal/services"
	"comment-service/internal/transport"
	"context"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	db := config.DatabaseConnect()

	if err := db.AutoMigrate(&models.Comment{}, &models.Reaction{}); err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	rdb := config.NewRedis()

	broadcastURL := os.Getenv("BROADCAST_SERVICE_URL")
	if broadcastURL == "" {
		broadcastURL = "http://localhost:8081"
	}
	broadcastClient := client.NewBroadcastClient(broadcastURL)
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

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		kafka.RunWorker(ctx, cache)
	}()

	defer cancel()

	if err := router.Run(":8080"); err != nil {
		panic("failed to run server: " + err.Error())
	}

}
