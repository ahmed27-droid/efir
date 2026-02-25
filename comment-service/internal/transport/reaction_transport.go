package transport

import (
	"comment-service/internal/dto"
	"comment-service/internal/errs"
	"comment-service/internal/services"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReactionTransport struct {
	reactionService services.ReactionServices
}

func NewReactionTransport(reactionService services.ReactionServices) *ReactionTransport {
	return &ReactionTransport{reactionService: reactionService}
}

func (t *ReactionTransport) CreateReaction(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var req dto.CreateReactionDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid reaction body"})
		return
	}

	reaction, err := t.reactionService.CreateReaction(uint(id), req)
	if err != nil {
		switch {
		case errors.Is(err, errs.ErrPostNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		case errors.Is(err, errs.ErrBroadcastNotActive):
			c.JSON(http.StatusForbidden, gin.H{"error": "broadcast is not active"})
		case errors.Is(err, errs.ErrBroadcastService), errors.Is(err, errs.ErrUnexpectedStatusCode):
			c.JSON(http.StatusBadGateway, gin.H{"error": "broadcast service error"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create reaction"})
		}
		return
	}

	c.JSON(http.StatusCreated, reaction)
}

func (t *ReactionTransport) UpdateReaction(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req dto.UpdateReactionDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid reaction body"})
		return
	}

	reaction, err := t.reactionService.UpdateReaction(uint(id), req)
	if err != nil {
		switch {
		case errors.Is(err, errs.ErrReactionNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "reaction not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update reaction"})
		}
		return
	}

	c.JSON(http.StatusOK, reaction)
}

func (t *ReactionTransport) DeleteReaction(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := t.reactionService.DeleteReaction(uint(id)); err != nil {
		switch {
		case errors.Is(err, errs.ErrReactionNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "reaction not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete reaction"})
		}
		return
	}

	c.Status(http.StatusNoContent)
}

func (t *ReactionTransport) ListReaction(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	reactions, err := t.reactionService.ListReaction(uint(id))
	if err != nil {
		switch {
		case errors.Is(err, errs.ErrPostNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		case errors.Is(err, errs.ErrBroadcastService), errors.Is(err, errs.ErrUnexpectedStatusCode):
			c.JSON(http.StatusBadGateway, gin.H{"error": "broadcast service error"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list reaction"})
		}
		return
	}

	c.JSON(http.StatusOK, reactions)
}
