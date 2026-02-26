package main

import (
	"notification-service/internal/config"
	"notification-service/internal/models"
	"notification-service/internal/repository"
	"notification-service/internal/services"
	"notification-service/internal/transport"

	"github.com/gin-gonic/gin"
)

func main() {

	db := config.DatabaseConnect()

	if err := db.AutoMigrate(&models.Notification{}, &models.Subscription{}); err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	subRepo := repository.NewSubscriptionRepository(db)
	notifRepo := repository.NewNotificationRepository(db)

	notifService := services.NewNotificationService(notifRepo, subRepo)
	subService := services.NewSubscriptionService(subRepo)

	notifTransport := transport.NewNotificationTransport(notifService)
	subTransport := transport.NewSubscriptionTransport(subService)

	router := gin.Default()

	transport.RegisterRoutes(
		router,
		subTransport,
		notifTransport,
	)

	if err := router.Run(":8080"); err != nil {
		panic("failed to run server: " + err.Error())
	}

}
