package main

import (
	"gateway/internal/auth"
	"gateway/internal/config"
	"gateway/internal/proxy"
	"gateway/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.Load()

	r := gin.Default()

	userProxy, _ := proxy.CreateProxy(cfg.UserServiceURL)
	broadcastProxy, _ := proxy.CreateProxy(cfg.BroadcastServiceURL)
	commentProxy, _ := proxy.CreateProxy(cfg.CommentServiceURL)
	notificationProxy, _ := proxy.CreateProxy(cfg.NotificationServiceURL)

	jwtManager := auth.NewJWTManager(cfg.JWTSecret)

	// simple health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	routes.Register(r, jwtManager, userProxy, broadcastProxy, commentProxy, notificationProxy)

	r.Run(":8092")
}
