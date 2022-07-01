package error

type ServerErrorMessage = string

const (
	UnauthorizedMessage        ServerErrorMessage = "Unauthorized"
	InvalidCredentialsMessage  ServerErrorMessage = "There is no user with such credentials"
	InvalidAccessTokenMessage  ServerErrorMessage = "The access token is invalid"
	InvalidRefreshTokenMessage ServerErrorMessage = "The refresh token is invalid"
	CredentialsExistsMessage   ServerErrorMessage = "A user with such credentials already exists"
)
