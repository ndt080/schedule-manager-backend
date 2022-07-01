package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ndt080/schedule-manager-backend/internal/domain"
	"github.com/ndt080/schedule-manager-backend/internal/mappers"
	"github.com/ndt080/schedule-manager-backend/internal/repository/models"
)

type WorkspaceMemberRepository struct {
	database *sqlx.DB
}

func NewWorkspaceMemberRepository(database *sqlx.DB) *WorkspaceMemberRepository {
	return &WorkspaceMemberRepository{database: database}
}

func (repository *WorkspaceMemberRepository) Create(member *domain.WorkspaceMember) (*domain.WorkspaceMember, error) {
	query := `INSERT INTO workspace_member(workspace, member, status) 
			  VALUES ($1, $2, $3)
			  RETURNING id`

	var id int64
	err := repository.database.QueryRow(query, member.WorkspaceId, member.Member.ID, member.Status).Scan(&id)
	if err != nil {
		return nil, err
	}

	return repository.Get(id)
}

func (repository *WorkspaceMemberRepository) UpdateStatus(member int64, status domain.WorkspaceMemberStatus) error {
	query := `UPDATE workspace_member 
			  SET status=$1
			  WHERE member=$2;`
	_, err := repository.database.Exec(query, status, member)
	return err
}

func (repository *WorkspaceMemberRepository) Get(id int64) (*domain.WorkspaceMember, error) {
	data := models.WorkspaceMemberDB{}
	query := `SELECT 
				u.id as member_id,
				u.email as member_email,
				u.username as member_username,
				u.image as member_image,
       			wm.id as id,
       			wm.workspace as workspace,
       			wm.status as status
			  FROM workspace_member as wm
			  INNER JOIN users u on u.id = wm.member
			  WHERE wm.id=$1`
	err := repository.database.Get(&data, query, id)
	if err != nil {
		return nil, err
	}

	member := mappers.MapRowToWorkspaceMember(data)
	return &member, nil
}

func (repository *WorkspaceMemberRepository) GetAll() (*[]domain.WorkspaceMember, error) {
	var rows []models.WorkspaceMemberDB
	query := `SELECT 
				u.id as member_id,
				u.email as member_email,
				u.username as member_username,
				u.image as member_image,
       			wm.id as id,
       			wm.workspace as workspace,
       			wm.status as status
			  FROM workspace_member as wm
			  INNER JOIN users u on u.id = wm.member`
	err := repository.database.Select(&rows, query)
	if err != nil {
		return nil, err
	}

	members := mappers.MapRowsToWorkspaceMembers(rows)
	return &members, nil
}

func (repository *WorkspaceMemberRepository) GetAllByWorkspace(id int64) (*[]domain.WorkspaceMember, error) {
	var rows []models.WorkspaceMemberDB
	query := `SELECT 
				u.id as member_id,
				u.email as member_email,
				u.username as member_username,
				u.image as member_image,
       			wm.id as id,
       			wm.workspace as workspace,
       			wm.status as status
			  FROM workspace_member as wm
			  INNER JOIN users u on u.id = wm.member
			  WHERE wm.workspace=$1`
	err := repository.database.Select(&rows, query, id)
	if err != nil {
		return nil, err
	}

	members := mappers.MapRowsToWorkspaceMembers(rows)
	return &members, nil
}

func (repository *WorkspaceMemberRepository) GetAllByUser(id int64) (*[]domain.WorkspaceMember, error) {
	var rows []models.WorkspaceMemberDB
	query := `SELECT 
				u.id as member_id,
				u.email as member_email,
				u.username as member_username,
				u.image as member_image,
       			wm.id as id,
       			wm.workspace as workspace,
       			wm.status as status
			  FROM workspace_member as wm
			  INNER JOIN users u on u.id = wm.member
			  WHERE wm.member=$1`
	err := repository.database.Select(&rows, query, id)
	if err != nil {
		return nil, err
	}

	members := mappers.MapRowsToWorkspaceMembers(rows)
	return &members, nil
}

func (repository *WorkspaceMemberRepository) Remove(id int64) error {
	query := `DELETE FROM workspace_member WHERE id=$1;`
	_, err := repository.database.Exec(query, id)
	return err
}
