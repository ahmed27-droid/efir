package dto

type CreateBroadcastRequest struct {
	CategoryID uint   `json:"category_id" binding:"required,gt=0"`
	Name       string `json:"name" binding:"required"`
}
