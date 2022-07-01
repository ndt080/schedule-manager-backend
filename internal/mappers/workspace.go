package mappers

import (
	"github.com/ndt080/schedule-manager-backend/internal/domain"
	"github.com/ndt080/schedule-manager-backend/internal/repository/models"
)

func MapRowsToWorkspaceMembers(rows []models.WorkspaceMemberDB) []domain.WorkspaceMember {
	var members []domain.WorkspaceMember
	for _, row := range rows {
		members = append(members, domain.WorkspaceMember{
			ID: row.ID,
			Member: domain.User{
				ID:         row.MemberId,
				Email:      row.MemberEmail,
				Username:   row.MemberUsername,
				Image:      row.MemberImage,
				IsVerified: true,
			},
			Status:      row.Status,
			WorkspaceId: row.Workspace,
		})
	}
	return members
}

func MapRowToWorkspaceMember(row models.WorkspaceMemberDB) domain.WorkspaceMember {
	return domain.WorkspaceMember{
		ID: row.ID,
		Member: domain.User{
			ID:         row.MemberId,
			Email:      row.MemberEmail,
			Username:   row.MemberUsername,
			Image:      row.MemberImage,
			IsVerified: true,
		},
		Status:      row.Status,
		WorkspaceId: row.Workspace,
	}
}
