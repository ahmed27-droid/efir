package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	JWTSecret              string
	UserServiceURL         string
	BroadcastServiceURL    string
	CommentServiceURL      string
	NotificationServiceURL string
}

func Load() *Config {

	if err := godotenv.Load(".env"); err != nil {
		log.Println(".env not found, using system environment")
	}

	cfg := &Config{
		JWTSecret:              os.Getenv("JWT_SECRET"),
		UserServiceURL:         os.Getenv("USER_SERVICE_URL"),
		BroadcastServiceURL:    os.Getenv("BROADCAST_SERVICE_URL"),
		CommentServiceURL:      os.Getenv("COMMENT_SERVICE_URL"),
		NotificationServiceURL: os.Getenv("NOTIFICATION_SERVICE_URL"),
	}

	if cfg.UserServiceURL == "" {
		log.Fatal("USER_SERVICE_URL is not set")
	}
	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	if cfg.BroadcastServiceURL == "" {
		log.Fatal("BROADCAST_SERVICE_URL is not set")
	}
	if cfg.CommentServiceURL == "" {
		log.Fatal("COMMENT_SERVICE_URL is not set")
	}
	if cfg.NotificationServiceURL == "" {
		log.Fatal("NOTIFICATION_SERVICE_URL is not set")
	}

	return cfg
}
