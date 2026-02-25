package models

import (
	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model
	UserID     uint `json:"user_id" gorm:"not null;index:idx_user_category,unique"`
	CategoryID uint `json:"category_id" gorm:"not null;index:idx_user_category,unique"`
}
