package domain

type User struct {
	ID           int64  `db:"id" json:"id" binding:"omitempty"`
	Username     string `db:"username" json:"username" binding:"required"`
	Image        string `db:"image" json:"image" binding:"omitempty"`
	Email        string `db:"email" json:"email" binding:"required,email"`
	Status       string `db:"status" json:"status" binding:"required"`
	IsVerified   bool   `db:"is_verified" json:"isVerified" binding:"omitempty"`
	PasswordHash string `db:"password_hash" json:"passwordHash" binding:"required,gte=8"`
}
