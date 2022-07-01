package response

import (
	error2 "github.com/ndt080/schedule-manager-backend/internal/delivery/http/response/error"
	s "github.com/ndt080/schedule-manager-backend/internal/delivery/http/response/success"
)

func NewServerSuccessResponse(msg string) *s.ServerSuccessResponse {
	return &s.ServerSuccessResponse{
		Success: true,
		Msg:     msg,
	}
}

func NewServerInternalError(err string) *error2.ServerErrorResponse {
	return &error2.ServerErrorResponse{
		Success: false,
		Error:   err,
	}
}

func NewServerBadRequestError(err string) *error2.ServerErrorResponse {
	return &error2.ServerErrorResponse{
		Success:   false,
		ErrorCode: error2.BadRequestCode,
		Error:     err,
	}
}

func NewServerCredentialsExistsError() *error2.ServerErrorResponse {
	return &error2.ServerErrorResponse{
		Success:   false,
		ErrorCode: error2.CredentialsExistsCode,
		Error:     error2.CredentialsExistsMessage,
	}
}

func NewServerInvalidCredentialsError() *error2.ServerErrorResponse {
	return &error2.ServerErrorResponse{
		Success:   false,
		ErrorCode: error2.InvalidCredentialsCode,
		Error:     error2.InvalidCredentialsMessage,
	}
}

func NewServerUnauthorizedError(err string) *error2.ServerErrorResponse {
	return &error2.ServerErrorResponse{
		Success:   false,
		ErrorCode: error2.UnauthorizedCode,
		Error:     err,
	}
}

func NewServerInvalidRefreshTokenError() *error2.ServerErrorResponse {
	return &error2.ServerErrorResponse{
		Success:   false,
		ErrorCode: error2.InvalidRefreshTokenCode,
		Error:     error2.InvalidRefreshTokenMessage,
	}
}

func NewServerInvalidAccessTokenError() *error2.ServerErrorResponse {
	return &error2.ServerErrorResponse{
		Success:   false,
		ErrorCode: error2.InvalidAccessTokenCode,
		Error:     error2.InvalidAccessTokenMessage,
	}
}
