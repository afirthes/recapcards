package store

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
	"log"
)

type Post struct {
	ID        int64    `json:"id"`
	Content   string   `json:"content"`
	Title     string   `json:"title"`
	UserID    int64    `json:"user_id"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type Posts interface {
	Create(context.Context, *Post) error
}

// Make sure that the PostStorage implements the Posts interface
var _ Posts = (*PostStorage)(nil)

type PostStorage struct {
	db *sql.DB
}

func (s *PostStorage) Create(ctx context.Context, post *Post) error {

	query := `
	INSERT INTO posts (content, title, user_id, tags)
	VALUES ($1, $2, $3, $4)
	RETURNING id`

	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Content,
		post.Title,
		post.UserID,
		pq.Array(post.Tags),
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return err
	}

	log.Printf("Created post with id %d", post.ID)
	return nil
}
