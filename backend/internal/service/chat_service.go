package service

import (
	entities "chater/internal/domain/entity"
	models "chater/internal/domain/entity"
	repo "chater/internal/domain/repository"
	"chater/internal/domain/valueobject"
	"context"
)

type ChatService struct {
	chatRepo repo.ChatRepository
	userRepo repo.UserRepository
}

func NewChatService(chatRepo repo.ChatRepository, userRepo repo.UserRepository) *ChatService {
	return &ChatService{
		chatRepo: chatRepo,
		userRepo: userRepo,
	}
}

func (cc *ChatService) CreateChat(ctx context.Context, name string, ownerID uint) (*models.Chat, error) {
	chatName, err := valueobject.NewChatName(name)
	if err != nil {
		return nil, err
	}

	owner, err := cc.userRepo.FindByID(ctx, ownerID)
	if err != nil {
		return nil, err
	}

	newChat := &models.Chat{
		Name:    chatName,
		OwnerID: ownerID,
		Members: []*entities.User{owner},
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
