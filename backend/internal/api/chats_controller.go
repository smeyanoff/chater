package api

import (
	"chater/internal/logging"
	"chater/internal/service"
	"fmt"
	"net/http"
	"strconv"

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
// @Tags Chats, V1
// @Produce  json
// @Security BearerAuth
// @Success 200 {object} chatsResponse
// @Failure 401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /v1/chats [get]
func (c *ChatController) GetChatsForUser(ctx *gin.Context) {
	logging.Logger.Debug("Getting chats response...")

	// Получаем user_id из middleware, который проверил JWT токен
	userID, exists := ctx.Get("user_id")
	if !exists {
		logging.Logger.Error(ErrUnauthorized)
		ctx.JSON(http.StatusUnauthorized, errorResponse{Error: ErrUnauthorized})
		return
	}

	// Вызываем сервис для получения чатов
	chats, err := c.chatService.GetUserChats(ctx, userID.(uint))
	if err != nil {
		logging.Logger.Error(fmt.Sprintf("Error getting chats: %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, errorResponse{Error: "Error getting chats"})
		return
	}

	// Преобразуем данные чатов в response структуру
	response := mapChats(chats, userID.(uint))

	// Отправляем ответ
	ctx.JSON(http.StatusOK, chatsResponse{Chats: response})
	logging.Logger.Debug("Getting chats succeded")
}

// CreateChat godoc
// @Summary Создание нового чата
// @Description Создаёт новый чат с указанным именем и возвращает его данные
// @Tags Chats, V1
// @Accept  json
// @Produce  json
// @Param   chat body createChatRequest true "Данные для создания чата"
// @Success 200 {object} chatResponse "Информация о созданном чате"
// @Failure 400 {object} errorResponse "Неверный запрос"
// @Failure 500 {object} errorResponse "Ошибка на сервере"
// @Security BearerAuth
// @Router /v1/chats [post]
func (cc *ChatController) CreateChat(ctx *gin.Context) {
	var request createChatRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{Error: ErrInvalidRequest})
		return
	}

	ownerID, exists := ctx.Get("user_id") // Получаем ID пользователя (например, из JWT)
	if !exists {
		logging.Logger.Error(fmt.Sprintf("Create chat error: %s", ErrUnauthorized))
		ctx.JSON(http.StatusUnauthorized, errorResponse{Error: ErrUnauthorized})
		return
	}

	chat, err := cc.chatService.CreateChat(ctx, request.Name, ownerID.(uint))
	if err != nil {
		logging.Logger.Error(fmt.Sprintf("Create chat error: %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, errorResponse{Error: "Failed to create chat"})
		return
	}

	ctx.JSON(http.StatusOK, mapChat(chat, ownerID.(uint)))
}

// AddGroupToChat godoc
// @Summary Add a group to a chat
// @Description Adds a specified group to a chat, requires the user to be authorized
// @Tags Chats, V1
// @Param chat_id path int true "Chat ID"
// @Param group body groupAddToChatRequest true "Group to be added to chat"
// @Security BearerAuth
// @Success 200 {object} successResponse
// @Failure 400 {object} errorResponse "Invalid Chat ID or Request Format"
// @Failure 401 {object} errorResponse "Unauthorized"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /v1/chats/{chat_id}/groups [post]
func (cc *ChatController) AddGroupToChat(ctx *gin.Context) {
	chatID := ctx.Param("chat_id")
	chatIDuint, err := strconv.ParseUint(chatID, 10, 32)
	if err != nil {
		logging.Logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse{Error: "Invalid Chat ID"})
		return
	}

	var request groupAddToChatRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{Error: ErrInvalidRequest})
		return
	}

	userID, exists := ctx.Get("user_id") // Получаем ID пользователя (например, из JWT)
	if !exists {
		ctx.JSON(http.StatusUnauthorized, errorResponse{Error: ErrUnauthorized})
		return
	}

	if err := cc.chatService.AddGroupToChat(ctx, userID.(uint), uint(chatIDuint), request.GroupID); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
	}

	ctx.JSON(http.StatusOK, successResponse{Message: "Group added successfully"})
}
