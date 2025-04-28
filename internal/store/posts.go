package store

import (
	"context"
	"database/sql"
)

type Posts interface {
	Create(ctx context.Context) error
}

// Make sure that the PostStorage implements the Posts interface
var _ Posts = (*PostStorage)(nil)

type PostStorage struct {
	db *sql.DB
}

func (s *PostStorage) Create(ctx context.Context) error {
	return nil
}
