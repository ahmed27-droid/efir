package config

import (
	"log"
	"os"
)

type Config struct {
	JWTSecret          string
	PostServiceURL     string
	CategoryServiceURL string
	ReactionServiceURL string
}

func Load() *Config {
	cfg := &Config{
		JWTSecret:          os.Getenv("JWT_SECRET"),
		PostServiceURL:     os.Getenv("POST_SERVICE_URL"),
		CategoryServiceURL: os.Getenv("CATEGORY_SERVICE_URL"),
		ReactionServiceURL: os.Getenv("REACTION_SERVICE_URL"),
	}

	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET is not set")
	}
	return cfg
}
