package transport

import (
	"broadcast-service/internal/kafka"
	"broadcast-service/internal/repository"
	"broadcast-service/internal/services"
	"log/slog"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(
	router *gin.Engine,
	logger *slog.Logger,
	db *gorm.DB,
	kafkaProducer *kafka.Producer,
) {
	categoryRepo := repository.NewCategoryRepo(db)
	broadcastRepo := repository.NewBroadcastRepo(db)

	categoryService := services.NewCategoryService(categoryRepo)
	broadcastService := services.NewBroadcastService(logger, broadcastRepo, categoryRepo, kafkaProducer)

	categoryHandler := NewCategoryHandler(categoryService)
	broadcastHandler := NewBroadcastHandler(logger, broadcastService)

	categoryHandler.RegisterRoutes(router)
	broadcastHandler.RegisterRoutes(router)
}
