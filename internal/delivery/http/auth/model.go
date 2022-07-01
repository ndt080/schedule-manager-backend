package auth

import "github.com/ndt080/schedule-manager-backend/internal/domain"

type Credential struct {
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,gte=8"`
}

type UserRequest struct {
	Username string `json:"username" binding:"required,max=64"`
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,gte=8"`
}

type UserResponse struct {
	User   domain.User   `json:"user" binding:"required"`
	Tokens domain.Tokens `json:"tokens" binding:"omitempty"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

type RefreshTokenResponse struct {
	Tokens domain.Tokens `json:"tokens" binding:"required"`
}
