package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	JWTSecret string
}

func Load() *Config {
	err := godotenv.Load("/Users/dinislambakkaev/efir/user-service/.env")
	if err != nil {
		log.Println(".env not found, using system env")
	}

	cfg := &Config{
		JWTSecret: os.Getenv("JWT_SECRET"),
	}

	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	return cfg
}