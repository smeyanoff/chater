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

func NewGroupService(groupRepo repo.GroupRepository, userRepo repo.UserRepository) *GroupService {
	return &GroupService{groupRepo: groupRepo, userRepo: userRepo}
}

// Проверка возможности удаления группы пользователем
func (gc *GroupService) checkRights(ctx context.Context, ownerID uint, userID uint) error {
	isUserAdmin, err := gc.groupRepo.CheckUserIsAdmin(ctx, userID)
	if err != nil {
		return err
	}
	// Если пользователь владелец группы или администратор
	if ownerID == userID || !isUserAdmin {
		return errors.New("only group owner and admin can do this")
	}

	return nil
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
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	// Поиск пользователя по ID
	owner, err := gc.userRepo.FindUserByID(ctx, ownerID)
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

func (gc *GroupService) DeleteGroup(ctx context.Context, ownerID uint, groupID uint) error {

	group, err := gc.groupRepo.FindGroupByID(ctx, groupID)
	if err != nil {
		return err
	}

	if err := gc.checkRights(ctx, group.OwnerID, ownerID); err != nil {
		return err
	}

	// Проверка удаления системных групп
	if group.Name.String() == "admins" {
		return errors.New("group 'admins' couldn't be deleted")
	}

	if err := gc.groupRepo.Delete(ctx, group.ID); err != nil {
		return err
	}

	return nil
}

func (gc *GroupService) AddUserToGroup(ctx context.Context, ownerID uint, userToAddID uint, groupID uint) error {

	group, err := gc.groupRepo.FindGroupByID(ctx, groupID)
	if err != nil {
		return err
	}

	userToAdd, err := gc.userRepo.FindUserByID(ctx, userToAddID)
	if err != nil {
		return err
	}

	if err := gc.checkRights(ctx, group.OwnerID, ownerID); err != nil {
		return err
	}

	if err := gc.groupRepo.AddUserToGroup(ctx, group, userToAdd); err != nil {
		return err
	}

	return nil
}

func (gc *GroupService) DeleteUserFromGroup(ctx context.Context, ownerID uint, userToRemoveID uint, groupID uint) error {
	group, err := gc.groupRepo.FindGroupByID(ctx, groupID)
	if err != nil {
		return err
	}

	userToRemove, err := gc.userRepo.FindUserByID(ctx, userToRemoveID)
	if err != nil {
		return err
	}

	if err := gc.checkRights(ctx, group.OwnerID, ownerID); err != nil {
		return err
	}

	if err := gc.groupRepo.RemoveUserFromGroup(ctx, group, userToRemove); err != nil {
		return err
	}

	return nil
}
