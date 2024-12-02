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

func (r *gormChatRepository) GetChatByID(ctx context.Context, chatID uint) (*models.Chat, error) {
	var chat *models.Chat
	if err := r.db.WithContext(ctx).
		Preload("ChatUsers").
		Preload("ChatGroups").
		First(&chat, chatID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("chat not found")
		} else {
			return nil, err
		}
	}
	return chat, nil
}

func (r *gormChatRepository) GetUserChats(ctx context.Context, userID uint) ([]*models.Chat, error) {
	var chats []*models.Chat

	err := r.db.WithContext(ctx).
		Joins("JOIN chat_users ON chat_users.chat_id = chats.id").
		Preload("ChatUsers").
		Preload("ChatGroups", func(db *gorm.DB) *gorm.DB {
			// Указываем таблицу groups, которая соединяется с group_users
			return db.Joins("JOIN group_users ON group_users.group_id = groups.id").
				Where("group_users.user_id = ?", userID)
		}).
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
