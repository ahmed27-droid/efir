package client

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"comment-Service/internal/cache"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	reader *kafka.Reader
}

func NewKafkaConsumer(brokers []string, topic, groupID string) *KafkaConsumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 1,
		MaxBytes: 10e6,
	})
	return &KafkaConsumer{reader: r}
}

type broadcastEvent struct {
	ID       uint `json:"id"`
	IsActive bool `json:"isActive"`
}

func (kc *KafkaConsumer) Consume(ctx context.Context, bc cache.BroadcastCache) error {
	for {
		m, err := kc.reader.ReadMessage(ctx)
		if err != nil {
			if err == context.Canceled || strings.Contains(err.Error(), "context canceled") {
				return nil
			}
			log.Println("kafka read error:", err)
			continue
		}

		var ev broadcastEvent
		if err := json.Unmarshal(m.Value, &ev); err != nil {
			log.Println("failed to unmarshal kafka message:", err, "value:", string(m.Value))
			continue
		}

		if ev.IsActive {
			bc.SetActive(ev.ID)
		} else {
			bc.SetInActive(ev.ID)
		}
	}
}

func (kc *KafkaConsumer) Close() error {
	return kc.reader.Close()
}
