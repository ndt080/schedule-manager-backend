package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ndt080/schedule-manager-backend/internal/domain"
	"time"
)

type User interface {
	CreateUser(user *domain.User) (*domain.User, error)
	GetUserByEmail(username string) (*domain.User, error)
	GetUserById(id int64) (*domain.User, error)
	GetUsersById(id []int64) (*[]domain.User, error)
	ConfirmUserVerification(id int64) error
}

type Workspace interface {
	Create(workspace *domain.Workspace) (*domain.Workspace, error)
	Update(workspace *domain.Workspace) error
	Remove(id int64) error
	Get(id int64) (*domain.Workspace, error)
	GetAll() (*[]domain.Workspace, error)
	SearchByName(name string) (*[]domain.Workspace, error)
}

type WorkspaceMember interface {
	Create(member *domain.WorkspaceMember) (*domain.WorkspaceMember, error)
	UpdateStatus(member int64, status domain.WorkspaceMemberStatus) error
	Remove(id int64) error
	Get(id int64) (*domain.WorkspaceMember, error)
	GetAll() (*[]domain.WorkspaceMember, error)
	GetAllByWorkspace(id int64) (*[]domain.WorkspaceMember, error)
	GetAllByUser(id int64) (*[]domain.WorkspaceMember, error)
}

type WorkspaceTask interface {
	Create(task *domain.WorkspaceTask) (*domain.WorkspaceTask, error)
	UpdateName(id int64, name string) error
	UpdateDescription(id int64, description string) error
	Remove(id int64) error
	Get(id int64) (*domain.WorkspaceTask, error)
	GetAll() (*[]domain.WorkspaceTask, error)
	GetAllByWorkspace(id int64) (*[]domain.WorkspaceTask, error)
}

type WorkspaceQueue interface {
	Create(queue *domain.WorkspaceQueue) (*domain.WorkspaceQueue, error)
	Remove(id int64) error
	Get(id int64) (*domain.WorkspaceQueue, error)
	GetAllWithoutMembers() (*[]domain.WorkspaceQueue, error)
	GetAllByWorkspaceWithoutMembers(id int64) ([]domain.WorkspaceQueue, error)

	GetMembers(id int64) (*[]int64, error)
	AddMember(queueId int64, userId int64) (int64, error)
	RemoveMember(id int64) error
	RemoveMemberByUser(uid int64, qid int64) error
}

type WorkspaceSchedule interface {
	Create(schedule *domain.WorkspaceSchedule) (*domain.WorkspaceSchedule, error)
	Remove(id int64) error
	Get(id int64) (*domain.WorkspaceSchedule, error)
	GetByStartDate(date time.Time) (*domain.WorkspaceSchedule, error)
	GetAllWithoutRecords() (*[]domain.WorkspaceSchedule, error)
	GetAllByWorkspaceWithoutRecords(id int64) (*[]domain.WorkspaceSchedule, error)

	AddRecord(record *domain.WorkspaceScheduleRecord) (*domain.WorkspaceScheduleRecord, error)
	UpdateRecord(record domain.WorkspaceScheduleRecord) error
	RemoveRecord(id int64) error
	GetRecord(id int64) (*domain.WorkspaceScheduleRecord, error)
	GetRecords(scheduleId int64) (*[]domain.WorkspaceScheduleRecord, error)
}

type Repository struct {
	User              User
	Workspace         Workspace
	WorkspaceMember   WorkspaceMember
	WorkspaceTask     WorkspaceTask
	WorkspaceQueue    WorkspaceQueue
	WorkspaceSchedule WorkspaceSchedule
}

func NewRepository(database *sqlx.DB) *Repository {
	return &Repository{
		User:              NewUserRepository(database),
		Workspace:         NewWorkspaceRepository(database),
		WorkspaceMember:   NewWorkspaceMemberRepository(database),
		WorkspaceTask:     NewWorkspaceTaskRepository(database),
		WorkspaceQueue:    NewWorkspaceQueueRepository(database),
		WorkspaceSchedule: NewWorkspaceScheduleRepository(database),
	}
}
