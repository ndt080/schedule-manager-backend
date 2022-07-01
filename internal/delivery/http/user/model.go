package user

import "github.com/ndt080/schedule-manager-backend/internal/domain"

type UsersDataRequest struct {
	IDs []int64 `json:"ids" binding:"required"`
}

type UsersDataResponse struct {
	Users []domain.User `json:"users" binding:"required"`
}
