package kafka

import (
	kafka "broadcast-service/internal/kafka/events"
	"context"
	"encoding/json"

	kafka_go "github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka_go.Writer
}

func NewProducer(brokers []string) *Producer {
	return &Producer{
		writer: &kafka_go.Writer{
			Addr:     kafka_go.TCP(brokers...),
			Balancer: &kafka_go.LeastBytes{},
		},
	}
}

func (p *Producer) BroadcastStarted(
	ctx context.Context,
	event kafka.BroadcastStartedEvent,
) error {
	return p.writer.WriteMessages(ctx, kafka_go.Message{
		Topic: BroadcastStarted,
		Value: mustJSON(event),
	})
}

func (p *Producer) Close() error {
	return p.writer.Close()
}

func mustJSON(v any) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}
