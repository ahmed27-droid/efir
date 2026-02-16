package models

import "gorm.io/gorm"

type ReactionType string

const (
	ReactionLike  ReactionType = "like"
	ReactionFire  ReactionType = "fire"
	ReactionShock ReactionType = "shock"
	ReactionSad   ReactionType = "sad"
	ReactionLaugh ReactionType = "laugh"
)

type Reaction struct {
	gorm.Model
	PostID uint         `json:"post_id" gorm:"not null;index"`
	UserID uint         `json:"user_id" gorm:"not null;index"`
	Type   ReactionType `json:"type" gorm:"type:varchar(10);not null"`
}
