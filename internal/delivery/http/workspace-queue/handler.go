package workspace_queue

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ndt080/schedule-manager-backend/internal/delivery/http/response"
	"github.com/ndt080/schedule-manager-backend/internal/domain"
	"github.com/ndt080/schedule-manager-backend/internal/service"
	"net/http"
	"strconv"
)

type WorkspaceQueueHandler struct {
	services *service.Service
}

func NewWorkspaceQueueHandler(services *service.Service) *WorkspaceQueueHandler {
	return &WorkspaceQueueHandler{
		services: services,
	}
}

// CreateWorkspaceQueue godoc
// @Description  Create workspace queue
// @Tags         Workspace
// @Accept       json
// @Produce      json
// @Param        id       path      string                 true  "Workspace id"
// @Param        input    body      WorkspaceQueueRequest  true  "Workspace queue data"
// @Success      200      {object}  domain.WorkspaceQueue
// @Failure      400,500  {object}  error.ServerErrorResponse
// @Router       /workspace/{id}/queue [post]
// @Security     AuthorizationKey
func (handler *WorkspaceQueueHandler) CreateWorkspaceQueue(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	json := WorkspaceQueueRequest{}
	if err := context.ShouldBindJSON(&json); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	queue := domain.WorkspaceQueue{
		Name:        json.Name,
		WorkspaceId: id,
	}

	data, err := handler.services.Workspace.CreateWorkspaceQueue(queue)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	context.JSON(http.StatusOK, data)
}

// JoinToWorkspaceQueue godoc
// @Description  Create workspace queue
// @Tags         Workspace
// @Accept       json
// @Produce      json
// @Param        id       path      string  true  "Workspace queue id"
// @Success      200      {object}  domain.WorkspaceQueue
// @Failure      400,500  {object}  error.ServerErrorResponse
// @Router       /workspace/queue/{id}/join [post]
// @Security     AuthorizationKey
func (handler *WorkspaceQueueHandler) JoinToWorkspaceQueue(context *gin.Context) {
	user, err := handler.services.Authorization.GetAuthorizedUser(context)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	data, err := handler.services.Workspace.JoinToWorkspaceQueue(user, id)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	context.JSON(http.StatusOK, data)
}

// LeaveWorkspaceQueue godoc
// @Description  Create workspace queue
// @Tags         Workspace
// @Accept       json
// @Produce      json
// @Param        id       path      string  true  "Workspace queue id"
// @Param        uid      path      string  true  "User id"
// @Success      200      {object}  domain.WorkspaceQueue
// @Failure      400,500  {object}  error.ServerErrorResponse
// @Router       /workspace/queue/{id}/leave/{uid} [delete]
// @Security     AuthorizationKey
func (handler *WorkspaceQueueHandler) LeaveWorkspaceQueue(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	uid, err := strconv.ParseInt(context.Param("uid"), 10, 64)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	data, err := handler.services.Workspace.LeaveWorkspaceQueue(uid, id)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	context.JSON(http.StatusOK, data)
}

// RemoveWorkspaceQueue godoc
// @Description  Remove workspace queue by id
// @Tags         Workspace
// @Accept       json
// @Produce      json
// @Param        id       path      string  true  "Workspace queue id"
// @Success      200      {object}  success.ServerSuccessResponse
// @Failure      400,500  {object}  error.ServerErrorResponse
// @Router       /workspace/queue/{id} [delete]
// @Security     AuthorizationKey
func (handler *WorkspaceQueueHandler) RemoveWorkspaceQueue(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	if err := handler.services.Workspace.RemoveWorkspaceQueue(id); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, response.NewServerInternalError(err.Error()))
		return
	}

	message := fmt.Sprintf("Workspace task with id %d was removed", id)
	context.JSON(http.StatusOK, response.NewServerSuccessResponse(message))
}
