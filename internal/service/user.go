package service

import (
	"github.com/ndt080/schedule-manager-backend/internal/domain"
	"github.com/ndt080/schedule-manager-backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository repository.User
}

func NewUserService(repository repository.User) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (service *UserService) ConfirmUserVerification(userId int64) error {
	return service.repository.ConfirmUserVerification(userId)
}

func (service *UserService) HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func (service *UserService) CreateUser(user *domain.User) (*domain.User, error) {
	user.PasswordHash = service.HashPassword(user.PasswordHash)
	user.IsVerified = false
	return service.repository.CreateUser(user)
}

func (service *UserService) GetUserByEmail(username string) (*domain.User, error) {
	return service.repository.GetUserByEmail(username)
}

func (service *UserService) GetUserById(id int64) (*domain.User, error) {
	user, err := service.repository.GetUserById(id)
	if err != nil {
		return nil, err
	}

	user.PasswordHash = ""
	return user, nil
}

func (service *UserService) GetUsersById(id []int64) (*[]domain.User, error) {
	users, err := service.repository.GetUsersById(id)
	if err != nil {
		return nil, err
	}

	for _, user := range *users {
		user.PasswordHash = ""
	}

	return users, nil
}

func (service *UserService) CheckExistsUser(username string) (bool, error) {
	user, err := service.repository.GetUserByEmail(username)
	return user != nil && err != nil, err
}
