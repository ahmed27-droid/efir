package dto

type CreateCommentDTO struct {
	PostID   uint   `json:"post_id" binding:"required,gt=0"`
	AuthorID uint   `json:"author_id" binding:"required,gt=0"`
	Content  string `json:"content" binding:"required,min=5,max=1000"`
}

type UpdateCommentDTO struct {
	Content *string `json:"content" binding:"omitempty,min=5,max=1000"`
}
