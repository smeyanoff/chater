package repository

import (
	models "chater/internal/domain/entity"
	_ "chater/internal/domain/valueobject"
	"context"
)

type ChatRepository interface {
	Save(ctx context.Context, chat *models.Chat) error
	GetChatByID(ctx context.Context, chatID uint) (*models.Chat, error)
	GetUserChats(ctx context.Context, userID uint) ([]*models.Chat, error)
	AddChatUser(ctx context.Context, chat *models.Chat, userToAdd *models.User) error
	RemoveChatUser(ctx context.Context, chat *models.Chat, userToRemove *models.User) error
	Delete(ctx context.Context, chatID uint) error
	AddGroup(ctx context.Context, chat *models.Chat, group *models.Group) error
	RemoveGroup(ctx context.Context, chat *models.Chat, group *models.Group) error
}
