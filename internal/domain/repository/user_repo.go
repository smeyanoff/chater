package repository

import (
	"chater/internal/domain/models"
	"context"
)

type UserRepository interface {
	Save(ctx context.Context, user *models.User) error
	FindByID(ctx context.Context, id uint) (*models.User, error)
	FindByUsername(ctx context.Context, username string) (*models.User, error)
	Delete(ctx context.Context, id uint) error
}
