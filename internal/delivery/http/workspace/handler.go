package workspace

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ndt080/schedule-manager-backend/internal/delivery/http/response"
	"github.com/ndt080/schedule-manager-backend/internal/domain"
	"github.com/ndt080/schedule-manager-backend/internal/service"
	manager "github.com/ndt080/schedule-manager-backend/pkg/auth"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	services     *service.Service
	tokenManager manager.TokenManager
}

func NewWorkspaceHandler(services *service.Service, tokenManager manager.TokenManager) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
	}
}

// GetWorkspaceById godoc
// @Description  Get workspace by id
// @Tags         Workspace
// @Accept       json
// @Produce      json
// @Param        id       path      string                true  "Workspace id"
// @Success      200      {object}  domain.WorkspaceData
// @Failure      400,500  {object}  error.ServerErrorResponse
// @Router       /workspace/{id} [get]
// @Security     AuthorizationKey
func (handler *Handler) GetWorkspaceById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	data, err := handler.services.Workspace.GetWorkspaceById(id)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, response.NewServerInternalError(err.Error()))
		return
	}

	context.JSON(http.StatusOK, data)
}

// GetUserWorkspaces godoc
// @Description  Get current user workspaces
// @Tags         Workspace
// @Accept       json
// @Produce      json
// @Success      200      {object}  []domain.WorkspaceData
// @Failure      400,500  {object}  error.ServerErrorResponse
// @Router       /workspaces/me [get]
// @Security     AuthorizationKey
func (handler *Handler) GetUserWorkspaces(context *gin.Context) {
	user, err := handler.services.Authorization.GetAuthorizedUser(context)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	data, err := handler.services.Workspace.GetWorkspacesByUser(user.ID)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, response.NewServerInternalError(err.Error()))
		return
	}

	context.JSON(http.StatusOK, data)
}

// RemoveWorkspace godoc
// @Description  Remove workspace by id
// @Tags         Workspace
// @Accept       json
// @Produce      json
// @Param        id       path      string                  true  "Workspace id"
// @Success      200      {object}  success.ServerSuccessResponse
// @Failure      400,500  {object}  error.ServerErrorResponse
// @Router       /workspace/{id} [delete]
// @Security     AuthorizationKey
func (handler *Handler) RemoveWorkspace(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	if err := handler.services.Workspace.RemoveWorkspace(id); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, response.NewServerInternalError(err.Error()))
		return
	}

	message := fmt.Sprintf("Object with id %d was removed", id)
	context.JSON(http.StatusOK, response.NewServerSuccessResponse(message))
}

// CreateWorkspace godoc
// @Description  Create workspace
// @Tags         Workspace
// @Accept       json
// @Produce      json
// @Param        input    body      WorkspaceRequest  true  "Workspace data"
// @Success      200      {object}  domain.WorkspaceData
// @Failure      400,500  {object}  error.ServerErrorResponse
// @Router       /workspace [post]
// @Security     AuthorizationKey
func (handler *Handler) CreateWorkspace(context *gin.Context) {
	user, err := handler.services.Authorization.GetAuthorizedUser(context)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	json := WorkspaceRequest{}
	if err := context.ShouldBindJSON(&json); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	workspace := domain.Workspace{
		Name:        json.Name,
		Description: json.Description,
		Image:       json.Image,
		CreatedAt:   time.Now(),
	}

	data, err := handler.services.Workspace.CreateWorkspace(*user, workspace)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	context.JSON(http.StatusOK, data)
}

// UpdateWorkspace godoc
// @Description  Update workspace
// @Tags         Workspace
// @Accept       json
// @Produce      json
// @Param        input    body      WorkspaceRequest  true  "Workspace data"
// @Success      200      {object}  domain.WorkspaceData
// @Failure      400,500  {object}  error.ServerErrorResponse
// @Router       /workspace [put]
// @Security     AuthorizationKey
func (handler *Handler) UpdateWorkspace(context *gin.Context) {
	json := WorkspaceRequest{}
	if err := context.ShouldBindJSON(&json); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	workspace := domain.Workspace{
		ID:          json.ID,
		Name:        json.Name,
		Description: json.Description,
		Image:       json.Image,
	}

	data, err := handler.services.Workspace.UpdateWorkspace(workspace)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	context.JSON(http.StatusOK, data)
}

// SearchWorkspaces godoc
// @Description  Search workspaces
// @Tags         Workspace
// @Accept       json
// @Produce      json
// @Param        name     query     string  false  "Workspace name"
// @Success      200      {object}  []domain.WorkspaceData
// @Failure      400,500  {object}  error.ServerErrorResponse
// @Router       /workspaces/search [get]
// @Security     AuthorizationKey
func (handler *Handler) SearchWorkspaces(context *gin.Context) {
	name, _ := context.GetQuery("name")

	data, err := handler.services.Workspace.SearchWorkspaces(name)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, response.NewServerInternalError(err.Error()))
		return
	}

	context.JSON(http.StatusOK, data)
}

// CreateWorkspaceTask godoc
// @Description  Create workspace task
// @Tags         Workspace
// @Accept       json
// @Produce      json
// @Param        id       path      string  true  "Workspace id"
// @Param        input    body      WorkspaceTaskRequest  true  "Workspace task data"
// @Success      200      {object}  domain.WorkspaceTask
// @Failure      400,500  {object}  error.ServerErrorResponse
// @Router       /workspace/{id}/task [post]
// @Security     AuthorizationKey
func (handler *Handler) CreateWorkspaceTask(context *gin.Context) {
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

	json := WorkspaceTaskRequest{}
	if err := context.ShouldBindJSON(&json); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	task := domain.WorkspaceTask{
		Name:        json.Name,
		Description: json.Description,
		Creator:     user.ID,
		WorkspaceId: id,
	}

	data, err := handler.services.Workspace.CreateWorkspaceTask(task)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	context.JSON(http.StatusOK, data)
}

// RemoveWorkspaceTask godoc
// @Description  Remove workspace by id
// @Tags         Workspace
// @Accept       json
// @Produce      json
// @Param        id       path      string  true  "Workspace task id"
// @Success      200      {object}  success.ServerSuccessResponse
// @Failure      400,500  {object}  error.ServerErrorResponse
// @Router       /workspace/task/{id} [delete]
// @Security     AuthorizationKey
func (handler *Handler) RemoveWorkspaceTask(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	if err := handler.services.Workspace.RemoveWorkspaceTask(id); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, response.NewServerInternalError(err.Error()))
		return
	}

	message := fmt.Sprintf("Workspace task with id %d was removed", id)
	context.JSON(http.StatusOK, response.NewServerSuccessResponse(message))
}

// CreateWorkspaceMember godoc
// @Description  Create workspace task
// @Tags         Workspace
// @Accept       json
// @Produce      json
// @Param        id       path      string  true  "Workspace id"
// @Param        input    body      WorkspaceMemberRequest  true  "Workspace member data"
// @Success      200      {object}  domain.WorkspaceMember
// @Failure      400,500  {object}  error.ServerErrorResponse
// @Router       /workspace/{id}/member [post]
// @Security     AuthorizationKey
func (handler *Handler) CreateWorkspaceMember(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	json := WorkspaceMemberRequest{}
	if err := context.ShouldBindJSON(&json); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	user, err := handler.services.User.GetUserByEmail(json.Email)
	if err != nil || user == nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, response.NewServerInvalidCredentialsError())
		return
	}

	data, err := handler.services.Workspace.AddMemberToWorkspace(domain.WorkspaceMember{
		Member:      *user,
		Status:      json.Status,
		WorkspaceId: id,
	})
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	context.JSON(http.StatusOK, data)
}

// GetInviteToken godoc
// @Description  Get invite token
// @Tags         Workspace
// @Accept       json
// @Produce      json
// @Param        id       path      string  true  "Workspace id"
// @Success      200      {object}  WorkspaceInviteResponse
// @Failure      400,500  {object}  error.ServerErrorResponse
// @Router       /workspace/{id}/invite [get]
// @Security     AuthorizationKey
func (handler *Handler) GetInviteToken(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	token, err := handler.tokenManager.NewInviteToken(id, 24*time.Hour)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, response.NewServerInternalError(err.Error()))
		return
	}

	context.JSON(http.StatusOK, WorkspaceInviteResponse{
		Token: token,
	})
}
