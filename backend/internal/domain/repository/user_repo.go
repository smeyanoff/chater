package repository

import (
	models "chater/internal/domain/entity"
	"context"
)

type UserRepository interface {
	Save(ctx context.Context, user *models.User) error
	FindUserByID(ctx context.Context, userID uint) (*models.User, error)
	FindByUsername(ctx context.Context, username string) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	Delete(ctx context.Context, userID uint) error
}
