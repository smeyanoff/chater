package repository

import (
	models "chater/internal/domain/entity"
	"context"
)

type MessageRepository interface {
	Save(ctx context.Context, message *models.Message) (*models.Message, error)
	GetLastMessage(ctx context.Context, chatID uint) (*models.Message, error)
	GetMessages(ctx context.Context, chatID uint) ([]*models.Message, error)
}
