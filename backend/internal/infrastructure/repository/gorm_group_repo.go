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

	if err := r.db.WithContext(ctx).Preload("GroupUsers").First(&group, groupID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("group not found")
		}
		return nil, err
	}

	return &group, nil
}

// Найти группу по имени
func (r *gormGroupRepository) FindGroupByName(ctx context.Context, groupName string) (*models.Group, error) {
	var group models.Group

	// Выполняем запрос с проверкой по имени группы
	if err := r.db.WithContext(ctx).
		Preload("GroupUsers").
		Where("name = ?", groupName).
		First(&group).Error; err != nil { // Добавляем вызов First для выполнения запроса
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Возвращаем nil вместо ошибки, если группа не найдена
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
		Preload("GroupUsers").
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
	if err := r.db.WithContext(ctx).Model(group).Association("GroupUsers").Append(userToAdd); err != nil {
		return err
	}

	return nil
}

// Проверка, что пользователь является админом
func (r *gormGroupRepository) CheckUserIsAdmin(ctx context.Context, userID uint) (bool, error) {
	var group models.Group

	// Присоединяем таблицу пользователей и проверяем, состоит ли пользователь с указанным userID в группе "admins"
	err := r.db.WithContext(ctx).
		Preload("GroupUsers").
		Joins("JOIN group_users ON group_users.group_id = groups.id AND group_users.user_id = ?", userID).
		Where("groups.name = ?", "admins").
		First(&group).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil // Если пользователь не найден в группе "admins", возвращаем false
		}
		return false, err // Возвращаем ошибку, если возникли другие проблемы
	}

	return true, nil // Пользователь найден в группе "admins", возвращаем true
}

// Удалить пользователя из группы
func (r *gormGroupRepository) RemoveUserFromGroup(ctx context.Context, group *models.Group, userToRemove *models.User) error {

	if err := r.db.WithContext(ctx).Model(group).Association("GroupUsers").Delete(userToRemove); err != nil {
		return err
	}
	return nil
}
