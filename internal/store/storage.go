package store

import (
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("resource not found")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Posts Posts
	Users Users
}

func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStorage{db},
		Users: &UserStorage{db},
	}
}
