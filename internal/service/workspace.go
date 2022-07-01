package service

import (
	"github.com/ndt080/schedule-manager-backend/internal/domain"
	"github.com/ndt080/schedule-manager-backend/internal/repository"
)

type WorkspaceService struct {
	repository *repository.Repository
}

func NewWorkspaceService(repository *repository.Repository) *WorkspaceService {
	return &WorkspaceService{
		repository: repository,
	}
}

func (service *WorkspaceService) SearchWorkspaces(query string) (*[]domain.WorkspaceData, error) {
	workspaces, err := service.repository.Workspace.SearchByName(query)
	if err != nil {
		return nil, err
	}

	var workspacesData []domain.WorkspaceData

	for _, workspace := range *workspaces {
		members, err := service.repository.WorkspaceMember.GetAllByWorkspace(workspace.ID)
		if err != nil {
			return nil, err
		}

		workspacesData = append(workspacesData, domain.WorkspaceData{
			Workspace: workspace,
			Members:   *members,
		})
	}
	return &workspacesData, nil
}

func (service *WorkspaceService) GetWorkspaceById(id int64) (*domain.WorkspaceData, error) {
	workspace, err := service.repository.Workspace.Get(id)
	if err != nil {
		return nil, err
	}

	members, err := service.GetWorkspaceMembers(workspace.ID)
	if err != nil {
		return nil, err
	}

	tasks, err := service.GetWorkspaceTasks(workspace.ID)
	if err != nil {
		return nil, err
	}

	schedules, err := service.GetWorkspaceSchedules(workspace.ID)
	if err != nil {
		return nil, err
	}

	queues, err := service.GetWorkspaceQueues(workspace.ID)
	if err != nil {
		return nil, err
	}

	data := &domain.WorkspaceData{
		Workspace: *workspace,
		Members:   *members,
		Tasks:     *tasks,
		Schedules: *schedules,
		Queues:    *queues,
	}
	return data, nil
}

func (service *WorkspaceService) GetWorkspacesByUser(userId int64) (*[]domain.WorkspaceData, error) {
	memberships, err := service.GetWorkspaceMembersByUser(userId)
	if err != nil {
		return nil, err
	}

	var workspaces []domain.WorkspaceData

	for _, membership := range *memberships {
		workspace, err := service.repository.Workspace.Get(membership.WorkspaceId)
		if err != nil {
			return nil, err
		}

		members, err := service.repository.WorkspaceMember.GetAllByWorkspace(workspace.ID)
		if err != nil {
			return nil, err
		}

		workspaces = append(workspaces, domain.WorkspaceData{
			Workspace: *workspace,
			Members:   *members,
		})
	}
	return &workspaces, nil
}

func (service *WorkspaceService) GetWorkspaceMembers(workspaceId int64) (*[]domain.WorkspaceMember, error) {
	return service.repository.WorkspaceMember.GetAllByWorkspace(workspaceId)
}

func (service *WorkspaceService) GetWorkspaceMembersByUser(userId int64) (*[]domain.WorkspaceMember, error) {
	return service.repository.WorkspaceMember.GetAllByUser(userId)
}

func (service *WorkspaceService) GetWorkspaceTasks(workspaceId int64) (*[]domain.WorkspaceTask, error) {
	return service.repository.WorkspaceTask.GetAllByWorkspace(workspaceId)
}

func (service *WorkspaceService) GetWorkspaceQueues(workspaceId int64) (*[]domain.WorkspaceQueue, error) {
	queues, err := service.repository.WorkspaceQueue.GetAllByWorkspaceWithoutMembers(workspaceId)
	if err != nil {
		return nil, err
	}

	for index, queue := range queues {
		members, err := service.repository.WorkspaceQueue.GetMembers(queue.ID)
		if err != nil {
			return nil, err
		}

		queues[index].Members = *members
	}

	return &queues, nil
}

func (service *WorkspaceService) GetWorkspaceSchedules(workspaceId int64) (*[]domain.WorkspaceSchedule, error) {
	schedules, err := service.repository.WorkspaceSchedule.GetAllByWorkspaceWithoutRecords(workspaceId)
	if err != nil {
		return nil, err
	}

	for _, schedule := range *schedules {
		records, err := service.repository.WorkspaceSchedule.GetRecords(schedule.ID)
		if err != nil {
			return nil, err
		}

		schedule.Records = *records
	}
	return schedules, nil
}

func (service *WorkspaceService) RemoveWorkspace(id int64) error {
	return service.repository.Workspace.Remove(id)
}

func (service *WorkspaceService) CreateWorkspace(user domain.User, workspace domain.Workspace) (*domain.WorkspaceData, error) {
	newWorkspace, err := service.repository.Workspace.Create(&workspace)
	if err != nil {
		return nil, err
	}

	owner, err := service.repository.WorkspaceMember.Create(&domain.WorkspaceMember{
		Member:      user,
		Status:      domain.Owner,
		WorkspaceId: newWorkspace.ID,
	})
	if err != nil {
		return nil, err
	}

	var members []domain.WorkspaceMember
	members = append(members, *owner)

	data := domain.WorkspaceData{
		Workspace: *newWorkspace,
		Members:   members,
	}
	return &data, nil
}

func (service *WorkspaceService) AddMemberToWorkspace(member domain.WorkspaceMember) (*domain.WorkspaceMember, error) {
	return service.repository.WorkspaceMember.Create(&member)
}

func (service *WorkspaceService) UpdateWorkspace(workspace domain.Workspace) (*domain.WorkspaceData, error) {
	err := service.repository.Workspace.Update(&workspace)
	if err != nil {
		return nil, err
	}

	return service.GetWorkspaceById(workspace.ID)
}

func (service *WorkspaceService) CreateWorkspaceTask(task domain.WorkspaceTask) (*domain.WorkspaceTask, error) {
	return service.repository.WorkspaceTask.Create(&task)
}

func (service *WorkspaceService) RemoveWorkspaceTask(id int64) error {
	return service.repository.WorkspaceTask.Remove(id)
}

func (service *WorkspaceService) CreateWorkspaceQueue(queue domain.WorkspaceQueue) (*domain.WorkspaceQueue, error) {
	return service.repository.WorkspaceQueue.Create(&queue)
}

func (service *WorkspaceService) RemoveWorkspaceQueue(id int64) error {
	return service.repository.WorkspaceQueue.Remove(id)
}

func (service *WorkspaceService) JoinToWorkspaceQueue(user *domain.User, id int64) (*domain.WorkspaceQueue, error) {
	_, err := service.repository.WorkspaceQueue.AddMember(id, user.ID)
	if err != nil {
		return nil, err
	}

	return service.repository.WorkspaceQueue.Get(id)
}

func (service *WorkspaceService) LeaveWorkspaceQueue(uid int64, qid int64) (*domain.WorkspaceQueue, error) {
	if err := service.repository.WorkspaceQueue.RemoveMemberByUser(uid, qid); err != nil {
		return nil, err
	}

	return service.repository.WorkspaceQueue.Get(qid)
}
