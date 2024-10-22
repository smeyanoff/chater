package repository

import (
	models "chater/internal/domain/entity"
	_ "chater/internal/domain/valueobject"
	"context"
)

type ChatRepository interface {
	Save(ctx context.Context, chat *models.Chat) error
	FindAllUserChatsWithLastMessage(ctx context.Context, userID uint) ([]*models.Chat, error)
	AddChatMember(ctx context.Context, chat *models.Chat, memberToAdd *models.User) error
	RemoveChatMember(ctx context.Context, chat *models.Chat, memberToRemove *models.User) error
	Delete(ctx context.Context, ownerID uint, chatID uint) error
}
