package dto

type CreateReactionDTO struct {
	UserID uint   `json:"user_id" binding:"required,gt=0"`
	Type   string `json:"type" binding:"required,oneof=like fire shock sad laugh"`
}

type UpdateReactionDTO struct {
	Type *string `json:"type" binding:"omitempty,oneof=like fire shock sad laugh"`
}
