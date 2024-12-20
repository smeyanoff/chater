package api

import (
	entities "chater/internal/domain/entity"
	"time"
)

// Преобразование участников чатов в структуру ответа
func mapMembers(members []*entities.User) []memberResponse {
	var result []memberResponse
	for _, member := range members {
		result = append(result, mapMember(member))
	}
	return result
}

func mapMember(member *entities.User) memberResponse {
	return memberResponse{
		ID:       member.ID,
		Username: member.Username,
	}
}

// Преобразование сообщений чатов в структуру ответа
func mapMessages(messages []*entities.Message, userID uint) []messageResponse {
	var result []messageResponse
	for _, message := range messages {
		result = append(result, mapMessage(message, userID))
	}
	return result
}

func mapMessage(message *entities.Message, userID uint) messageResponse {
	isCurrentUser := userID == message.SenderID

	return messageResponse{
		ID:        message.ID,
		Sender:    message.Sender.Username,
		IsCurrent: isCurrentUser,
		Content:   message.Content,
		CreatedAt: message.CreatedAt.Format(time.RFC3339),
	}
}

// Преобразование сущностей чатов в структуру ответа
func mapChats(chats []*entities.Chat, userID uint) []chatResponse {
	var response []chatResponse
	for _, chat := range chats {
		response = append(response, mapChat(chat, userID))
	}
	return response
}

func mapChat(chat *entities.Chat, userID uint) chatResponse {
	return chatResponse{
		ID:      chat.ID,
		Name:    chat.Name.String(),
		Groups:  mapGroups(chat.ChatGroups, userID),
		Members: mapMembers(chat.ChatUsers),
	}
}

func mapGroups(groups []*entities.Group, userID uint) []groupResponse {
	var response []groupResponse
	for _, group := range groups {
		response = append(response, mapGroup(group, userID))
	}
	return response
}

func mapGroup(group *entities.Group, userID uint) groupResponse {
	isOwner := userID == group.OwnerID
	return groupResponse{
		ID:      group.ID,
		Name:    group.Name.String(),
		IsOwner: isOwner,
	}
}
