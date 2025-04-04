package sqlite

import (
	"database/sql"

	"github.com/kartikey1188/go-students-api/internal/config"
	_ "github.com/mattn/go-sqlite3" // this import for sqlite3 is required for the driver to register itself with the database/sql package, and the underscore indicates that we don't need to use it directly in our code; but we still need it to be included in the build
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath) // same as (*cfg).StoragePath, Go automatically dereferences struct pointers when accessing fields
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE,
	age INTEGER NOT NULL
	)`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{
		Db: db,
	}, nil
}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {
	stmt, err := s.Db.Prepare("INSERT INTO students (name, email, age) VALUES (?, ?, ?)")

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(name, email, age)

	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}
