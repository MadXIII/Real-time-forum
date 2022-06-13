package sqlite

import (
	"fmt"

	newErr "github.com/madxiii/real-time-forum/error"
	"github.com/madxiii/real-time-forum/model"

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
	for _, category := range categories { // FIXME without Prepare in For
		categoryRow, err := c.db.Prepare(`INSERT INTO category (name) VALUES (?) 
		`)
		if err != nil {
			return fmt.Errorf("InsertCategories, Prepare: %w", err)
		}

		defer categoryRow.Close()

		_, err = categoryRow.Exec(category)
		if err != nil && err.Error() != "UNIQUE constraint failed: category.name" { // FIXME unreadable
			return fmt.Errorf("InsertCategories, Exec: %w", err)
		}
	}
	return nil
}

// GetCategories - Get all Categories from db
func (c *Category) GetCategories() ([]model.Categories, error) {
	var categories []model.Categories

	rows, err := c.db.Query(`
		SELECT * FROM category
	`)
	if err != nil {
		return nil, fmt.Errorf("GetCategories, Query: %w", err)
	}
	defer rows.Close()

	var category model.Categories

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
