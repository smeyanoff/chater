package repository

import (
	models "chater/internal/domain/entity"
	_ "chater/internal/domain/valueobject"
	"context"
)

type GroupRepository interface {
	Save(ctx context.Context, group *models.Group) error
	Delete(ctx context.Context, groupID uint) error
	FindGroupByID(ctx context.Context, groupID uint) (*models.Group, error)
	FindGroupByName(ctx context.Context, groupName string) (*models.Group, error)
	FindAllUserGroups(ctx context.Context, user uint) ([]*models.Group, error)
	CheckUserIsAdmin(ctx context.Context, userID uint) (bool, error)
	CheckUserIsGroupMember(ctx context.Context, groupID uint, userID uint) (bool, error)
	AddUserToGroup(ctx context.Context, group *models.Group, userToAdd *models.User) error
	RemoveUserFromGroup(ctx context.Context, group *models.Group, userToRemove *models.User) error
}
