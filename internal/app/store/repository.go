package store

import "github.com/Tradash/gopherSchool/internal/app/model"

// UserRepository ...
type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(string) (*model.User, error)
}
