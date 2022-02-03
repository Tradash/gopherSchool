package apiserver

import (
	"database/sql"
	"github.com/Tradash/gopherSchool/internal/app/store/sqlstore"
	"github.com/gorilla/sessions"
	"net/http"
)

// Start ...
func Start(config *Config) error {
	db, err := newDB(config.DatabaseUrl)
	if err != nil {
		return err
	}
	defer db.Close()
	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := newServer(store, sessionStore)
	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(dbUrl string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
