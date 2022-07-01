package error

type ServerErrorCode = string

const (
	UnauthorizedCode        ServerErrorCode = "Unauthorized"
	InvalidCredentialsCode  ServerErrorCode = "InvalidCredentials"
	InvalidAccessTokenCode  ServerErrorCode = "InvalidAccessToken"
	InvalidRefreshTokenCode ServerErrorCode = "InvalidRefreshToken"
	CredentialsExistsCode   ServerErrorCode = "CredentialsExists"
	BadRequestCode          ServerErrorCode = "BadRequest"
)
