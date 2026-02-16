package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	PostID uint `json:"post_id" gorm:"not null;index"`
	UserID uint  `json:"user_id" gorm:"not null;index"`
	Content string `json:"content" gorm:"type:text;not null"`
}