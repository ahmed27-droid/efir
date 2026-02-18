package main

import (
	"commen-sService/internal/config"
	"commen-sService/internal/models"
	"log"

	"context"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func main() {

	db := config.DatabaseConnect()

	db.AutoMigrate(&models.Comment{}, &models.Reaction{})

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatal("Redis error:", err)
	}
	// broadcastClient := client.NewBroadcastClient("http://localhost:8081")
	// cache := cache.NewRedisCache(rdb)
	// commentRepo := repository.NewCommentRepository(db)
	// commentService := services.NewCommentServices(commentRepo, cache, *broadcastClient)

	router := gin.Default()

	router.Run()
}
