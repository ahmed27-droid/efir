package main

import (
	"gateway/internal/config"
	"gateway/internal/proxy"
	"gateway/internal/routes"
	"gateway/internal/auth"

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

	routes.Register(r, jwtManager, userProxy, broadcastProxy, commentProxy, notificationProxy)

	r.Run(":8092")
}
