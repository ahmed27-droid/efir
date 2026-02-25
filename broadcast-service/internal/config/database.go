package config

import (
	"broadcast-service/internal/models"
	"fmt"
	"log/slog"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnect(logger *slog.Logger) *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("BROADCAST_DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbMode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v",
		dbHost, dbUser, dbPass, dbName, dbPort, dbMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Error("failed to connect", "error", err)
		os.Exit(1)
	}

	if err := db.AutoMigrate(
		&models.Category{},
		&models.Broadcast{},
		&models.Post{},
	); err != nil {
		logger.Error("failed to migrate database", "error", err)
		os.Exit(1)
	}

	return db
}
