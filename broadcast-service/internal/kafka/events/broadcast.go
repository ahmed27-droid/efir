package kafka

type BroadcastStartedEvent struct {
	BroadcastID uint64 `json:"broadcast_id"`
	Title       string `json:"title"`
	Category    string `json:"category"`
}
