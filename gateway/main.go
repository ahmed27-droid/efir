package main

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func createProxy(target string) *httputil.ReverseProxy {
	u, err := url.Parse(target)
	if err != nil {
		return nil
	}
	return httputil.NewSingleHostReverseProxy(u)
}

func main() {
	r := gin.Default()

	userProxy := createProxy("http://localhost:8081")
	broadcastProxy := createProxy("http://localhost:8082")
	commentProxy := createProxy("http://localhost:8083")
	notificationProxy := createProxy("http://localhost:8084")

	r.Any("/api/auth/*path", func(c *gin.Context) {
		c.Request.URL.Path = c.Param("path")
		userProxy.ServeHTTP(c.Writer, c.Request)
	})

	r.Any("/api/users/*path", func(c *gin.Context) {
		c.Request.URL.Path = "/users" + c.Param("path")
		userProxy.ServeHTTP(c.Writer, c.Request)
	})

	r.Any("/api/posts/*path", func(c *gin.Context) {
		broadcastProxy.ServeHTTP(c.Writer, c.Request)
	})

	r.Any("/api/categories/*path", func(c *gin.Context) {
		broadcastProxy.ServeHTTP(c.Writer, c.Request)
	})

	r.Any("/api/comments/*path", func(c *gin.Context) {
		commentProxy.ServeHTTP(c.Writer, c.Request)
	})

	r.Any("/api/reactions/*path", func(c *gin.Context) {
		commentProxy.ServeHTTP(c.Writer, c.Request)
	})

	r.Any("/api/notifications/*path", func(c *gin.Context) {
		notificationProxy.ServeHTTP(c.Writer, c.Request)
	})

	r.Any("/api/subscriptions/*path", func(c *gin.Context) {
		notificationProxy.ServeHTTP(c.Writer, c.Request)
	})

	r.Run(":8086")

}
