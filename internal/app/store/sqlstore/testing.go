package sqlstore

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"
)

// TestDb ...
func TestDb(t *testing.T, databaseUrl string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("mysql", databaseUrl)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec(fmt.Sprintf("truncate %s", strings.Join(tables, ", ")))
		}
		db.Close()
	}
}
