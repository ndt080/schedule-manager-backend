package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ndt080/schedule-manager-backend/internal/domain"
)

type WorkspaceQueueRepository struct {
	database *sqlx.DB
}

func NewWorkspaceQueueRepository(database *sqlx.DB) *WorkspaceQueueRepository {
	return &WorkspaceQueueRepository{database: database}
}

func (repository *WorkspaceQueueRepository) Create(queue *domain.WorkspaceQueue) (*domain.WorkspaceQueue, error) {
	query := `INSERT INTO workspace_queue(name, workspace) 
			  VALUES ($1, $2)
			  RETURNING id`

	var id int64
	err := repository.database.QueryRow(query, queue.Name, queue.WorkspaceId).Scan(&id)
	if err != nil {
		return nil, err
	}

	return repository.Get(id)
}

func (repository *WorkspaceQueueRepository) AddMember(queueId int64, userId int64) (int64, error) {
	query := `INSERT INTO workspace_queue_member(queue, member) 
			  VALUES ($1, $2)
			  RETURNING id`

	var id int64
	err := repository.database.QueryRow(query, queueId, userId).Scan(&id)
	return id, err
}

func (repository *WorkspaceQueueRepository) Get(id int64) (*domain.WorkspaceQueue, error) {
	queue := domain.WorkspaceQueue{}
	query := `SELECT * FROM workspace_queue WHERE id=$1`
	err := repository.database.Get(&queue, query, id)
	if err != nil {
		return nil, err
	}

	members, err := repository.GetMembers(id)
	queue.Members = *members
	return &queue, nil
}

func (repository *WorkspaceQueueRepository) GetMembers(id int64) (*[]int64, error) {
	var members []int64
	mQuery := `SELECT member FROM workspace_queue_member WHERE queue=$1`
	err := repository.database.Select(&members, mQuery, id)
	return &members, err
}

func (repository *WorkspaceQueueRepository) GetAllWithoutMembers() (*[]domain.WorkspaceQueue, error) {
	var queues []domain.WorkspaceQueue
	query := `SELECT * FROM workspace_queue`
	err := repository.database.Select(&queues, query)
	return &queues, err
}

func (repository *WorkspaceQueueRepository) GetAllByWorkspaceWithoutMembers(id int64) ([]domain.WorkspaceQueue, error) {
	var queues []domain.WorkspaceQueue
	query := `SELECT * FROM workspace_queue WHERE workspace=$1`
	err := repository.database.Select(&queues, query, id)
	return queues, err
}

func (repository *WorkspaceQueueRepository) Remove(id int64) error {
	query := `DELETE FROM workspace_queue WHERE id=$1;`
	_, err := repository.database.Exec(query, id)
	return err
}

func (repository *WorkspaceQueueRepository) RemoveMember(id int64) error {
	query := `DELETE FROM workspace_queue_member WHERE id=$1;`
	_, err := repository.database.Exec(query, id)
	return err
}

func (repository *WorkspaceQueueRepository) RemoveMemberByUser(uid int64, qid int64) error {
	query := `DELETE FROM workspace_queue_member WHERE queue=$1 AND member=$2;`
	_, err := repository.database.Exec(query, qid, uid)
	return err
}
