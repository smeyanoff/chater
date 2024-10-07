package repository

import (
	models "chater/internal/domain/entity"
	"chater/internal/domain/repository"
	_ "chater/internal/domain/valueobject"
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

func (r *gormChatRepository) FindAllByUserIdWithLastMessage(ctx context.Context, userId uint) ([]*models.Chat, error) {
	var chats []*models.Chat

	// Загрузить чаты пользователя с последним сообщением для каждого чата
	err := r.db.
		Joins("JOIN chat_users ON chat_users.chat_id = chats.id").
		Where("chat_users.user_id = ?", userId).
		Preload("Members"). // Загрузка участников чатов
		Preload("Messages", func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at DESC").Limit(1) // Загрузка только последнего сообщения
		}).
		Find(&chats).Error

	if err != nil {
		return nil, err
	}

	return chats, nil
}

func (r *gormChatRepository) Delete(ctx context.Context, id uint) error {
	err := r.db.WithContext(ctx).Delete(&models.Chat{}, id).Error
	return err
}
