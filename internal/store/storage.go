package store

import (
	"database/sql"
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
