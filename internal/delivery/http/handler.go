package http

import (
	"github.com/gin-gonic/gin"
	"github.com/ndt080/schedule-manager-backend/internal/delivery/http/auth"
	"github.com/ndt080/schedule-manager-backend/internal/delivery/http/user"
	"github.com/ndt080/schedule-manager-backend/internal/delivery/http/workspace"
	workspace_queue "github.com/ndt080/schedule-manager-backend/internal/delivery/http/workspace-queue"
	"github.com/ndt080/schedule-manager-backend/internal/service"
	manager "github.com/ndt080/schedule-manager-backend/pkg/auth"
)

type Auth interface {
	InitRoutes(auth *gin.RouterGroup)
	InitRoutesWithIdentifyUser(auth *gin.RouterGroup)
	SignIn(context *gin.Context)
	SignUp(context *gin.Context)
	RefreshToken(context *gin.Context)
	CheckAuthStatus(context *gin.Context)
	ConfirmEmailAgain(context *gin.Context)
}

type User interface {
	InitRoutes(auth *gin.RouterGroup)
	GetCurrentUser(context *gin.Context)
	GetUsers(context *gin.Context)
}

type Workspace interface {
	InitRoutes(auth *gin.RouterGroup)
}

type WorkspaceQueue interface {
	InitRoutes(auth *gin.RouterGroup)
}

type Handler struct {
	service        *service.Service
	tokenManager   manager.TokenManager
	Auth           Auth
	User           User
	Workspace      Workspace
	WorkspaceQueue WorkspaceQueue
}

func NewHandler(service *service.Service, tokenManager manager.TokenManager) *Handler {
	return &Handler{
		service:        service,
		tokenManager:   tokenManager,
		Auth:           auth.NewAuthHandler(service, tokenManager),
		User:           user.NewUserHandler(service),
		Workspace:      workspace.NewWorkspaceHandler(service, tokenManager),
		WorkspaceQueue: workspace_queue.NewWorkspaceQueueHandler(service),
	}
}
