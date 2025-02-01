package services

import (
	"github.com/mostafizur-raahman/go-with-gone/internal/models"
	"github.com/mostafizur-raahman/go-with-gone/internal/repositories"
)

type UserService interface {
	CreateUser(user models.User) (int, error)
	AllUsers() ([]models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (us *userService) CreateUser(user models.User) (int, error) {
	return us.repo.Create(user)
}

func (us *userService) AllUsers() ([]models.User, error) {
	return us.repo.GetAll()
}
