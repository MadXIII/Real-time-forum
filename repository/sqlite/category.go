package sqlite

import (
	"fmt"
	newErr "forum/internal/error"
	"forum/internal/models"

	"github.com/jmoiron/sqlx"
)

type Category struct {
	db *sqlx.DB
}

func NewCategory(db *sqlx.DB) *Category {
	return &Category{db: db}
}

// InsertCategories - insert Categories in db, while we Init db
func (c *Category) InsertCategories(categories []string) error {
	for _, category := range categories {
		categoryRow, err := c.db.Prepare(`INSERT INTO category (name) VALUES (?)
		`)
		defer categoryRow.Close()

		if err != nil {
			return fmt.Errorf("InsertCategories, Prepare: %w", err)
		}

		_, err = categoryRow.Exec(category)

		if err != nil && err.Error() != "UNIQUE constraint failed: category.name" {
			return fmt.Errorf("InsertCategories, Exec: %w", err)
		}
	}

	return nil
}

// GetCategories - Get all Categories from db
func (c *Category) GetCategories() ([]models.Categories, error) {
	var categories []models.Categories

	rows, err := c.db.Query(`
		SELECT * FROM category
	`)

	defer rows.Close()

	if err != nil {
		return nil, fmt.Errorf("GetCategories, Query: %w", err)
	}

	var category models.Categories

	for rows.Next() {
		if err = rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, fmt.Errorf("GetCategories, Scan: %w", err)
		}
		categories = append(categories, category)
	}

	return categories, nil
}

// CheckCategoryID - check request Categories for correct ID
func (c *Category) CheckCategoryID(categoryID int) error {
	if categoryID < 1 || categoryID > 4 {
		return newErr.ErrWrongCategory
	}
	var id int
	if err := c.db.QueryRow(`SELECT id FROM category WHERE id = ?`, categoryID).Scan(&id); err != nil {
		return fmt.Errorf("CheckCategoryID, QueryRow: %w", err)
	}

	return nil
}
