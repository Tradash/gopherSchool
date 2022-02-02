package sqlstore

import (
	"database/sql"
	"github.com/Tradash/gopherSchool/internal/app/store"
	_ "github.com/go-sql-driver/mysql"
)

// Store ...
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{store: s}
	return s.userRepository
}
