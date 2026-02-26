package kafka

import (
	"context"
	"log"
	"notification-service/internal/services"
	"os"
	"strings"
)

type Notification struct {
	notif services.NotificationService
}

func (n Notification) Broadcast(ctx context.Context) {
	brokersEnv := os.Getenv("KAFKA_BROKERS")
	if brokersEnv == "" {
		brokersEnv = "localhost:9092"
	}
	topic := os.Getenv("KAFKA_TOPIC")
	if topic == "" {
		topic = "broadcast.started"
	}
	group := os.Getenv("KAFKA_GROUP")
	if group == "" {
		group = "comment-service-consumer"
	}

	brokers := strings.Split(brokersEnv, ",")

	consumer := NewKafkaConsumer(brokers, topic, group)
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Println("failed to close kafka consumer:", err)
		}
	}()
	ev, err := consumer.ConsumeBroadcast(ctx)
	if err != nil {
		log.Println("kafka consumer stopped with error:", err)
		return
	}

	if err := n.notif.NotifyBroadcastStarted(ev); err != nil {
		log.Println("failed to notify broadcast started:", err)
	}
}

func (n Notification) Posts(ctx context.Context) {
	brokersEnv := os.Getenv("KAFKA_BROKERS")
	if brokersEnv == "" {
		brokersEnv = "localhost:9092"
	}
	topic := os.Getenv("KAFKA_TOPIC")
	if topic == "" {
		topic = "broadcast.started"
	}
	group := os.Getenv("KAFKA_GROUP")
	if group == "" {
		group = "comment-service-consumer"
	}

	brokers := strings.Split(brokersEnv, ",")

	consumer := NewKafkaConsumer(brokers, topic, group)
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Println("failed to close kafka consumer:", err)
		}
	}()
	ev, err := consumer.ConsumePost(ctx)
	if err != nil {
		log.Println("kafka consumer stopped with error:", err)
		return
	}

	if err := n.notif.NotifyPost(ev); err != nil {
		log.Println("failed to notify broadcast started:", err)
	}
}
