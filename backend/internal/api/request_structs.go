package api

// registerRequest структура для запроса на регистрацию
type registerRequest struct {
	Username string `json:"username" binding:"required"` // Имя пользователя
	Email    string `json:"email" binding:"required"`    // Емейл пользователя
	Password string `json:"password" binding:"required"` // Пароль пользователя
}

// loginRequest структура для запроса на вход
type loginRequest struct {
	Username string `json:"username" binding:"required"` // Имя пользователя
	Password string `json:"password" binding:"required"` // Пароль пользователя
}

type sendMessageRequest struct {
	Content string `json:"content" binding:"required"` // Текст сообщения
}

type createChatRequest struct {
	Name    string `json:"name" binding:"required"` // Название чата
	GroupID uint   `json:"groupID"`
}

type createGroupRequest struct {
	Name string `json:"name" binding:"required"` // Название группы
}

type userAddToGroupRequest struct {
	UserID uint `json:"userID" binding:"required"` // ID пользователя
}

type groupAddToChatRequest struct {
	GroupID uint `json:"groupID" binding:"required"` // ID пользователя
}
