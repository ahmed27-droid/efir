package transport

import "github.com/gin-gonic/gin"

func RegisterRoutes(
	router *gin.Engine,
	commentTransport *CommentTransport,
	reactionTransport *ReactionTransport,
) {
	router.GET("/comments", commentTransport.ListComments)
	router.POST("/posts/:id/comments", commentTransport.CreateComment)
	router.PATCH("/comments/:id", commentTransport.UpdateComment)
	router.DELETE("/comments/:id", commentTransport.DeleteComment)

	router.GET("/reactions", reactionTransport.ListReaction)
	router.POST("posts/:id/reactions", reactionTransport.CreateReaction)
	router.PATCH("/reactions/:id", reactionTransport.UpdateReaction)
	router.DELETE("/reactions/:id", reactionTransport.DeleteReaction)
}
