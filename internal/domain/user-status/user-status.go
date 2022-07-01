package user_status

type UserStatus string

const (
	User      UserStatus = "user"
	Admin     UserStatus = "admin"
	Moderator UserStatus = "moderator"
)
