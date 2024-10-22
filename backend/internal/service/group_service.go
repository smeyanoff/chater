package service

import (
	models "chater/internal/domain/entity"
	repo "chater/internal/domain/repository"
	"chater/internal/domain/valueobject"
	"context"
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

	// Поиск пользователя по ID
	owner, err := gc.userRepo.FindByID(ctx, ownerID)
	if err != nil {
		return nil, err
	}

	newGroup := &models.Group{
		Name:         groupName,
		OwnerID:      ownerID,
		GroupMembers: []*models.User{owner},
	}
	if err := gc.groupRepo.Save(ctx, newGroup); err != nil {
		return nil, err
	}
	return newGroup, nil
}

func (gc *GroupService) DeleteGroup(ctx context.Context, ownerID uint)

// // Проверяем, является ли текущий пользователь владельцем группы
// if group.OwnerID != groupOwnerID {
// 	return errors.New("only the group owner can add users to the group")
// }

// // Проверяем, является ли пользователь уже участником группы
// for _, member := range group.GroupMembers {
// 	if member.ID == userToAdd.ID {
// 		return errors.New("user is already in the group")
// 	}
// }

// // Найти группу по ID
// group, err := r.FindGroupByID(ctx, groupID)
// if err != nil {
// 	return err
// }

// Найти группу по ID
// group, err := r.FindGroupByID(ctx, groupID)
// if err != nil {
// 	return err
// }

// // Проверяем, является ли текущий пользователь владельцем группы
// if group.OwnerID != groupOwnerID {
// 	return errors.New("only the group owner can remove users from the group")
// }

// // Проверяем, является ли пользователь участником группы
// var userFound bool
// for _, member := range group.GroupMembers {
// 	if member.ID == userToRemove.ID {
// 		userFound = true
// 		break
// 	}
// }
