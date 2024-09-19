package api

import (
	entities "chater/internal/domain/entity"
	"chater/internal/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ChatController struct {
	chatService *service.ChatService
}

// Конструктор для ChatController
func NewChatController(chatService *service.ChatService) *ChatController {
	return &ChatController{
		chatService: chatService,
	}
}

// GetChatsForUser godoc
// @Summary Get all chats for the authenticated user
// @Description Returns a list of all chats that the authenticated user participates in, including chat members and recent messages.
// @Tags chats
// @Produce  json
// @Security BearerAuth
// @Success 200 {object} chatsResponse
// @Failure 401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /chats [get]
func (c *ChatController) GetChatsForUser(ctx *gin.Context) {
	// Получаем user_id из middleware, который проверил JWT токен
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, errorResponse{Error: "Unauthorized"})
		return
	}

	// Вызываем сервис для получения чатов
	chats, err := c.chatService.GetUserChats(ctx, userID.(uint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse{Error: "Failed to fetch chats"})
		return
	}

	// Преобразуем данные чатов в response структуру
	response := mapChatsToResponse(chats)

	// Отправляем ответ
	ctx.JSON(http.StatusOK, chatsResponse{Chats: response})
}

// Преобразование сущностей чатов в структуру ответа
func mapChatsToResponse(chats []*entities.Chat) []chatResponse {
	var response []chatResponse
	for _, chat := range chats {
		chatResponse := chatResponse{
			ID:        chat.ID,
			Name:      chat.Name.String(),
			CreatedAt: chat.CreatedAt.Format(time.RFC3339),
			UpdatedAt: chat.UpdatedAt.Format(time.RFC3339),
			Members:   mapMembers(chat.Members),
			Messages:  mapMessages(chat.Messages),
		}
		response = append(response, chatResponse)
	}
	return response
}

// Преобразование участников чатов в структуру ответа
func mapMembers(members []*entities.User) []chatMember {
	var result []chatMember
	for _, member := range members {
		result = append(result, chatMember{
			ID:       member.ID,
			Username: member.Username,
		})
	}
	return result
}

// Преобразование сообщений чатов в структуру ответа
func mapMessages(messages []*entities.Message) []messageDetail {
	var result []messageDetail
	for _, message := range messages {
		result = append(result, messageDetail{
			ID:        message.ID,
			SenderID:  message.SenderID,
			Sender:    message.Sender.Username, // Или можно получить это через отношения
			Content:   message.Content,
			CreatedAt: message.CreatedAt.Format(time.RFC3339),
		})
	}
	return result
}
