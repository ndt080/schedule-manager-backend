package workspace

import "github.com/ndt080/schedule-manager-backend/internal/domain"

type WorkspaceRequest struct {
	ID          int64  `db:"id" json:"id" binding:"omitempty"`
	Name        string `db:"name" json:"name" binding:"required"`
	Description string `db:"description" json:"description" binding:"omitempty"`
	Image       string `db:"image" json:"image" binding:"omitempty"`
}

type WorkspaceTaskRequest struct {
	ID          int64  `db:"id" json:"id" binding:"omitempty"`
	Name        string `db:"name" json:"name" binding:"required"`
	Description string `db:"description" json:"description" binding:"omitempty"`
}

type WorkspaceMemberRequest struct {
	Email  string                       `db:"email" json:"email" binding:"required"`
	Status domain.WorkspaceMemberStatus `db:"status" json:"status" binding:"required"`
}

type WorkspaceInviteResponse struct {
	Token string `db:"token" json:"token" binding:"required"`
}
