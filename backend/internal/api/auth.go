// internal/api/auth_handler.go
package api

import (
	"chater/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with username, email, and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body registerRequest true "User Data"
// @Success 201 {object} successResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{Error: err.Error()})
		return
	}

	err := h.authService.Register(c.Request.Context(), req.Username, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, successResponse{Message: "User registered successfully"})
}

// Login godoc
// @Summary Log in a user
// @Description Log in a user and return a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body loginRequest true "User Credentials"
// @Success 200 {object} tokenResponse
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Router /login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{Error: err.Error()})
		return
	}

	token, err := h.authService.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse{Error: "Invalid username or password"})
		return
	}

	c.JSON(http.StatusOK, tokenResponse{Token: token})
}
