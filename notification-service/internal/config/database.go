package config

import (
	"fmt"
	"os"

	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnect() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("No .env file found, relying on environment variables")
	}
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)

	for {
		db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
		if err == nil {
			return db
		}
		fmt.Println("database connection failed, retrying:", err)
		time.Sleep(2 * time.Second)
	}
}
