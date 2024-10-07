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
	ChatID  string `json:"chat_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type createChatRequest struct {
	Name string `json:"name" binding:"required"` // Название чата
}
