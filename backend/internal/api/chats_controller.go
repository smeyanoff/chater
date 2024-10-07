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
		ctx.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
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

// CreateChat godoc
// @Summary Создание нового чата
// @Description Создаёт новый чат с указанным именем и возвращает его данные
// @Tags chats
// @Accept  json
// @Produce  json
// @Param   chat body createChatRequest true "Данные для создания чата"
// @Success 200 {object} chatResponse "Информация о созданном чате"
// @Failure 400 {object} errorResponse "Неверный запрос"
// @Failure 500 {object} errorResponse "Ошибка на сервере"
// @Security BearerAuth
// @Router /chats [post]
func (cc *ChatController) CreateChat(ctx *gin.Context) {
	var request createChatRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{Error: "Invalid request"})
		return
	}

	ownerID := ctx.MustGet("user_id").(uint) // Получаем ID пользователя (например, из JWT)

	chat, err := cc.chatService.CreateChat(ctx, request.Name, ownerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse{Error: "Failed to create chat"})
		return
	}

	ctx.JSON(http.StatusOK, chatResponse{
		ID:        chat.ID,
		Name:      chat.Name.String(),
		CreatedAt: chat.CreatedAt.String(),
		UpdatedAt: chat.UpdatedAt.String(),
	})
}
