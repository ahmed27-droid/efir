package transport

import (
	"errors"
	"net/http"
	"notification-service/errs"
	"notification-service/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NotificationTransport struct {
	services services.NotificationService
}

func NewNotificationTransport(service services.NotificationService) *NotificationTransport {
	return &NotificationTransport{services: service}
}

func (t *NotificationTransport) MarkAsRead(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	userid, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	if err := t.services.MarkAsRead(uint(id), uint(userid)); err != nil {
		switch {
		case errors.Is(err, errs.ErrNotificationNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "notification not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to mark notification as read"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "notification marked as read"})

}

func (t *NotificationTransport) MarkAllAsRead(c *gin.Context) {
	userid, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	if err := t.services.MarkAllAsRead(uint(userid)); err != nil {
		switch {
		case errors.Is(err, errs.ErrNotificationNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "no notifications found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to mark notifications as read"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "all notifications marked as read"})
}

func (t *NotificationTransport) GetUnreadCount(c *gin.Context) {
	userid, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	count, err := t.services.GetUnreadCount(uint(userid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get unread count"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"unread_count": count})
}
