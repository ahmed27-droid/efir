package models

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	UserID  uint   `json:"user_id" gorm:"not null;index"`
	Message string `json:"message" gorm:"type:text;not null"`
	IsRead  bool   `json:"is_read" gorm:"default:false;index"`
}
