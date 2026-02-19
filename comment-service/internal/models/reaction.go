package models

import "gorm.io/gorm"

type ReactionType string

const (
	Like  ReactionType = "like"
	Fire  ReactionType = "fire"
	Shock ReactionType = "shock"
	Sad   ReactionType = "sad"
	Laugh ReactionType = "laugh"
)

type Reaction struct {
	gorm.Model
	PostID uint         `json:"post_id" gorm:"not null;index"`
	UserID uint         `json:"user_id" gorm:"not null;index"`
	Type   ReactionType `json:"type" gorm:"type:varchar(10);not null"`
}
