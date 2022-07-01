package models

import "github.com/ndt080/schedule-manager-backend/internal/domain"

type WorkspaceMemberDB struct {
	ID             int64                        `db:"id"`
	MemberId       int64                        `db:"member_id"`
	MemberEmail    string                       `db:"member_email"`
	MemberUsername string                       `db:"member_username"`
	MemberImage    string                       `db:"member_image"`
	Status         domain.WorkspaceMemberStatus `db:"status"`
	Workspace      int64                        `db:"workspace"`
}
