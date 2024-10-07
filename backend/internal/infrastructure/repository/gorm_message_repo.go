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
func (r *gormMessageRepository) CreateMessage(ctx context.Context, message *entities.Message) error {
	return r.db.Create(message).Error
}

// Получение всех сообщений по ID чата
func (r *gormMessageRepository) GetMessagesByChatID(ctx context.Context, chatID uint) ([]*entities.Message, error) {
	var messages []*entities.Message
	err := r.db.Where("chat_id = ?", chatID).Order("created_at asc").Find(&messages).Error
	return messages, err
}
