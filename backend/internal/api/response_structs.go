package api

type messagesResponse struct {
	Messages []messageResponse `json:"messages"` // Сообщения чата
}

type chatsResponse struct {
	Chats []chatResponse `json:"chats"` // Чаты
}

type chatResponse struct {
	ID        uint                 `json:"id"`         // Идентификатор чата
	Name      string               `json:"name"`       // Название чата
	CreatedAt string               `json:"created_at"` // Время создания чата
	UpdatedAt string               `json:"updated_at"` // Время последнего обновления чата
	Members   []chatMemberResponse `json:"members"`    // Список участников чата
	Messages  []messageResponse    `json:"messages"`   // Последние сообщения чата
}

type chatMemberResponse struct {
	ID       uint   `json:"id"`       // Идентификатор пользователя
	Username string `json:"username"` // Имя пользователя
}

type messageResponse struct {
	ID        uint   `json:"id"`         // Идентификатор сообщения
	SenderID  uint   `json:"sender_id"`  // Идентификатор отправителя
	Sender    string `json:"sender"`     // Имя отправителя
	Content   string `json:"content"`    // Содержание сообщения
	CreatedAt string `json:"created_at"` // Время отправки сообщения
}

// successResponse структура успешного ответа
type successResponse struct {
	Message string `json:"message"`
}

// tokenResponse структура ответа с JWT-токеном
type tokenResponse struct {
	Token string `json:"token"`
}

// errorResponse структура ошибки
type errorResponse struct {
	Error string `json:"error"`
}
