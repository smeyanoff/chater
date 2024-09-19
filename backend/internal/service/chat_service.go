package service

import (
	models "chater/internal/domain/entity"
	repo "chater/internal/domain/repository"
	"chater/internal/domain/valueobject"
	"context"
)

type ChatService struct {
	chatRepo repo.ChatRepository
}

func NewChatService(chatRepo repo.ChatRepository) *ChatService {
	return &ChatService{chatRepo: chatRepo}
}

func (cc *ChatService) CreateChat(ctx context.Context, name string) (*models.Chat, error) {
	chatName, err := valueobject.NewChatName(name)
	if err != nil {
		return nil, err
	}

	newChat := &models.Chat{
		Name: chatName,
	}
	if err := cc.chatRepo.Save(ctx, newChat); err != nil {
		return nil, err
	}
	return newChat, nil
}

func (cc *ChatService) GetUserChats(ctx context.Context, userId uint) ([]*models.Chat, error) {
	models, err := cc.chatRepo.FindAllByUserIdWithLastMessage(ctx, userId)
	if err != nil {
		return nil, err
	}
	return models, nil
}
