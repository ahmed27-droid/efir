package transport

import (
	"comment-Service/internal/dto"
	"comment-Service/internal/errs"
	"comment-Service/internal/services"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentTransport struct {
	comService services.CommentServices
}

func NewCommentTransport(comService services.CommentServices) *CommentTransport {
	return &CommentTransport{comService: comService}
}

func (t *CommentTransport) CreateComment(c *gin.Context) {
	var req dto.CreateCommentDTO
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	comment, err := t.comService.CreateComment(uint(id), &req)
	if err != nil {
		switch {
		case errors.Is(err, errs.ErrPostNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		case errors.Is(err, errs.ErrBroadcastNotActive):
			c.JSON(http.StatusForbidden, gin.H{"error": "broadcast is not active"})
		case errors.Is(err, errs.ErrBroadcastService), errors.Is(err, errs.ErrUnexpectedStatusCode):
			c.JSON(http.StatusBadGateway, gin.H{"error": "broadcast service error"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create comment"})
		}
		return
	}

	c.JSON(http.StatusCreated, comment)
}

func (t *CommentTransport) UpdateComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var req dto.UpdateCommentDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	comment, err := t.comService.UpdateComment(uint(id), &req)
	if err != nil {
		switch {
		case errors.Is(err, errs.ErrCommentNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "comment not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update comment"})
		}
		return
	}

	c.JSON(http.StatusOK, comment)
}

func (t *CommentTransport) DeleteComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := t.comService.DeleteComment(uint(id)); err != nil {
		switch {
		case errors.Is(err, errs.ErrCommentNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "comment not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete comment"})
		}
		return
	}

	c.Status(http.StatusNoContent)
}

func (t *CommentTransport) ListComments(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit"})
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page"})
		return
	}

	comments, err := t.comService.ListComments(uint(id), page, limit)
	if err != nil {
		switch {
		case errors.Is(err, errs.ErrPostNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		case errors.Is(err, errs.ErrBroadcastService), errors.Is(err, errs.ErrUnexpectedStatusCode):
			c.JSON(http.StatusBadGateway, gin.H{"error": "broadcast service error"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list comments"})
		}
		return
	}

	c.JSON(http.StatusOK, comments)
}
