package domain

import (
	"time"
)

type WorkspaceMemberStatus string

const (
	Member WorkspaceMemberStatus = "member"
	Owner  WorkspaceMemberStatus = "owner"
	Editor WorkspaceMemberStatus = "editor"
)

type WorkspaceData struct {
	Workspace Workspace           `json:"workspace" binding:"required"`
	Members   []WorkspaceMember   `json:"members" binding:"omitempty"`
	Tasks     []WorkspaceTask     `json:"tasks" binding:"omitempty"`
	Schedules []WorkspaceSchedule `json:"schedules" binding:"omitempty"`
	Queues    []WorkspaceQueue    `json:"queues" binding:"omitempty"`
}

type Workspace struct {
	ID          int64     `db:"id" json:"id" binding:"omitempty"`
	Name        string    `db:"name" json:"name" binding:"required"`
	Description string    `db:"description" json:"description" binding:"omitempty"`
	Image       string    `db:"image" json:"image" binding:"omitempty"`
	CreatedAt   time.Time `db:"created_at" json:"createdAt" binding:"omitempty"`
}

type WorkspaceMember struct {
	ID          int64                 `db:"id" json:"id" binding:"omitempty"`
	Member      User                  `db:"ws_user" json:"member" binding:"required"`
	Status      WorkspaceMemberStatus `db:"status" json:"status" binding:"required"`
	WorkspaceId int64                 `db:"workspace" json:"workspaceId" binding:"required"`
}

type WorkspaceTask struct {
	ID          int64  `db:"id" json:"id" binding:"omitempty"`
	Name        string `db:"name" json:"name" binding:"required"`
	Description string `db:"description" json:"description" binding:"omitempty"`
	Creator     int64  `db:"creator" json:"creator" binding:"omitempty"`
	WorkspaceId int64  `db:"workspace" json:"workspaceId" binding:"required"`
}

type WorkspaceSchedule struct {
	ID          int64                     `db:"id" json:"id" binding:"omitempty"`
	StartDate   time.Time                 `db:"start" json:"startDate" binding:"required"`
	Records     []WorkspaceScheduleRecord `json:"records" binding:"required"`
	WorkspaceId int64                     `db:"workspace" json:"workspaceId" binding:"required"`
}

type WorkspaceScheduleRecord struct {
	ID            int64     `db:"id" json:"id" binding:"omitempty"`
	Description   string    `db:"description" json:"description" binding:"omitempty"`
	StartDateTime time.Time `db:"start_datetime" json:"startDateTime" binding:"required"`
	EndDateTime   time.Time `db:"end_datetime" json:"endDateTime" binding:"required"`
	TaskId        int64     `db:"task" json:"taskId" binding:"required"`
	ScheduleId    int64     `db:"schedule" json:"scheduleId" binding:"required"`
}

type WorkspaceQueue struct {
	ID          int64   `db:"id" json:"id" binding:"omitempty"`
	Name        string  `db:"name" json:"name" binding:"required"`
	Members     []int64 `json:"members" binding:"omitempty"`
	WorkspaceId int64   `db:"workspace" json:"workspaceId" binding:"required"`
}
