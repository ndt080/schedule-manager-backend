package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ndt080/schedule-manager-backend/internal/domain"
)

type WorkspaceTaskRepository struct {
	database *sqlx.DB
}

func NewWorkspaceTaskRepository(database *sqlx.DB) *WorkspaceTaskRepository {
	return &WorkspaceTaskRepository{database: database}
}

func (repository *WorkspaceTaskRepository) Create(task *domain.WorkspaceTask) (*domain.WorkspaceTask, error) {
	query := `INSERT INTO workspace_task(workspace, name, description, creator) 
			  VALUES ($1, $2, $3, $4)
			  RETURNING id`

	var id int64
	err := repository.database.QueryRow(query, task.WorkspaceId, task.Name, task.Description, task.Creator).Scan(&id)
	if err != nil {
		return nil, err
	}

	return repository.Get(id)
}

func (repository *WorkspaceTaskRepository) UpdateName(id int64, name string) error {
	query := `UPDATE workspace_task 
			  SET name=$1
			  WHERE id=$2;`
	_, err := repository.database.Exec(query, name, id)
	return err
}

func (repository *WorkspaceTaskRepository) UpdateDescription(id int64, description string) error {
	query := `UPDATE workspace_task 
			  SET description=$1
			  WHERE id=$2;`
	_, err := repository.database.Exec(query, description, id)
	return err
}

func (repository *WorkspaceTaskRepository) Get(id int64) (*domain.WorkspaceTask, error) {
	task := domain.WorkspaceTask{}
	query := `SELECT * FROM workspace_task WHERE id=$1`
	err := repository.database.Get(&task, query, id)
	return &task, err
}

func (repository *WorkspaceTaskRepository) GetAll() (*[]domain.WorkspaceTask, error) {
	var tasks []domain.WorkspaceTask
	query := `SELECT * FROM workspace_task`
	err := repository.database.Select(&tasks, query)
	return &tasks, err
}

func (repository *WorkspaceTaskRepository) GetAllByWorkspace(id int64) (*[]domain.WorkspaceTask, error) {
	var tasks []domain.WorkspaceTask
	query := `SELECT * FROM workspace_task WHERE workspace=$1`
	err := repository.database.Select(&tasks, query, id)
	return &tasks, err
}

func (repository *WorkspaceTaskRepository) Remove(id int64) error {
	query := `DELETE FROM workspace_task WHERE id=$1;`
	_, err := repository.database.Exec(query, id)
	return err
}
