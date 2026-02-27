package kafka

import (
	"comment-service/internal/cache"
	"comment-service/internal/client"
	"context"
	"log"
	"os"
	"strings"
)

func RunWorker(ctx context.Context, bc cache.BroadcastCache) {
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
