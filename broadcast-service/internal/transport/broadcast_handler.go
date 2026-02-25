package transport

import (
	"broadcast-service/internal/dto"
	"broadcast-service/internal/services"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BroadcastHandler struct {
	logger           *slog.Logger
	broadcastService *services.BroadcastService
}

func NewBroadcastHandler(
	logger *slog.Logger,
	broadcastService *services.BroadcastService,
) *BroadcastHandler {
	return &BroadcastHandler{
		logger:           logger,
		broadcastService: broadcastService,
	}
}

func (h *BroadcastHandler) RegisterRoutes(r *gin.Engine) {
	b := r.Group("/broadcast")
	{
		b.POST("/:id/start", h.Start)
		b.POST("/", h.Create)
	}
}

func (h *BroadcastHandler) Create(c *gin.Context) {
	var broadcastDto dto.CreateBroadcastRequest
	if err := c.ShouldBindJSON(&broadcastDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	broadcast, err := h.broadcastService.Create(&broadcastDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": broadcast})
}

func (h *BroadcastHandler) Start(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.broadcastService.Start(ctx, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
