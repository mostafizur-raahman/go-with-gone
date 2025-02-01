package repositories

import (
	"errors"
	"sync"

	"github.com/mostafizur-raahman/go-with-gone/internal/models"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	// GetById(id int) (models.User, error)
	Create(user models.User) (int, error)
	// Update(id int, user models.User) error
	// Delete(id int) error
}

// TODO: we use postgreSQL
type userRepository struct {
	users map[int]models.User
	mu    sync.Mutex
}

func NewUserRepository() UserRepository {
	return &userRepository{
		users: make(map[int]models.User),
	}
}

func (r *userRepository) Create(user models.User) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := len(r.users) + 1
	user.ID = id
	r.users[id] = user

	return id, nil
}

func (r *userRepository) GetAll() ([]models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var userList []models.User
	for _, user := range r.users {
		userList = append(userList, user)
	}

	if len(userList) == 0 {
		return nil, errors.New("no users found")
	}

	return userList, nil
}
