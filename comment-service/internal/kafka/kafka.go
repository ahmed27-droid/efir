package kafka

import (
	"context"
	"log"
	"os"
	"strings"

	"comment-Service/internal/cache"
	"comment-Service/internal/client"
)

func RunWorker(ctx context.Context, bc cache.BroadcastCache) {
	brokersEnv := os.Getenv("KAFKA_BROKERS")
	if brokersEnv == "" {
		brokersEnv = "localhost:9092"
	}
	topic := os.Getenv("KAFKA_TOPIC")
	if topic == "" {
		topic = "broadcast-status"
	}
	group := os.Getenv("KAFKA_GROUP")
	if group == "" {
		group = "comment-service-consumer"
	}

	brokers := strings.Split(brokersEnv, ",")

	consumer := client.NewKafkaConsumer(brokers, topic, group)
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Println("failed to close kafka consumer:", err)
		}
	}()

	if err := consumer.Consume(ctx, bc); err != nil {
		log.Println("kafka consumer stopped with error:", err)
	} else {
		log.Println("kafka consumer stopped")
	}
}
