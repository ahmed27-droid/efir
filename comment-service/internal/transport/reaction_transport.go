package transport

import (
	"comment-Service/internal/dto"
	"comment-Service/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReactionTransport struct {
	reactionService services.ReactionServices
}

func NewReactionTransport(reactionService services.ReactionServices) *ReactionTransport{
	return &ReactionTransport{reactionService: reactionService}
}

func (t *ReactionTransport) CreateReaction(c *gin.Context){
	var req dto.CreateReactionDTO

	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid reaction body"})
		return
	}

	reaction, err := t.reactionService.CreateReaction(req)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"failed to create reaction"})
		return
	}

	c.JSON(http.StatusCreated, reaction)
}

func (t *ReactionTransport) UpdateReaction(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid id"})
		return
	}

	var req dto.UpdateReactionDTO

	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid reaction body"})
		return
	}

	reaction, err := t.reactionService.UpdateReaction(uint(id), req)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"failed to update reaction"})
		return
	}

	c.JSON(http.StatusOK, reaction)
}

func (t *ReactionTransport) DeleteReaction(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid id"})
		return
	}

	if err := t.reactionService.DeleteReaction(uint(id)); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"filed to delete reaction"})
		return
	}

	c.Status(http.StatusNoContent)
}

func (t *ReactionTransport) ListReaction(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid id"})
		return
	}

	reactions, err := t.reactionService.ListReaction(uint(id))
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"failed to list reaction"})
	}

	c.JSON(http.StatusOK, reactions)
}