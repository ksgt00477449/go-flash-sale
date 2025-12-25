package repository

import (
	"errors"
	"go-flash-sale/internal/model"
	"sync"
)

type UserRepository struct {
	users map[string]*model.User
	mu    sync.RWMutex
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]*model.User),
	}
}

func (r *UserRepository) Create(user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.Email]; exists {
		return errors.New("email already registered!")
	}
	r.users[user.Email] = user
	return nil
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[email]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}
