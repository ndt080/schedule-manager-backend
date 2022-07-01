package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/ndt080/schedule-manager-backend/internal/domain"
	"strings"
)

type UserRepository struct {
	database *sqlx.DB
}

func NewUserRepository(database *sqlx.DB) *UserRepository {
	return &UserRepository{database: database}
}

func (repository *UserRepository) CreateUser(user *domain.User) (*domain.User, error) {
	query := `INSERT INTO users(email, password_hash, username) VALUES ($1, $2, $3);`
	_, err := repository.database.Exec(query, user.Email, user.PasswordHash, user.Username)
	if err != nil {
		return nil, err
	}
	user, err = repository.GetUserByEmail(user.Email)
	user.PasswordHash = ""
	return user, err
}

func (repository *UserRepository) GetUserByEmail(username string) (*domain.User, error) {
	user := domain.User{}
	query := `SELECT * FROM users WHERE email=$1 `
	err := repository.database.Get(&user, query, username)
	return &user, err
}

func (repository *UserRepository) ConfirmUserVerification(userId int64) error {
	query := `UPDATE users SET is_verified=true WHERE id=$1;`
	_, err := repository.database.Exec(query, userId)
	return err
}

func (repository *UserRepository) GetUserById(id int64) (*domain.User, error) {
	user := domain.User{}
	query := `SELECT * FROM users WHERE id=$1 `
	err := repository.database.Get(&user, query, id)
	return &user, err
}

func (repository *UserRepository) GetUsersById(ids []int64) (*[]domain.User, error) {
	var IDs []string
	for _, i := range ids {
		IDs = append(IDs, fmt.Sprint(i))
	}

	var users []domain.User
	query := `SELECT * FROM users WHERE id IN ($1)`
	err := repository.database.Select(&users, query, strings.Join(IDs, ", "))
	return &users, err
}
