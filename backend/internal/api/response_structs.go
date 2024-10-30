package api

type groupResponse struct {
	ID      uint             `json:"id"`
	Name    string           `json:"name"`
	IsOwner bool             `json:"isOwner"`
	Members []memberResponse `json:"members"`
}

type messagesResponse struct {
	Messages []messageResponse `json:"messages"` // Сообщения чата
}

type chatsResponse struct {
	Chats []chatResponse `json:"chats"` // Чаты
}

type chatResponse struct {
	ID        uint              `json:"id"`        // Идентификатор чата
	Name      string            `json:"name"`      // Название чата
	CreatedAt string            `json:"createdAt"` // Время создания чата
	UpdatedAt string            `json:"updatedAt"` // Время последнего обновления чата
	Members   []memberResponse  `json:"members"`   // Список участников чата
	Messages  []messageResponse `json:"messages"`  // Последние сообщения чата
}

type memberResponse struct {
	ID       uint   `json:"id"`       // Идентификатор пользователя
	Username string `json:"username"` // Имя пользователя
}

type messageResponse struct {
	ID        uint   `json:"id"`        // Идентификатор сообщения
	SenderID  uint   `json:"senderID"`  // Идентификатор отправителя
	Sender    string `json:"sender"`    // Имя отправителя
	IsCurrent bool   `json:"isCurrent"` // Сообщение отправлено текущим пользователем
	Content   string `json:"content"`   // Содержание сообщения
	CreatedAt string `json:"createdAt"` // Время отправки сообщения
}

// successResponse структура успешного ответа
type successResponse struct {
	Message string `json:"message"`
}

// errorResponse структура ошибки
type errorResponse struct {
	Error string `json:"error"`
}
