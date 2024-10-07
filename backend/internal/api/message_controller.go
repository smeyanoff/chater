package api

import (
	"chater/internal/service"
	"net/http"
	"strconv"
	"time"

	_ "chater/internal/domain/entity"

	"github.com/gin-gonic/gin"
)

type MessageController struct {
	messageService *service.MessageService
}

func NewMessageController(messageService *service.MessageService) *MessageController {
	return &MessageController{messageService: messageService}
}

// SendMessage godoc
// @Summary Отправка сообщения в чат
// @Description Позволяет отправить сообщение в чат, указав идентификатор чата и текст сообщения
// @Tags messages
// @Accept  json
// @Produce  json
// @Param  chat_id path uint true "ID чата"
// @Param  message body sendMessageRequest true "Данные для отправки сообщения"
// @Success 200 {object} messageResponse "Успешное отправленное сообщение"
// @Failure 400 {object} errorResponse "Ошибка в запросе"
// @Failure 500 {object} errorResponse "Ошибка на стороне сервера"
// @Security BearerAuth
// @Router /chats/{chat_id}/messages [post]
func (mc *MessageController) SendMessage(ctx *gin.Context) {
	var request sendMessageRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{Error: "Invalid request"})
		return
	}

	userID := ctx.MustGet("user_id").(uint) // Получаем ID отправителя (например, из JWT)
	chatID := request.ChatID

	// Преобразование строки chatID в uint
	chatIDUint, err := strconv.ParseUint(chatID, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{Error: "Invalid chat ID"})
		return
	}

	message, err := mc.messageService.SendMessage(ctx, uint(chatIDUint), userID, request.Content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse{Error: "Failed to send message"})
		return
	}

	ctx.JSON(http.StatusOK, messageResponse{
		ID:        message.ID,
		SenderID:  message.SenderID,
		Sender:    message.Sender.Username, // Или можно получить это через отношения
		Content:   message.Content,
		CreatedAt: message.CreatedAt.Format(time.RFC3339),
	})
}

// GetMessages godoc
// @Summary Получение сообщений чата
// @Description Возвращает список всех сообщений в чате по его ID
// @Tags messages
// @Produce  json
// @Param  chat_id path uint true "ID чата"
// @Success 200 {object} messagesResponse "Список сообщений"
// @Failure 400 {object} errorResponse "Ошибка в запросе"
// @Failure 500 {object} errorResponse "Ошибка на стороне сервера"
// @Security BearerAuth
// @Router /chats/{chat_id}/messages [get]
func (mc *MessageController) GetMessages(ctx *gin.Context) {
	chatID := ctx.Param("chat_id")

	// Преобразование строки chatID в uint
	chatIDUint, err := strconv.ParseUint(chatID, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{Error: "Invalid chat ID"})
		return
	}

	messages, err := mc.messageService.GetMessages(ctx, uint(chatIDUint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse{Error: "Failed to load messages"})
		return
	}

	ctx.JSON(http.StatusOK, messagesResponse{Messages: mapMessages(messages)})
}
