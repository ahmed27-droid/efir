package config

import (
	"log"
	"os"
)

type Config struct {
	JWTSecret string
}

func Load() *Config {

	cfg := &Config{
		JWTSecret: os.Getenv("JWT_SECRET"),
	}

	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	return cfg
}
