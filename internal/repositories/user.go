package repositories

import (
	"sync"

	"github.com/mostafizur-raahman/go-with-gone/internal/models"
)

type UserRepositories interface {
	// GetAll() ([]models.User, error)
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

func NewUserRepository() UserRepositories {
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
