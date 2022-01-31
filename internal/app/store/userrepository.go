package store

import (
	"github.com/Tradash/gopherSchool/internal/app/model"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create ...
func (r *UserRepository) Create(u *model.User) (*model.User, error) {

	result, err := r.store.db.Exec("insert into users (email, encrypted_password) values (? ,?)", u.Email, u.EncryptedPassword)
	if err != nil {
		return nil, err
	}
	u.ID, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return u, nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow("select id, email, encrypted_password from users where email = ?", email).Scan(&u.ID, &u.Email, &u.EncryptedPassword); err != nil {
		return nil, err
	}
	return u, nil
}
