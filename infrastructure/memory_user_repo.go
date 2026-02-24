package infrastructure

import (
	"sync"

	"github.com/hmdyt/orange/domain"
)

type MemoryUserRepository struct {
	mu    sync.Mutex
	users map[string]*domain.User
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		users: make(map[string]*domain.User),
	}
}

func (r *MemoryUserRepository) Save(user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[user.ID] = user
	return nil
}
