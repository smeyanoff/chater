package api

import (
	"chater/internal/logging"
	"chater/internal/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GroupController struct {
	groupService *service.GroupService
}

func NewGroupController(groupService *service.GroupService) *GroupController {
	return &GroupController{groupService: groupService}
}

// CreateGroup создает новую группу
// @Summary Создание группы
// @Description Создает новую группу с указанным именем для авторизованного пользователя
// @Tags Groups, V1
// @Accept json
// @Produce json
// @Param createGroupRequest body createGroupRequest true "Данные для создания группы"
// @Success 200 {object} groupResponse "Группа успешно создана"
// @Failure 400 {object} errorResponse "Неверный запрос"
// @Failure 401 {object} errorResponse "Пользователь не авторизован"
// @Failure 500 {object} errorResponse "Ошибка при создании группы"
// @Security BearerAuth
// @Router /v1/groups [post]
func (gc *GroupController) CreateGroup(ctx *gin.Context) {
	logging.Logger.Debug("Create group request...")
	var request createGroupRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logging.Logger.Error(err.Error())
		logging.Logger.Sugar().Debug(request)
		ctx.JSON(http.StatusBadRequest, errorResponse{Error: ErrInvalidRequest})
		return
	}

	userID, exists := ctx.Get("user_id")
	if !exists {
		logging.Logger.Error(ErrUnauthorized)
		ctx.JSON(http.StatusUnauthorized, errorResponse{Error: ErrUnauthorized})
		return
	}

	group, err := gc.groupService.CreateGroup(ctx, request.Name, userID.(uint))
	if err != nil {
		logging.Logger.Error(fmt.Sprintf("Error create group: %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError,
			errorResponse{Error: fmt.Sprintf("Error create group: %s", err.Error())})
	}

	ctx.JSON(http.StatusOK, mapGroup(group, userID.(uint)))
	logging.Logger.Debug("Group created")
}

// DeleteGroup удаляет группу по её идентификатору
// @Summary Удаление группы
// @Description Удаляет группу по её идентификатору, если пользователь является её владельцем
// @Tags Groups, V1
// @Param group_id path string true "ID группы для удаления"
// @Success 200 {object} successResponse "Группа успешно удалена"
// @Failure 400 {object} errorResponse "Неверный идентификатор группы"
// @Failure 401 {object} errorResponse "Пользователь не авторизован"
// @Failure 500 {object} errorResponse "Ошибка при удалении группы"
// @Security BearerAuth
// @Router /v1/groups/{group_id} [delete]
func (gc *GroupController) DeleteGroup(ctx *gin.Context) {

	groupID := ctx.Param("group_id")
	logging.Logger.Debug(fmt.Sprintf("Delete %s group request...", groupID))
	groupIDuint, err := strconv.ParseUint(groupID, 10, 32)
	if err != nil {
		logging.Logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse{Error: "Invalid Group ID"})
		return
	}

	userID, exists := ctx.Get("user_id")
	if !exists {
		logging.Logger.Error(ErrUnauthorized)
		ctx.JSON(http.StatusUnauthorized, errorResponse{Error: "Unauthorized"})
		return
	}

	if err := gc.groupService.DeleteGroup(ctx, userID.(uint), uint(groupIDuint)); err != nil {
		logging.Logger.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError,
			errorResponse{Error: fmt.Sprintf("Error delete group: %s", err.Error())})
		return
	}

	ctx.JSON(http.StatusOK, successResponse{Message: "Group deleted successfully"})
	logging.Logger.Debug("Group deleted")
}

// AddUserToChat добавляет пользователя в группу
// @Summary Добавление пользователя в группу
// @Description Добавляет пользователя в группу по идентификатору группы
// @Tags Groups, V1
// @Accept json
// @Produce json
// @Param group_id path string true "ID группы"
// @Param addUserToGroupRequest body userGroupRequest true "Данные для добавления пользователя"
// @Success 200 {object} successResponse "Пользователь успешно добавлен в группу"
// @Failure 400 {object} errorResponse "Неверный запрос или неверный ID группы"
// @Failure 401 {object} errorResponse "Пользователь не авторизован"
// @Failure 500 {object} errorResponse "Ошибка при добавлении пользователя в группу"
// @Security BearerAuth
// @Router /v1/groups/{group_id}/users [post]
func (gc *GroupController) AddUserToGroup(ctx *gin.Context) {

	groupID := ctx.Param("group_id")
	logging.Logger.Debug(fmt.Sprintf("Add user %s group request...", groupID))

	var request userGroupRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logging.Logger.Error(err.Error())
		logging.Logger.Sugar().Debug(request)
		ctx.JSON(http.StatusBadRequest, errorResponse{Error: ErrInvalidRequest})
		return
	}

	groupIDuint, err := strconv.ParseUint(groupID, 10, 32)
	if err != nil {
		logging.Logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse{Error: "Invalid Group ID"})
		return
	}

	ownerID, exists := ctx.Get("user_id")
	if !exists {
		logging.Logger.Error(ErrUnauthorized)
		ctx.JSON(http.StatusUnauthorized, errorResponse{Error: ErrUnauthorized})
		return
	}

	if err := gc.groupService.AddUserToGroup(ctx, ownerID.(uint), request.UserID, uint(groupIDuint)); err != nil {
		logging.Logger.Error(fmt.Sprintf("Error add user to group: %s", err.Error()))
		ctx.JSON(http.StatusUnauthorized, errorResponse{Error: fmt.Sprintf("Error add user to group: %s", err.Error())})
		return
	}

	ctx.JSON(http.StatusOK, successResponse{Message: "User added to group"})
	logging.Logger.Debug("User added suceffully")
}

// DeleteUserFromGroup удаляет пользователя из группы
// @Summary Удаление пользователя из группы
// @Description Удаляет указанного пользователя из группы, если действие выполняет владелец или администратор группы
// @Tags Groups, V1
// @Accept json
// @Produce json
// @Param group_id path string true "ID группы"
// @Param userGroupRequest body userGroupRequest true "ID пользователя для удаления из группы"
// @Success 200 {object} successResponse "Пользователь успешно удалён из группы"
// @Failure 400 {object} errorResponse "Неверный запрос или неверный ID группы"
// @Failure 401 {object} errorResponse "Пользователь не авторизован"
// @Failure 500 {object} errorResponse "Ошибка при удалении пользователя из группы"
// @Security BearerAuth
// @Router /v1/groups/{group_id}/users [delete]
func (gc *GroupController) DeleteUserFromGroup(ctx *gin.Context) {

	groupID := ctx.Param("group_id")
	logging.Logger.Debug(fmt.Sprintf("Delete user from group %s request...", groupID))

	var request userGroupRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logging.Logger.Error(err.Error())
		logging.Logger.Sugar().Debug(request)
		ctx.JSON(http.StatusBadRequest, errorResponse{Error: ErrInvalidRequest})
		return
	}

	groupIDuint, err := strconv.ParseUint(groupID, 10, 32)
	if err != nil {
		logging.Logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse{Error: ErrInvalidGroupID})
		return
	}

	ownerID, exists := ctx.Get("user_id")
	if !exists {
		logging.Logger.Error(ErrUnauthorized)
		ctx.JSON(http.StatusUnauthorized, errorResponse{Error: ErrUnauthorized})
		return
	}

	if err := gc.groupService.DeleteUserFromGroup(ctx, ownerID.(uint), request.UserID, uint(groupIDuint)); err != nil {
		logging.Logger.Error(fmt.Sprintf("Error deleting user from group: %s", err.Error()))
		ctx.JSON(http.StatusUnauthorized, errorResponse{Error: fmt.Sprintf("Error deleting user from group: %s", err.Error())})
		return
	}

	ctx.JSON(http.StatusOK, successResponse{Message: "User deleted from group"})
	logging.Logger.Debug("User deleted successfully")
}

// GetAllUserGroups godoc
// @Summary Получить все группы пользователя
// @Description Возвращает список всех групп, в которых состоит текущий авторизованный пользователь
// @Tags Groups
// @Produce json
// @Security BearerAuth
// @Success 200 {object} groupsResponse "Список групп пользователя"
// @Failure 401 {object} errorResponse "Пользователь не авторизован"
// @Failure 500 {object} errorResponse "Ошибка сервера"
// @Router /v1/groups [get]
func (gc *GroupController) GetAllUserGroups(ctx *gin.Context) {
	logging.Logger.Debug("Get all user group response...")
	userID, exists := ctx.Get("user_id")
	if !exists {
		logging.Logger.Error(ErrUnauthorized)
		ctx.JSON(http.StatusUnauthorized, errorResponse{Error: ErrUnauthorized})
		return
	}

	groups, err := gc.groupService.GetAllUserGroups(ctx, userID.(uint))
	if err != nil {
		logging.Logger.Error(fmt.Sprintf("Error get user groups: %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("Error get user groups: %s", err.Error()))
	}

	ctx.JSON(http.StatusOK, groupsResponse{Groups: mapGroups(groups, userID.(uint))})
}
