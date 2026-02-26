package kafka

import (
	"context"
	"encoding/json"
	"log"
	"notification-service/internal/dto"
	"strings"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewKafkaConsumer(brokers []string, topic, groupID string) *Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 1,
		MaxBytes: 10e6,
	})
	return &Consumer{reader: r}
}

func (kc *Consumer) ConsumeBroadcast(ctx context.Context) (dto.BroadcastStartedEvent, error) {
	for {
		m, err := kc.reader.ReadMessage(ctx)
		if err != nil {
			if err == context.Canceled || strings.Contains(err.Error(), "context canceled") {
				return dto.BroadcastStartedEvent{}, nil
			}
			log.Println("kafka read error:", err)
			continue
		}
		var ev dto.BroadcastStartedEvent
		if err := json.Unmarshal(m.Value, &ev); err != nil {
			log.Println("failed to unmarshal kafka message:", err, "value:", string(m.Value))
			continue
		}
		return ev, nil
	}
}

func (kc *Consumer) ConsumePost(ctx context.Context) (dto.PostCreatedEvent, error) {
	for {
		m, err := kc.reader.ReadMessage(ctx)
		if err != nil {
			if err == context.Canceled || strings.Contains(err.Error(), "context canceled") {
				return dto.PostCreatedEvent{}, nil
			}
			log.Println("kafka read error:", err)
			continue
		}
		var ev dto.PostCreatedEvent
		if err := json.Unmarshal(m.Value, &ev); err != nil {
			log.Println("failed to unmarshal kafka message:", err, "value:", string(m.Value))
			continue
		}
		return ev, nil
	}
}

func (kc *Consumer) Close() error {
	return kc.reader.Close()
}
