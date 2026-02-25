package dto

type CreateSubscriptionDTO struct {
	UserID     uint   `json:"user_id"`
	CategoryID uint `json:"category_id"`
}
