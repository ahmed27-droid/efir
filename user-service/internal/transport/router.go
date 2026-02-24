package transport

import "github.com/gin-gonic/gin"

func RegisterRoutes(
	router *gin.Engine,
	userHandler *UserHandler,
	authHandler *AuthHandler,
){

	
		router.POST("/register", authHandler.Register)
		router.POST("/login", authHandler.Login)
		router.GET("/users/:id", userHandler.GetByID)
		router.PATCH("/users/:id", userHandler.UpdateProfile)
	}


