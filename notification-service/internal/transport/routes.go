package transport

import "github.com/gin-gonic/gin"

func RegisterRoutes(
	router *gin.Engine,
	subTransport *SubscriptionTransport,
	notifTransport *NotificationTransport,
) {
	
	router.POST("/subscribe", subTransport.Subscribe)
	router.DELETE("/unsubscribe/:id", subTransport.Unsubscribe)

	router.GET("/notifications/:user_id", notifTransport.GetUnreadCount)
	router.PATCH("/notifications/:id/read/:user_id", notifTransport.MarkAsRead)
	router.PATCH("/notifications/read/:user_id", notifTransport.MarkAllAsRead)
}
