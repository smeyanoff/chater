package repository

import (
	models "chater/internal/domain/entity"
	"chater/internal/domain/repository"
	_ "chater/internal/domain/valueobject"
	"context"
	"errors"

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

func (r *gormChatRepository) FindChatByID(ctx context.Context, chatID uint) (*models.Chat, error) {
	var chat models.Chat
	if err := r.db.WithContext(ctx).
		Preload("ChatUsers").
		Preload("Messages").
		First(&chat, chatID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("chat not found")
		} else {
			return nil, err
		}
	}

	return &chat, nil
}

func (r *gormChatRepository) FindAllChatsWithLastMessage(ctx context.Context, userID uint) ([]*models.Chat, error) {
	var chats []*models.Chat

	// Загрузить чаты пользователя с последним сообщением для каждого чата
	err := r.db.WithContext(ctx).
		Joins("JOIN chat_users ON chat_users.chat_id = chats.id").
		Preload("ChatUsers").
		Preload("ChatGroups").
		Preload("Messages", func(db *gorm.DB) *gorm.DB {
			return db.Joins("JOIN (SELECT chat_id, MAX(created_at) AS max_created_at FROM messages GROUP BY chat_id) last_messages ON messages.chat_id = last_messages.chat_id AND messages.created_at = last_messages.max_created_at")
		}).
		Preload("Messages.Sender").
		Where("chat_users.user_id = ?", userID).
		Find(&chats).Error

	if err != nil {
		return nil, err
	}
	return chats, nil
}

func (r *gormChatRepository) Delete(ctx context.Context, chatID uint) error {
	return r.db.WithContext(ctx).Delete(&models.Chat{}, chatID).Error
}

func (r *gormChatRepository) AddChatUser(ctx context.Context, chat *models.Chat, userToAdd *models.User) error {
	// Добавление пользователя в чат
	if err := r.db.WithContext(ctx).Model(chat).Association("ChatUsers").Append(userToAdd); err != nil {
		return err
	}

	return nil
}

func (r *gormChatRepository) RemoveChatUser(ctx context.Context, chat *models.Chat, userToRemove *models.User) error {
	// Удалить пользователя из чата
	if err := r.db.WithContext(ctx).Model(chat).Association("ChatUsers").Delete(userToRemove); err != nil {
		return err
	}
	return nil
}

func (r *gormChatRepository) AddGroup(ctx context.Context, chat *models.Chat, group *models.Group) error {
	if err := r.db.WithContext(ctx).Model(chat).Association("ChatGroups").Append(group); err != nil {
		return err
	}
	return nil
}

func (r *gormChatRepository) RemoveGroup(ctx context.Context, chat *models.Chat, group *models.Group) error {
	if err := r.db.WithContext(ctx).Model(chat).Association("ChatGroups").Delete(group); err != nil {
		return err
	}
	return nil
}
