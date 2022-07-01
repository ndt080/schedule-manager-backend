package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ndt080/schedule-manager-backend/internal/domain"
	"time"
)

type WorkspaceScheduleRepository struct {
	database *sqlx.DB
}

func NewWorkspaceScheduleRepository(database *sqlx.DB) *WorkspaceScheduleRepository {
	return &WorkspaceScheduleRepository{database: database}
}

func (repository *WorkspaceScheduleRepository) Create(schedule *domain.WorkspaceSchedule) (*domain.WorkspaceSchedule, error) {
	query := `INSERT INTO workspace_schedule(start, workspace) 
			  VALUES ($1, $2)
			  RETURNING id`

	var id int64
	err := repository.database.QueryRow(query, schedule.StartDate, schedule.WorkspaceId).Scan(&id)
	if err != nil {
		return nil, err
	}

	return repository.Get(id)
}

func (repository *WorkspaceScheduleRepository) AddRecord(record *domain.WorkspaceScheduleRecord) (*domain.WorkspaceScheduleRecord, error) {
	query := `INSERT INTO workspace_schedule_record(schedule, description, start_datetime, end_datetime, task) 
			  VALUES ($1, $2)
			  RETURNING id`

	var id int64
	err := repository.database.QueryRow(
		query,
		record.ScheduleId,
		record.Description,
		record.StartDateTime,
		record.EndDateTime,
		record.TaskId,
	).Scan(&id)
	if err != nil {
		return nil, err
	}

	return repository.GetRecord(id)
}

func (repository *WorkspaceScheduleRepository) GetRecord(id int64) (*domain.WorkspaceScheduleRecord, error) {
	var record domain.WorkspaceScheduleRecord
	query := `SELECT * FROM workspace_schedule_record WHERE id=$1`
	err := repository.database.Get(&record, query, id)
	return &record, err
}

func (repository *WorkspaceScheduleRepository) UpdateRecord(record domain.WorkspaceScheduleRecord) error {
	query := `UPDATE workspace_schedule_record
			  SET description=$1, start_datetime=$2, end_datetime=$3, task=$4
			  WHERE schedule=$5;`
	_, err := repository.database.Exec(query, record.Description, record.StartDateTime, record.EndDateTime, record.TaskId, record.ScheduleId)
	return err
}

func (repository *WorkspaceScheduleRepository) Get(id int64) (*domain.WorkspaceSchedule, error) {
	schedule := domain.WorkspaceSchedule{}
	sQuery := `SELECT * FROM workspace_schedule WHERE id=$1`
	err := repository.database.Get(&schedule, sQuery, id)
	if err != nil {
		return nil, err
	}

	var records []domain.WorkspaceScheduleRecord
	rQuery := `SELECT * FROM workspace_schedule_record WHERE schedule=$1`
	err = repository.database.Select(&records, rQuery, schedule.ID)
	if err != nil {
		return nil, err
	}

	schedule.Records = records
	return &schedule, nil
}

func (repository *WorkspaceScheduleRepository) GetByStartDate(date time.Time) (*domain.WorkspaceSchedule, error) {
	schedule := domain.WorkspaceSchedule{}
	sQuery := `SELECT * FROM workspace_schedule WHERE start::date=$1::date`
	err := repository.database.Get(&schedule, sQuery, date)
	if err != nil {
		return nil, err
	}

	var records []domain.WorkspaceScheduleRecord
	rQuery := `SELECT * FROM workspace_schedule_record WHERE schedule=$1`
	err = repository.database.Select(&records, rQuery, schedule.ID)
	if err != nil {
		return nil, err
	}

	schedule.Records = records
	return &schedule, nil
}

func (repository *WorkspaceScheduleRepository) GetAllWithoutRecords() (*[]domain.WorkspaceSchedule, error) {
	var schedules []domain.WorkspaceSchedule
	query := `SELECT * FROM workspace_schedule`
	err := repository.database.Select(&schedules, query)
	return &schedules, err
}

func (repository *WorkspaceScheduleRepository) GetAllByWorkspaceWithoutRecords(id int64) (*[]domain.WorkspaceSchedule, error) {
	var schedules []domain.WorkspaceSchedule
	query := `SELECT * FROM workspace_schedule WHERE workspace=$1`
	err := repository.database.Select(&schedules, query, id)
	return &schedules, err
}

func (repository *WorkspaceScheduleRepository) GetRecords(scheduleId int64) (*[]domain.WorkspaceScheduleRecord, error) {
	var records []domain.WorkspaceScheduleRecord
	query := `SELECT * FROM workspace_schedule_record WHERE schedule=$1`
	err := repository.database.Select(&records, query, scheduleId)
	return &records, err
}

func (repository *WorkspaceScheduleRepository) Remove(id int64) error {
	query := `DELETE FROM workspace_schedule WHERE id=$1;`
	_, err := repository.database.Exec(query, id)
	return err
}

func (repository *WorkspaceScheduleRepository) RemoveRecord(id int64) error {
	query := `DELETE FROM workspace_schedule_record WHERE id=$1;`
	_, err := repository.database.Exec(query, id)
	return err
}
