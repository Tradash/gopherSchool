package sqlstore

import (
	"database/sql"
	"github.com/Tradash/gopherSchool/internal/app/model"
	"github.com/Tradash/gopherSchool/internal/app/store"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create ...
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	result, err := r.store.db.Exec("insert into users (email, encrypted_password) values (? ,?)", u.Email, u.EncryptedPassword)
	if err != nil {

		return err
	}
	u.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow("select id, email, encrypted_password from users where email = ?", email).Scan(&u.ID, &u.Email, &u.EncryptedPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}

// Find ...
func (r *UserRepository) Find(id int64) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow("select id, email, encrypted_password from users where id = ?", id).Scan(&u.ID, &u.Email, &u.EncryptedPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}
