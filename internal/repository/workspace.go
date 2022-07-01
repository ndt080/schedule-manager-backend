package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ndt080/schedule-manager-backend/internal/domain"
)

type WorkspaceRepository struct {
	database *sqlx.DB
}

func NewWorkspaceRepository(database *sqlx.DB) *WorkspaceRepository {
	return &WorkspaceRepository{database: database}
}

func (repository *WorkspaceRepository) Create(workspace *domain.Workspace) (*domain.Workspace, error) {
	if workspace.Image == "" {
		workspace.Image = "https://i.ibb.co/JmYr3ys/5.png"
	}

	query := `INSERT INTO workspace(name, description, image) 
			  VALUES ($1, $2, $3) 
			  RETURNING id`

	var id int64
	err := repository.database.QueryRow(query, workspace.Name, workspace.Description, workspace.Image).Scan(&id)
	if err != nil {
		return nil, err
	}

	return repository.Get(id)
}

func (repository *WorkspaceRepository) Update(workspace *domain.Workspace) error {
	query := `UPDATE workspace 
			  SET name=$1, description=$2, image=$3
			  WHERE id=$4;`
	_, err := repository.database.Exec(query, workspace.Name, workspace.Description, workspace.Image, workspace.ID)
	return err
}

func (repository *WorkspaceRepository) Get(id int64) (*domain.Workspace, error) {
	workspace := domain.Workspace{}
	query := `SELECT * FROM workspace WHERE id=$1`
	err := repository.database.Get(&workspace, query, id)
	return &workspace, err
}

func (repository *WorkspaceRepository) SearchByName(name string) (*[]domain.Workspace, error) {
	var workspaces []domain.Workspace
	query := `SELECT * FROM workspace WHERE lower(name) LIKE '%' || lower($1) || '%'`
	err := repository.database.Select(&workspaces, query, name)
	return &workspaces, err
}

func (repository *WorkspaceRepository) GetAll() (*[]domain.Workspace, error) {
	var workspaces []domain.Workspace
	query := `SELECT * FROM workspace`
	err := repository.database.Select(&workspaces, query)
	return &workspaces, err
}

func (repository *WorkspaceRepository) Remove(id int64) error {
	query := `DELETE FROM workspace WHERE id=$1;`
	_, err := repository.database.Exec(query, id)
	return err
}
