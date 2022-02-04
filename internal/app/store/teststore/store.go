package teststore

import (
	"github.com/Tradash/gopherSchool/internal/app/model"
	"github.com/Tradash/gopherSchool/internal/app/store"
)

// Store ...
type Store struct {
	userRepository *UserRepository
}

// New ...
func New() *Store {
	return &Store{}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{store: s, users: make(map[int64]*model.User)}
	return s.userRepository
}
