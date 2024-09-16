package repository

import (
	models "chater/internal/domain/entity"
	"chater/internal/domain/repository"
	"context"

	"gorm.io/gorm"
)

type gormChatRepository struct {
	db *gorm.DB
}

func NewGormChatRepository(db *gorm.DB) repository.ChatRepository {
	return &gormChatRepository{db: db}
}

func (r *gormChatRepository) Save(ctx context.Context, chat *models.Chat) error {
	return r.db.WithContext(ctx).Create(chat).Error
}

func (r *gormChatRepository) FindByID(ctx context.Context, id uint) (*models.Chat, error) {
	var chat models.Chat
	err := r.db.WithContext(ctx).Preload("Messages").First(&chat, id).Error
	return &chat, err
}

func (r *gormChatRepository) FindAll(ctx context.Context) ([]*models.Chat, error) {
	var chats []*models.Chat
	err := r.db.WithContext(ctx).Preload("Messages").Find(&chats).Error
	return chats, err
}

func (r *gormChatRepository) FindAllByUserId(ctx context.Context, userId uint) ([]*models.Chat, error) {
	var user models.User
	err := r.db.Preload("Chats").First(&user, userId).Error
	if err != nil {
		return nil, err
	}

	return user.Chats, nil
}

func (r *gormChatRepository) Delete(ctx context.Context, id uint) error {
	err := r.db.WithContext(ctx).Delete(&models.Chat{}, id).Error
	return err
}
