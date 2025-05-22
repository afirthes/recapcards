package store

import (
	"context"
	"database/sql"
	"testing"
)

func TestCategoriesStore_Create(t *testing.T) {
	withDB(t, func(db *sql.DB, t *testing.T) error {

		ctx := context.Background()
		user, err := createUser(ctx, db)

		cs := &CategoriesStore{db}

		category := &Category{
			ID:        1,
			Title:     "Test Title",
			CreatorID: user.ID,
			IsPublic:  true,
		}

		err = cs.Create(ctx, category)
		if err != nil {
			t.Fatalf("failed to insert category: %v", err)
		}

		if category.ID == 0 || category.CreatedAt == "" || category.UpdatedAt == "" {
			t.Fatalf("expected question to be populated, got %+v", category)
		}
		return nil
	})
}
