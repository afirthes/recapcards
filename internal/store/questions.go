package store

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
)

type Question struct {
	ID         int64    `json:"id"`
	Title      string   `json:"title"`
	Question   string   `json:"question"`
	Answer     string   `json:"answer"`
	CreatorID  int64    `json:"user_id"`
	Tags       []string `json:"tags"`
	CreatedAt  string   `json:"created_at"`
	UpdatedAt  string   `json:"updated_at"`
	IsPublic   bool     `json:"is_public"`
	CategoryID int64    `json:"category_id"`
}

type QuestionStore struct {
	db *sql.DB
}

func (s *QuestionStore) Create(ctx context.Context, question *Question) error {
	query := `
		INSERT INTO questions (
			user_id,
			title,
			question,
			answer,
			tags,
			is_public
		) VALUES (
			$1, $2, $3, $4, $5, $6
		) RETURNING id, created_at, updated_at;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	return s.db.QueryRowContext(
		ctx,
		query,
		question.CreatorID,
		question.Title,
		question.Question,
		question.Answer,
		pq.Array(&question.Tags),
		true,
	).Scan(
		&question.ID,
		&question.CreatedAt,
		&question.UpdatedAt,
	)

}
