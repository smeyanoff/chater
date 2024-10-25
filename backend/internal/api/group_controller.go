package api

import (
	"chater/internal/logging"
	"chater/internal/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GroupController struct {
	groupService *service.GroupService
}

func NewGroupController(groupService *service.GroupService) *GroupController {
	return &GroupController{groupService: groupService}
}

func (gc *GroupController) CreateGroup(ctx *gin.Context) {
	logging.Logger.Debug("Create group request...")
	var request createGroupRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logging.Logger.Error(err.Error())
		logging.Logger.Sugar().Debug(request)
		ctx.JSON(http.StatusBadRequest, errorResponse{Error: "Invalid request"})
		return
	}

	userID, exists := ctx.Get("user_id")
	if !exists {
		logging.Logger.Error("UserID doesn't exist")
		ctx.JSON(http.StatusUnauthorized, errorResponse{Error: "Unauthorized"})
		return
	}

	group, err := gc.groupService.CreateGroup(ctx, request.Name, userID.(uint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			errorResponse{Error: fmt.Sprintf("Error create group: %s", err.Error())})
	}

	ctx.JSON(http.StatusOK, mapGroup(group, userID.(uint)))
}
