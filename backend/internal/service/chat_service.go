package service

import (
	models "chater/internal/domain/entity"
	repo "chater/internal/domain/repository"
	"chater/internal/domain/valueobject"
	"context"
	"errors"
)

type ChatService struct {
	chatRepo  repo.ChatRepository
	userRepo  repo.UserRepository
	groupRepo repo.GroupRepository
}

func NewChatService(chatRepo repo.ChatRepository, userRepo repo.UserRepository, groupRepo repo.GroupRepository) *ChatService {
	return &ChatService{
		chatRepo:  chatRepo,
		userRepo:  userRepo,
		groupRepo: groupRepo,
	}
}

/**
 * checkRights checks if the user has the rights to perform a specific action in the chat.
 * It verifies if the user is the owner of the chat group or an administrator.
 *
 * @param ctx context.Context - The context of the request.
 * @param ownerID uint - The ID of the owner of the chat group.
 * @param userID uint - The ID of the user whose rights are being checked.
 * @return error - Returns an error if the user does not have the necessary rights.
 */
func (cc *ChatService) checkRights(ctx context.Context, ownerID uint, userID uint) error {
	isUserAdmin, err := cc.groupRepo.CheckUserIsAdmin(ctx, userID)
	if err != nil {
		return err
	}
	// Если пользователь владелец группы или администратор
	if ownerID != userID && !isUserAdmin {
		return errors.New("only chat owner and admin can do this")
	}

	return nil
}

/**
* Creates a new instance of ChatService with the provided repositories for chat, user, and group.
*
* Parameters:
* - chatRepo: The repository for chat operations.
* - userRepo: The repository for user operations.
* - groupRepo: The repository for group operations.
*
* Returns:
* - A pointer to the newly created ChatService instance.
 */
func (cc *ChatService) CreateChat(ctx context.Context, chatName string, ownerID uint) (*models.Chat, error) {
	validatedChatName, err := valueobject.NewChatName(chatName)
	if err != nil {
		return nil, err
	}

	owner, err := cc.userRepo.FindUserByID(ctx, ownerID)
	if err != nil {
		return nil, err
	}

	newChat := &models.Chat{
		Name:      validatedChatName,
		OwnerID:   ownerID,
		ChatUsers: []*models.User{owner},
	}
	if err := cc.chatRepo.Save(ctx, newChat); err != nil {
		return nil, err
	}
	return newChat, nil
}

/**
*  DeleteChat deletes a chat based on the provided chatID after checking user rights.
* Parameters:
* - ctx: the context.Context for the operation
* - userID: the ID of the user performing the delete action
* - chatID: the ID of the chat to be deleted
* Returns:
* - error: an error if any operation fails
 */
func (cc *ChatService) DeleteChat(ctx context.Context, userID uint, chatID uint) error {
	chat, err := cc.chatRepo.FindChatByID(ctx, chatID)
	if err != nil {
		return err
	}

	if err := cc.checkRights(ctx, chat.OwnerID, userID); err != nil {
		return err
	}

	if err := cc.chatRepo.Delete(ctx, chatID); err != nil {
		return err
	}

	return nil
}

/**
 * AddUserToChat adds a user to a chat.
 *
 * Parameters:
 * - ctx: the context in which the operation is performed
 * - userToAddID: the ID of the user to add to the chat
 * - chatID: the ID of the chat to which the user is added
 *
 * Returns:
 * - error: an error if the operation fails
 */
func (cc *ChatService) AddUserToChat(ctx context.Context, userToAddID uint, chatID uint) error {
	user, err := cc.userRepo.FindUserByID(ctx, userToAddID)
	if err != nil {
		return err
	}

	chat, err := cc.chatRepo.FindChatByID(ctx, chatID)
	if err != nil {
		return err
	}

	if err := cc.chatRepo.AddChatUser(ctx, chat, user); err != nil {
		return err
	}

	return nil
}

/**
 * DeleteUserFromChat removes a user from a chat.
 *
 * Parameters:
 * - ctx: the context.Context for the operation
 * - ownerID: the ID of the chat owner
 * - userToRemoveID: the ID of the user to be removed from the chat
 * - chatID: the ID of the chat from which the user is to be removed
 *
 * Returns:
 * - error: an error if any operation fails during the user removal process
 */
func (cc *ChatService) DeleteUserFromChat(ctx context.Context, ownerID uint, userToRemoveID uint, chatID uint) error {

	if ownerID == userToRemoveID {
		return errors.New("chat owner self remove forbidden")
	}

	user, err := cc.userRepo.FindUserByID(ctx, userToRemoveID)
	if err != nil {
		return err
	}

	chat, err := cc.chatRepo.FindChatByID(ctx, chatID)
	if err != nil {
		return err
	}

	if err := cc.checkRights(ctx, chat.OwnerID, ownerID); err != nil {
		return err
	}

	if err := cc.chatRepo.RemoveChatUser(ctx, chat, user); err != nil {
		return err
	}

	return nil
}

/**
 * GetUserChats retrieves all chats for a specific user.
 *
 * Parameters:
 * - ctx: the context for the request
 * - userID: the ID of the user to retrieve chats for
 *
 * Returns:
 * - []*models.Chat: a slice of Chat models
 * - error: an error if the operation fails
 */
func (cc *ChatService) GetUserChats(ctx context.Context, userID uint) ([]*models.Chat, error) {
	chats, err := cc.chatRepo.FindAllChatsWithLastMessage(ctx, userID)
	if err != nil {
		return nil, err
	}
	return chats, nil
}

func (cc *ChatService) AddGroupToChat(ctx context.Context, userID uint, chatID uint, groupID uint) error {

	chat, err := cc.chatRepo.FindChatByID(ctx, chatID)
	if err != nil {
		return err
	}

	user, err := cc.userRepo.FindUserByID(ctx, userID)
	if err != nil {
		return err
	}

	group, err := cc.groupRepo.FindGroupByID(ctx, groupID)
	if err != nil {
		return err
	}

	if err := cc.checkRights(ctx, chat.OwnerID, user.ID); err != nil {
		return err
	}

	if err := cc.chatRepo.AddGroup(ctx, chat, group); err != nil {
		return err
	}
	return nil
}

func (cc *ChatService) RemoveGroupFromChat(ctx context.Context, userID uint, chatID uint, groupID uint) error {

	chat, err := cc.chatRepo.FindChatByID(ctx, chatID)
	if err != nil {
		return err
	}

	user, err := cc.userRepo.FindUserByID(ctx, userID)
	if err != nil {
		return err
	}

	group, err := cc.groupRepo.FindGroupByID(ctx, groupID)
	if err != nil {
		return err
	}

	if err := cc.checkRights(ctx, chat.OwnerID, user.ID); err != nil {
		return err
	}

	if err := cc.chatRepo.RemoveGroup(ctx, chat, group); err != nil {
		return err
	}
	return nil
}
