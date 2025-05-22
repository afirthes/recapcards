package store

import (
	"context"
	"database/sql"
	"testing"
)

func createCategoryWithUser(ctx context.Context, db *sql.DB, user *User) (*Category, error) {
	cs := &CategoriesStore{db: db}
	category := &Category{}
	category.CreatorID = user.ID

	err := cs.Create(ctx, category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func createUser(ctx context.Context, db *sql.DB) (*User, error) {
	us := &UserStore{db: db}
	user := &User{}

	err := withTx(db, ctx, func(tx *sql.Tx) error {
		err := us.Create(ctx, tx, user)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func TestQuestionStore_Create(t *testing.T) {
	withDB(t, func(db *sql.DB, t *testing.T) error {

		ctx := context.Background()

		user, err := createUser(ctx, db)
		if err != nil {
			t.Fatalf("failed to create user: %v", err)
		}

		category, err := createCategoryWithUser(ctx, db, user)
		if err != nil {
			t.Fatalf("failed to create category: %v", err)
		}

		qs := &QuestionStore{db: db}
		question := &Question{
			ID:         1,
			CategoryID: category.ID,
			Title:      "Test Title",
			Question:   "What is testcontainers?",
			Answer:     "A Go library for integration testing with real containers.",
			CreatorID:  user.ID,
			Tags:       []string{"go", "testing"},
			IsPublic:   true,
		}

		err = qs.Create(ctx, question)
		if err != nil {
			t.Fatalf("failed to insert question: %v", err)
		}

		if question.ID == 0 || question.CreatedAt == "" || question.UpdatedAt == "" {
			t.Fatalf("expected question to be populated, got %+v", question)
		}
		return nil
	})
}
