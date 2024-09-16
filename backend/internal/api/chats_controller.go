package api

import (
	"chater/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ChatController struct {
	chatService *service.ChatService
}

func NewChatController(chatService *service.ChatService) *ChatController {
	return &ChatController{chatService: chatService}
}

// GetChatsByUserIDHandler обрабатывает запрос на получение чатов по ID пользователя
func (cc *ChatController) GetChatsByUserIDHandler(c *gin.Context) {
	userIDParam := c.Param("user_id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Получаем чаты, связанные с этим пользователем через сервис
	chats, err := cc.chatService.GetUserChats(c.Request.Context(), uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve chats"})
		return
	}

	// Возвращаем список чатов
	c.JSON(http.StatusOK, chats)
}
