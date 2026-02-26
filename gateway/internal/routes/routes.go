package routes

import (
	"gateway/internal/auth"
	"gateway/internal/middleware"
	"log"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

func Register(
	r *gin.Engine,
	jwtManager *auth.JWTManager,
	userProxy *httputil.ReverseProxy,
	broadcastProxy *httputil.ReverseProxy,
	commentProxy *httputil.ReverseProxy,
	notificationProxy *httputil.ReverseProxy,
) {
	api := r.Group("/api")
	{
		api.Any("/auth/*path", func(c *gin.Context) {
			c.Request.URL.Path = c.Param("path")
			userProxy.ServeHTTP(c.Writer, c.Request)
		})
	}

	protected := api.Group("")
	log.Println("lalal")
	protected.Use(middleware.JWTMiddleware(jwtManager))
	{
		protected.Any("/users/*path", func(c *gin.Context) { 
			c.Request.URL.Path = "/users" + c.Param("path")
			userProxy.ServeHTTP(c.Writer, c.Request)
		})

		protected.Any("/posts/*path", func(c *gin.Context) {
			c.Request.URL.Path = c.Param("path")
			broadcastProxy.ServeHTTP(c.Writer, c.Request)
		})

		protected.Any("/categories/*path", func(c *gin.Context) {
			c.Request.URL.Path = c.Param("path")
			broadcastProxy.ServeHTTP(c.Writer, c.Request)
		})

		protected.Any("/comments/*path", func(c *gin.Context) {
			c.Request.URL.Path = c.Param("path")
			commentProxy.ServeHTTP(c.Writer, c.Request)
		})

		protected.Any("/reactions/*path", func(c *gin.Context) {
			c.Request.URL.Path = c.Param("path")
			commentProxy.ServeHTTP(c.Writer, c.Request)
		})

		protected.Any("/notifications/*path", func(c *gin.Context) {
			c.Request.URL.Path = c.Param("path")
			notificationProxy.ServeHTTP(c.Writer, c.Request)
		})

		protected.Any("/subscriptions/*path", func(c *gin.Context) {
			c.Request.URL.Path = c.Param("path")
			notificationProxy.ServeHTTP(c.Writer, c.Request)
		})
	}
}
