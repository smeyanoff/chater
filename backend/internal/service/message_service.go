package service

import (
	models "chater/internal/domain/entity"
	"chater/internal/domain/repository"
	"chater/internal/logging"
	"context"
	"errors"
)

const PermissionDeniedError = "permission denied"

type MessageService struct {
	messageRepo repository.MessageRepository
	chatRepo    repository.ChatRepository
	groupRepo   repository.GroupRepository
}

func NewMessageService(messageRepo repository.MessageRepository, chatRepo repository.ChatRepository, groupRepo repository.GroupRepository) *MessageService {
	return &MessageService{messageRepo: messageRepo, groupRepo: groupRepo, chatRepo: chatRepo}
}

func (s *MessageService) checkRights(ctx context.Context, chatID uint, senderID uint) (bool, error) {
	chat, err := s.chatRepo.GetChatByID(ctx, chatID)
	if err != nil {
		logging.Logger.Error(err.Error())
		return false, err
	}

	// Проверка, является ли пользователь участником чата
	for _, user := range chat.ChatUsers {
		if user.ID == senderID {
			return true, nil
		}
	}

	if chat.ChatGroups != nil {
		// Проверка, состоит ли пользователь в группе чата
		for _, group := range chat.ChatGroups {
			member, err := s.groupRepo.CheckUserIsGroupMember(ctx, group.ID, senderID)
			if err != nil {
				logging.Logger.Error(err.Error())
				return false, err
			}
			if member {
				return true, nil
			}
		}
	}
	return false, nil
}

// Создание нового сообщения
func (s *MessageService) SendMessage(ctx context.Context, chatID uint, senderID uint, content string) (*models.Message, error) {
	if hasRights, err := s.checkRights(ctx, chatID, senderID); err != nil {
		logging.Logger.Error(err.Error())
		return nil, err
	} else if !hasRights {
		return nil, errors.New(PermissionDeniedError)
	}
	message := &models.Message{
		Content:  content,
		ChatID:   chatID,
		SenderID: senderID,
	}
	message, err := s.messageRepo.Save(ctx, message)

	// Создаем сообщение в базе данных
	if err != nil {
		logging.Logger.Error(err.Error())
		return nil, err
	}

	return message, nil
}

// Получение всех сообщений чата
func (s *MessageService) GetMessages(ctx context.Context, chatID uint, userID uint) ([]*models.Message, error) {
	if hasRights, err := s.checkRights(ctx, chatID, userID); err != nil {
		logging.Logger.Error(err.Error())
		return nil, err
	} else if !hasRights {
		return nil, errors.New(PermissionDeniedError)
	}
	return s.messageRepo.GetMessagesByChatID(ctx, chatID)
}

func (s *MessageService) GetLastMessageByChatID(ctx context.Context, chatID uint, userID uint) (*models.Message, error) {
	if hasRights, err := s.checkRights(ctx, chatID, userID); err != nil {
		logging.Logger.Error(err.Error())
		return nil, err
	} else if !hasRights {
		return nil, errors.New(PermissionDeniedError)
	}
	return s.messageRepo.GetLastMessageByChatID(ctx, chatID)
}
