package service

import (
	models "chater/internal/domain/entity"
	repo "chater/internal/domain/repository"
	"chater/internal/domain/valueobject"
	"context"
	"errors"

	"gorm.io/gorm"
)

type GroupService struct {
	groupRepo repo.GroupRepository
	userRepo  repo.UserRepository
}

func NewGroupService(groupRepo repo.GroupRepository, userRepo repo.UserRepository) GroupService {
	return GroupService{groupRepo: groupRepo, userRepo: userRepo}
}

func (gc *GroupService) CreateGroup(ctx context.Context, name string, ownerID uint) (*models.Group, error) {
	// Валидация имени группы
	groupName, err := valueobject.NewGroupName(name)
	if err != nil {
		return nil, err
	}

	// Провека существования группы с таким именем
	existedGroup, err := gc.groupRepo.FindGroupByName(ctx, groupName.String())
	if existedGroup != nil {
		return nil, errors.New("group already exists")
	}
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}

	// Поиск пользователя по ID
	owner, err := gc.userRepo.FindByID(ctx, ownerID)
	if err != nil {
		return nil, err
	}

	newGroup := &models.Group{
		Name:       groupName,
		OwnerID:    ownerID,
		GroupUsers: []*models.User{owner},
	}
	if err := gc.groupRepo.Save(ctx, newGroup); err != nil {
		return nil, err
	}
	return newGroup, nil
}

func (gc *GroupService) DeleteGroup(ctx context.Context, ownerID uint, groupId uint) error {

	user, err := gc.userRepo.FindByID(ctx, ownerID)
	if err != nil {
		return err
	}

	group, err := gc.groupRepo.FindGroupByID(ctx, groupId)
	if err != nil {
		return err
	}

	// Проверка удаления системных групп
	if group.Name.String() == "admins" {
		return errors.New("group 'admins' couldn't be deleted")
	}

	// Проверка возможности удаления группы пользователем
	isUserAdmin, err := gc.groupRepo.CheckUserIsAdmin(ctx, user.ID)
	if err != nil {
		return err
	}
	// Если пользователь владелец группы или администратор
	if group.OwnerID != user.ID || !isUserAdmin {
		return errors.New("only group owner and admin can delete group")
	}

	if err := gc.groupRepo.Delete(ctx, group.ID); err != nil {
		return err
	}

	return nil
}
