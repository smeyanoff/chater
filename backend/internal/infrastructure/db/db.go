package db

import (
	"chater/internal/config"
	models "chater/internal/domain/entity"
	_ "chater/internal/domain/valueobject"
	"chater/internal/logging"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB устанавливает соединение с базой данных PostgreSQL
func ConnectDB(cfg *config.Config) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	maxRetries := 10
	retryDelay := time.Duration(36000)

	dsn := fmt.Sprintf("host=%s user=%s password=%s database=%s port=%d sslmode=disable",
		cfg.DB.Host, cfg.DB.User, cfg.DB.Pwd, cfg.DB.Table, cfg.DB.Port)

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Connected to the database successfully")
			return db, nil
		}

		logging.Logger.Info(
			fmt.Sprintf("Failed to connect to the database. Attempt %d/%d. Retrying in %s...",
				i+1, maxRetries, retryDelay))
		time.Sleep(retryDelay) // Задержка перед повторной попыткой
	}
	return nil, fmt.Errorf("could not connect to the database after %d attempts: %w", maxRetries, err)
}

// AutoMigrate выполняет автоматическую миграцию всех моделей
func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
		&models.Chat{},
		&models.Message{},
		&models.Group{},
	)
	if err != nil {
		return err
	}

	logging.Logger.Info("Database migration completed successfully")
	return nil
}
