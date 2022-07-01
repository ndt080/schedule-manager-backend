package service

import (
	"github.com/gin-gonic/gin"
	"github.com/ndt080/schedule-manager-backend/internal/delivery/smtp"
	"github.com/ndt080/schedule-manager-backend/internal/domain"
	"github.com/ndt080/schedule-manager-backend/internal/repository"
	"github.com/ndt080/schedule-manager-backend/pkg/auth"
)

type Authorization interface {
	IdentifyUser(context *gin.Context)
	CreateTokens(userId int64) (*domain.Tokens, error)
	SendVerificationEmail(user *domain.User, context *gin.Context) (bool, error)
	GetAuthorizedUser(context *gin.Context) (*domain.User, error)
}

type User interface {
	HashPassword(password string) string
	CreateUser(user *domain.User) (*domain.User, error)
	GetUserById(id int64) (*domain.User, error)
	GetUsersById(id []int64) (*[]domain.User, error)
	GetUserByEmail(username string) (*domain.User, error)
	CheckExistsUser(username string) (bool, error)
	ConfirmUserVerification(id int64) error
}

type Workspace interface {
	GetWorkspaceById(id int64) (*domain.WorkspaceData, error)
	GetWorkspacesByUser(userId int64) (*[]domain.WorkspaceData, error)
	GetWorkspaceMembers(workspaceId int64) (*[]domain.WorkspaceMember, error)
	GetWorkspaceMembersByUser(userId int64) (*[]domain.WorkspaceMember, error)
	GetWorkspaceTasks(workspaceId int64) (*[]domain.WorkspaceTask, error)
	GetWorkspaceQueues(workspaceId int64) (*[]domain.WorkspaceQueue, error)
	GetWorkspaceSchedules(workspaceId int64) (*[]domain.WorkspaceSchedule, error)
	CreateWorkspace(user domain.User, workspace domain.Workspace) (*domain.WorkspaceData, error)
	RemoveWorkspace(id int64) error
	SearchWorkspaces(query string) (*[]domain.WorkspaceData, error)
	UpdateWorkspace(workspace domain.Workspace) (*domain.WorkspaceData, error)

	CreateWorkspaceTask(task domain.WorkspaceTask) (*domain.WorkspaceTask, error)
	RemoveWorkspaceTask(id int64) error

	CreateWorkspaceQueue(data domain.WorkspaceQueue) (*domain.WorkspaceQueue, error)
	RemoveWorkspaceQueue(id int64) error
	JoinToWorkspaceQueue(user *domain.User, id int64) (*domain.WorkspaceQueue, error)
	LeaveWorkspaceQueue(uid int64, id int64) (*domain.WorkspaceQueue, error)

	AddMemberToWorkspace(member domain.WorkspaceMember) (*domain.WorkspaceMember, error)
}

type Service struct {
	Authorization Authorization
	User          User
	Workspace     Workspace
}

func NewService(repository *repository.Repository, smtpService *smtp.SmtpService, manager auth.TokenManager) *Service {
	return &Service{
		Authorization: NewAuthService(manager, smtpService, repository.User),
		User:          NewUserService(repository.User),
		Workspace:     NewWorkspaceService(repository),
	}
}
