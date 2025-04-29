package store

import (
	"context"
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"log"
)

type Post struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	Title     string    `json:"title"`
	UserID    int64     `json:"user_id"`
	Tags      []string  `json:"tags"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Comments  []Comment `json:"comments"`
}

type Posts interface {
	Create(context.Context, *Post) error
	GetByID(context.Context, int64) (*Post, error)
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

	var id int64
	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Content,
		post.Title,
		post.UserID,
		pq.Array(post.Tags),
	).Scan(
		&id,
	)

	if err != nil {
		return err
	}

	post.ID = id

	log.Printf("Created post with id %d", post.ID)
	return nil
}

func (s *PostStorage) GetByID(ctx context.Context, id int64) (*Post, error) {
	query := `
		SELECT id, content, title, user_id, tags, created_at, updated_at
		FROM posts
		WHERE id = $1
	`

	var post Post

	err := s.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
		&post.ID,
		&post.Content,
		&post.Title,
		&post.UserID,
		pq.Array(&post.Tags),
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &post, nil
}
