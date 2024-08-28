package api

// registerRequest структура для запроса на регистрацию
type registerRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// loginRequest структура для запроса на вход
type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
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
