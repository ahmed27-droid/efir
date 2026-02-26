package dto

import "time"

type BroadcastStartedEvent struct {
	BroadcastID uint       `json:"broadcast_id"`
	CategoryID  uint       `json:"category_id"`
	Title       string     `json:"title"`
	StartedAt   *time.Time `json:"started_at"`
}

type PostCreatedEvent struct {
	PostID      uint   `json:"post_id"`
	Title       string `json:"title"`
	CategoryID  uint   `json:"category_id"`
	BroadcastID uint   `json:"broadcast_id"`
	Importance  string `json:"importance"` // "breaking", "normal"
}
