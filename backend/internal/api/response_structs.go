package api

type groupsResponse struct {
	Groups []groupResponse `json:"groups"`
}

type groupResponse struct {
	ID      uint   `json:"id"`      // ID группы
	Name    string `json:"name"`    // Имя группы
	IsOwner bool   `json:"isOwner"` // Является ли пользователь владельцем группы
}

type messagesResponse struct {
	Messages []messageResponse `json:"messages"` // Сообщения чата
}

type chatsResponse struct {
	Chats []chatResponse `json:"chats"` // Чаты
}

type chatResponse struct {
	ID      uint             `json:"id"`      // Идентификатор чата
	Name    string           `json:"name"`    // Название чата
	Members []memberResponse `json:"members"` // Список участников чата
	Groups  []groupResponse  `json:"groups"`  // Список групп
}

type memberResponse struct {
	ID       uint   `json:"id"`       // Идентификатор пользователя
	Username string `json:"username"` // Имя пользователя
}

type messageResponse struct {
	ID        uint   `json:"id"`        // Идентификатор сообщения
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
