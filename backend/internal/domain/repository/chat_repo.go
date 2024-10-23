package repository

import (
	models "chater/internal/domain/entity"
	_ "chater/internal/domain/valueobject"
	"context"
)

type ChatRepository interface {
	Save(ctx context.Context, chat *models.Chat) error
	FindAllUserChatsWithLastMessage(ctx context.Context, userID uint) ([]*models.Chat, error)
	AddChatUser(ctx context.Context, chat *models.Chat, userToAdd *models.User) error
	RemoveChatUser(ctx context.Context, chat *models.Chat, userToRemove *models.User) error
	Delete(ctx context.Context, chatID uint) error
}
