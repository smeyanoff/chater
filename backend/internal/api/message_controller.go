package api

import (
	"chater/internal/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type MessageController struct {
	messageService *service.MessageService
}

func NewMessageController(messageService *service.MessageService) *MessageController {
	return &MessageController{messageService: messageService}
}

// SendMessageWebSocket - обработчик WebSocket соединений для отправки и получения сообщений
func (mc *MessageController) SendMessageWebSocket(c *gin.Context) {

	chatID := c.Param("chat_id")

	// Преобразование строки chatID в uint
	chatIDUint, err := strconv.ParseUint(chatID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{Error: "invalid chat_id"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: "failed to upgrade connection"})
		return
	}
	defer conn.Close()

	for {
		// Получаем идентификатор текущего пользователя
		userID, exists := c.Get("user_id")
		if !exists {
			conn.WriteJSON(errorResponse{Error: "Unauthorized"})
			return
		}

		var msg sendMessageRequest
		// Чтение сообщений от клиента
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Ошибка чтения сообщения:", err)
			break
		}

		// Обработка сообщения через сервисный слой
		response, err := mc.messageService.SendMessage(c.Request.Context(), uint(chatIDUint), userID.(uint), string(msg.Content))
		if err != nil {
			conn.WriteJSON(errorResponse{Error: "Failed to send message"})
			break
		}

		// Отправка ответа обратно клиенту
		if err := conn.WriteJSON(response); err != nil {
			log.Println("Ошибка отправки сообщения:", err)
			break
		}
	}
}

// GetMessages godoc
// @Summary Получение сообщений чата
// @Description Возвращает список всех сообщений в чате по его ID
// @Tags messages, api, v1
// @Produce  json
// @Param  chat_id path uint true "ID чата"
// @Success 200 {object} messagesResponse "Список сообщений"
// @Failure 400 {object} errorResponse "Ошибка в запросе"
// @Failure 500 {object} errorResponse "Ошибка на стороне сервера"
// @Security BearerAuth
// @Router /api/v1/chats/{chat_id}/messages [get]
func (mc *MessageController) GetMessages(ctx *gin.Context) {
	chatID := ctx.Param("chat_id")
	userID := ctx.MustGet("user_id").(uint)

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

	ctx.JSON(http.StatusOK, messagesResponse{Messages: mapMessages(messages, userID)})
}
