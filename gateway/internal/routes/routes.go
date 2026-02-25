package routes


import (
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

func Register(
	r *gin.Engine,
	userProxy *httputil.ReverseProxy,
	broadcastProxy *httputil.ReverseProxy,
	commentProxy *httputil.ReverseProxy,
	notificationProxy *httputil.ReverseProxy,
) {


	r.Any("/api/auth/*path", func(c *gin.Context) {
		c.Request.URL.Path = c.Param("path")
		userProxy.ServeHTTP(c.Writer, c.Request)
	})
	r.Any("/api/users/*path", func(c *gin.Context) { // /api/users/*path удостоверься, правильно ли написано "*path", может не "*path", а просто "*"?
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
}