package store

import (
	"context"
	"database/sql"
)

type Category struct {
	ID        int64  `json:"id"`
	ParentID  *int64 `json:"parent"`
	Title     string `json:"title"`
	CreatorID int64  `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	IsPublic  bool   `json:"is_public"`
}

type CategoriesStore struct {
	db *sql.DB
}

func (s *CategoriesStore) Create(ctx context.Context, category *Category) error {
	query := `
		INSERT INTO categories (
			parent,
			user_id,
			title,
			is_public
		) VALUES (
			$1, $2, $3, $4
		) RETURNING id, created_at, updated_at;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		category.ParentID,
		category.CreatorID,
		category.Title,
		true,
	).Scan(
		&category.ID,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil

}
