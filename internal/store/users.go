package store

import (
	"context"
	"database/sql"
)

type Users interface {
	Create(ctx context.Context) error
}

// Make sure that the UserStorage implements the Users interface
var _ Users = (*UserStorage)(nil)

type UserStorage struct {
	db *sql.DB
}

func (s *UserStorage) Create(ctx context.Context) error {
	return nil
}
