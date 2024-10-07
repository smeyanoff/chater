package api

import (
	entities "chater/internal/domain/entity"
	"time"
)

// Преобразование участников чатов в структуру ответа
func mapMembers(members []*entities.User) []chatMemberResponse {
	var result []chatMemberResponse
	for _, member := range members {
		result = append(result, chatMemberResponse{
			ID:       member.ID,
			Username: member.Username,
		})
	}
	return result
}

// Преобразование сообщений чатов в структуру ответа
func mapMessages(messages []*entities.Message) []messageResponse {
	var result []messageResponse
	for _, message := range messages {
		result = append(result, messageResponse{
			ID:        message.ID,
			SenderID:  message.SenderID,
			Sender:    message.Sender.Username, // Или можно получить это через отношения
			Content:   message.Content,
			CreatedAt: message.CreatedAt.Format(time.RFC3339),
		})
	}
	return result
}
