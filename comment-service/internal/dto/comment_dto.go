package dto

type CreateCommentDTO struct {
	UserID  uint   `json:"user_id" binding:"required,gt=0"`
	Content string `json:"content" binding:"required,min=5,max=1000"`
}

type UpdateCommentDTO struct {
	Content *string `json:"content" binding:"omitempty,min=5,max=1000"`
}
