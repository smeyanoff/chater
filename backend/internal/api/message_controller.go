package api

import (
	entities "chater/internal/domain/entity"
	"chater/internal/logging"
	"chater/internal/service"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Структура для представления клиента WebSocket
type Client struct {
	conn   *websocket.Conn
	userID uint
}

// Хранилище подключенных клиентов для каждого чата
type ChatClients struct {
	clients map[*Client]bool
	mu      sync.Mutex
}

// Обработчик для хранения клиентов в каждом чате
var chatClientsMap = make(map[uint]*ChatClients)

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

// MessageWebSocketController Обработчик WebSocket соединений для отправки и получения сообщений
// @Summary Подключение к WebSocket для чата
// @Description Подключитесь к WebSocket для получения и отправки сообщений в чате в реальном времени.
// @Tags WebSocket, Messages, V1
// @Produce json
// @Param chat_id path string true "Chat ID"
// @Success 101 {string} string "WebSocket connection established"
// @Failure 400 {object} errorResponse "Invalid Chat ID"
// @Failure 401 {object} errorResponse "Unauthorized"
// @Failure 500 {object} errorResponse "Failed to upgrade connection"
// @Router /v1/chats/{chat_id}/ws [get]
// @Security ApiKeyAuth
func (mc *MessageController) MessageWebSocketController(c *gin.Context) {

	chatID := c.Param("chat_id")
	chatIDUint, err := strconv.ParseUint(chatID, 10, 32)
	if err != nil {
		logging.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, errorResponse{Error: "Invalid Chat ID"})
		return
	}

	// Получение идентификатора пользователя
	userID, exists := c.Get("user_id")
	if !exists {
		logging.Logger.Error(ErrUnauthorized)
		c.JSON(http.StatusUnauthorized, errorResponse{Error: ErrUnauthorized})
		return
	}

	logging.Logger.Debug("Open client connection")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logging.Logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse{Error: "Failed to upgrade connection"})
		return
	}
	defer conn.Close()

	client := &Client{
		conn:   conn,
		userID: userID.(uint),
	}

	// Добавление клиента в список клиентов для этого чата
	addClientToChat(uint(chatIDUint), client)
	defer removeClientFromChat(uint(chatIDUint), client)

	for {
		var msg sendMessageRequest

		// Чтение сообщения от клиента
		err := conn.ReadJSON(&msg)
		logging.Logger.Debug(fmt.Sprintf("Read client %d message from chat %d", client.userID, chatIDUint))

		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				logging.Logger.Debug("WebSocket connection closed by client (code 1000)")
				break
			}
			logging.Logger.Error(err.Error())
			break
		}

		logging.Logger.Debug("Process message")
		// Обработка сообщения через сервисный слой
		response, err := mc.messageService.SendMessage(c.Request.Context(), uint(chatIDUint), userID.(uint), msg.Content)
		if err != nil {
			logging.Logger.Error(err.Error())
			conn.WriteJSON(errorResponse{Error: "Failed to send message"})
			break
		}

		// Отправка нового сообщения всем клиентам в чате
		broadcastMessageToChat(uint(chatIDUint), response)
	}
}

// Функция для добавления клиента в список клиентов чата
func addClientToChat(chatID uint, client *Client) {
	logging.Logger.Debug(fmt.Sprintf("Add client %d to chat %d", client.userID, chatID))
	if chatClientsMap[chatID] == nil {
		chatClientsMap[chatID] = &ChatClients{
			clients: make(map[*Client]bool),
		}
	}

	chatClientsMap[chatID].mu.Lock()
	chatClientsMap[chatID].clients[client] = true
	chatClientsMap[chatID].mu.Unlock()
}

// Функция для удаления клиента из списка клиентов чата
func removeClientFromChat(chatID uint, client *Client) {
	logging.Logger.Debug(fmt.Sprintf("Remove client %d to chat %d", client.userID, chatID))
	if chatClientsMap[chatID] != nil {
		chatClientsMap[chatID].mu.Lock()
		delete(chatClientsMap[chatID].clients, client)
		chatClientsMap[chatID].mu.Unlock()
	}
}

// Функция для рассылки сообщения всем клиентам в чате
func broadcastMessageToChat(chatID uint, response *entities.Message) {
	if chatClientsMap[chatID] != nil {
		chatClientsMap[chatID].mu.Lock()
		defer chatClientsMap[chatID].mu.Unlock()

		for client := range chatClientsMap[chatID].clients {
			logging.Logger.Debug(fmt.Sprintf("Broadcast message to client %d", client.userID))
			err := client.conn.WriteJSON(mapMessage(response, client.userID))
			if err != nil {
				logging.Logger.Error(err.Error())
				client.conn.Close()
				delete(chatClientsMap[chatID].clients, client)
			}
		}
	}
}

// GetMessages godoc
// @Summary Получение сообщений чата
// @Description Возвращает список всех сообщений в чате по его ID
// @Tags Messages, V1
// @Produce  json
// @Param  chat_id path uint true "ID чата"
// @Success 200 {object} messagesResponse "Список сообщений"
// @Failure 400 {object} errorResponse "Ошибка в запросе"
// @Failure 500 {object} errorResponse "Ошибка на стороне сервера"
// @Security BearerAuth
// @Router /v1/chats/{chat_id}/messages [get]
func (mc *MessageController) GetMessages(ctx *gin.Context) {
	logging.Logger.Debug("Get messaged response...")
	chatID := ctx.Param("chat_id")
	userID, exists := ctx.Get("user_id") // Получаем ID пользователя (например, из JWT)
	if !exists {
		logging.Logger.Error(ErrUnauthorized)
		ctx.JSON(http.StatusUnauthorized, errorResponse{Error: ErrUnauthorized})
		return
	}

	// Преобразование строки chatID в uint
	chatIDUint, err := strconv.ParseUint(chatID, 10, 32)
	if err != nil {
		logging.Logger.Error(ErrInvalidChatID)
		ctx.JSON(http.StatusBadRequest, errorResponse{Error: ErrInvalidChatID})
		return
	}

	messages, err := mc.messageService.GetMessages(ctx, uint(chatIDUint), userID.(uint))
	if err != nil {
		logging.Logger.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, errorResponse{Error: "Failed to load messages"})
		return
	}

	ctx.JSON(http.StatusOK, messagesResponse{Messages: mapMessages(messages, userID.(uint))})
	logging.Logger.Debug("Get messages succeded")
}

// GetMessages godoc
// @Summary Получение последнего сообщения чата
// @Description Возвращает последнее сообщения чата
// @Tags Messages, V1
// @Produce  json
// @Param  chat_id path uint true "ID чата"
// @Success 200 {object} messageResponse "Сообщение"
// @Failure 400 {object} errorResponse "Ошибка в запросе"
// @Failure 500 {object} errorResponse "Ошибка на стороне сервера"
// @Security BearerAuth
// @Router /v1/chats/{chat_id}/last [get]
func (mc *MessageController) GetLastMessage(ctx *gin.Context) {
	logging.Logger.Debug("Get last message response...")
	chatID := ctx.Param("chat_id")
	userID, exists := ctx.Get("user_id")
	if !exists {
		logging.Logger.Error(ErrUnauthorized)
		ctx.JSON(http.StatusUnauthorized, errorResponse{Error: ErrUnauthorized})
		return
	}

	// Преобразование строки chatID в uint
	chatIDUint, err := strconv.ParseUint(chatID, 10, 32)
	if err != nil {
		logging.Logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse{Error: ErrInvalidChatID})
		return
	}

	message, err := mc.messageService.GetLastMessageByChatID(ctx, uint(chatIDUint), userID.(uint))
	if err != nil {
		logging.Logger.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, errorResponse{Error: "Failed to load message"})
		return
	}

	ctx.JSON(http.StatusOK, mapMessage(message, userID.(uint)))
	logging.Logger.Debug("Getting last message succeded")
}
