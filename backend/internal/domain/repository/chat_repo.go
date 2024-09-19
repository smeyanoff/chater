package repository

import (
	models "chater/internal/domain/entity"
	_ "chater/internal/domain/valueobject"
	"context"
)

type ChatRepository interface {
	Save(ctx context.Context, user *models.Chat) error
	FindByID(ctx context.Context, id uint) (*models.Chat, error)
	FindAll(ctx context.Context) ([]*models.Chat, error)
	FindAllByUserIdWithLastMessage(ctx context.Context, userId uint) ([]*models.Chat, error)
	Delete(ctx context.Context, id uint) error
}
