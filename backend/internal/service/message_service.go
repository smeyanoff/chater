package service

import (
	models "chater/internal/domain/entity"
	"chater/internal/domain/repository"
	"context"
)

type MessageService struct {
	messageRepo repository.MessageRepository
}

func NewMessageService(messageRepo repository.MessageRepository) *MessageService {
	return &MessageService{messageRepo: messageRepo}
}

// Создание нового сообщения
func (s *MessageService) SendMessage(ctx context.Context, chatID uint, senderID uint, content string) (*models.Message, error) {
	message := &models.Message{
		Content:  content,
		ChatID:   chatID,
		SenderID: senderID,
	}
	message, err := s.messageRepo.Save(ctx, message)

	// Создаем сообщение в базе данных
	if err != nil {
		return nil, err
	}

	return message, nil
}

// Получение всех сообщений чата
func (s *MessageService) GetMessages(ctx context.Context, chatID uint) ([]*models.Message, error) {

	return s.messageRepo.GetMessagesByChatID(ctx, chatID)
}
