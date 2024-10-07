package repository

import (
	models "chater/internal/domain/entity"
	"context"
)

type MessageRepository interface {
	CreateMessage(ctx context.Context, message *models.Message) error
	GetMessagesByChatID(ctx context.Context, chatID uint) ([]*models.Message, error)
}
