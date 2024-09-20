// internal/domain/service/auth_service.go
package service

import (
	"chater/internal/domain/auth"
	models "chater/internal/domain/entity"
	"chater/internal/domain/repository"
	"chater/internal/domain/validation"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo       repository.UserRepository
	jwtSecret      string
	experationTime int
}

func NewAuthService(userRepo repository.UserRepository, jwtSecret string, experationTime int) *AuthService {
	return &AuthService{
		userRepo:       userRepo,
		jwtSecret:      jwtSecret,
		experationTime: experationTime,
	}
}

func (s *AuthService) generateToken(userID uint) (string, error) {
	// Определяем claims
	claims := auth.Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(s.experationTime))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Создаем токен с claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен секретным ключом
	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *AuthService) Register(ctx context.Context, username, email, password string) error {

	if err := validation.ValidateEmail(email); err != nil {
		return err
	}

	existedUser, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return err
	}

	if existedUser != nil {
		return errors.New("user email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	return s.userRepo.Save(ctx, user)
}

func (s *AuthService) Login(ctx context.Context, username, password string) (string, error) {
	user, err := s.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	tokenString, err := s.generateToken(user.ID)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
