package transport

import (
	"errors"
	"net/http"
	"notification-service/errs"
	"notification-service/internal/dto"
	"notification-service/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SubscriptionTransport struct {
	subService services.SubscriptionService
}

func NewSubscriptionTransport(subService services.SubscriptionService) *SubscriptionTransport {
	return &SubscriptionTransport{subService: subService}
}

func (t *SubscriptionTransport) Subscribe(c *gin.Context) {
	var req dto.CreateSubscriptionDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	if err := t.subService.Subscribe(req); err != nil {
		switch {
		case errors.Is(err, errs.ErrSubscriptionExists):
			c.JSON(http.StatusConflict, gin.H{"error": "subscription already exists"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create subscription"})
		}
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "subscribed successfully"})
}

func (t *SubscriptionTransport) Unsubscribe(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid id"})
		return
	}

	if err:= t.subService.Unsubscribe(uint(id)); err != nil{
		switch{
			case errors.Is(err, errs.ErrSubscriptionNotFound):
				c.JSON(http.StatusNotFound, gin.H{"error": "subscription not found"})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete subscription"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "unsubscribed successfully"})
}
