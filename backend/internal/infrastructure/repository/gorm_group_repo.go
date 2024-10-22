package repository

import (
	models "chater/internal/domain/entity"
	"chater/internal/domain/repository"
	_ "chater/internal/domain/valueobject"
	"context"
	"errors"

	"gorm.io/gorm"
)

type gormGroupRepository struct {
	db *gorm.DB
}

func NewGormGroupRepository(db *gorm.DB) repository.GroupRepository {
	return &gormGroupRepository{db: db}
}

// Сохранение группы
func (r *gormGroupRepository) Save(ctx context.Context, group *models.Group) error {
	return r.db.WithContext(ctx).Create(group).Error
}

// Удаление группы
func (r *gormGroupRepository) Delete(ctx context.Context, groupID uint) error {
	return r.db.WithContext(ctx).Delete(&models.Group{}, groupID).Error
}

// Найти группу по ID
func (r *gormGroupRepository) FindGroupByID(ctx context.Context, groupID uint) (*models.Group, error) {
	var group models.Group

	if err := r.db.WithContext(ctx).Preload("GroupMembers").First(&group, groupID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("group not found")
		}
		return nil, err
	}

	return &group, nil
}

// Найти все группы пользователя
func (r *gormGroupRepository) FindAllUserGroups(ctx context.Context, userID uint) ([]*models.Group, error) {
	var groups []*models.Group

	err := r.db.WithContext(ctx).
		Joins("JOIN group_users ON group_users.group_id = groups.id").
		Preload("GroupMembers").
		Where("group_users.user_id = ?", userID).
		Find(&groups).Error

	if err != nil {
		return nil, err
	}

	return groups, nil
}

// Добавить пользователя в группу
func (r *gormGroupRepository) AddUserToGroup(ctx context.Context, group *models.Group, userToAdd *models.User) error {

	// Добавление пользователя в группу
	if err := r.db.WithContext(ctx).Model(group).Association("GroupMembers").Append(userToAdd); err != nil {
		return err
	}

	return nil
}

// Удалить пользователя из группы
func (r *gormGroupRepository) RemoveUserFromGroup(ctx context.Context, group *models.Group, userToRemove *models.User) error {

	if err := r.db.WithContext(ctx).Model(group).Association("GroupMembers").Delete(userToRemove); err != nil {
		return err
	}
	return nil
}
