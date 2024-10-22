package repository

import (
	entities "chater/internal/domain/entity"
	"chater/internal/domain/repository"
	"context"

	"gorm.io/gorm"
)

type gormMessageRepository struct {
	db *gorm.DB
}

func NewGormMessageRepository(db *gorm.DB) repository.MessageRepository {
	return &gormMessageRepository{db: db}
}

// Создание нового сообщения
func (r *gormMessageRepository) Save(ctx context.Context, message *entities.Message) (*entities.Message, error) {
	// Создание сообщения
	if err := r.db.WithContext(ctx).Create(message).Error; err != nil {
		return nil, err
	}

	if err := r.db.WithContext(ctx).Preload("Sender").First(message, message.ID).Error; err != nil {
		return nil, err
	}

	return message, nil
}

// Получение всех сообщений по ID чата
func (r *gormMessageRepository) GetMessagesByChatID(ctx context.Context, chatID uint) ([]*entities.Message, error) {
	var messages []*entities.Message
	err := r.db.WithContext(ctx).Preload("Sender").Where("chat_id = ?", chatID).Order("created_at asc").Find(&messages).Error
	return messages, err
}
