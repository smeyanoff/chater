package db

import (
	"chater/internal/config"
	"chater/internal/domain/models"
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
	retryDelay := time.Duration(3)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Database.Host, cfg.Database.User, cfg.Database.Pwd, cfg.Database.DB, cfg.Database.Port)

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Connected to the database successfully")
			return db, nil
		}

		log.Printf("Failed to connect to the database. Attempt %d/%d. Retrying in %s...", i+1, maxRetries, retryDelay)
		time.Sleep(retryDelay) // Задержка перед повторной попыткой
	}
	return nil, fmt.Errorf("could not connect to the database after %d attempts: %w", maxRetries, err)
}

// AutoMigrate выполняет автоматическую миграцию всех моделей
func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{}, // добавьте здесь все модели, которые нужно мигрировать
		// &models.AnotherModel{}, // Пример: если есть другие модели, добавьте их сюда
	)
	if err != nil {
		return err
	}

	log.Println("Database migration completed successfully")
	return nil
}
