package transport

import (
	"broadcast-service/internal/services"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	logger          *slog.Logger
	categoryService *services.CategoryService
}

func NewCategoryHandler(categoryService *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}

func (h *CategoryHandler) RegisterRoutes(r *gin.Engine) {
	r.GET("/ping", h.Ping)
	r.GET("/categories", h.GetCategories)
}

func (h *CategoryHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (h *CategoryHandler) GetCategories(c *gin.Context) {

	categories, err := h.categoryService.List()
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categories})
}
