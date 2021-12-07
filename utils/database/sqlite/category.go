package sqlite

import (
	"fmt"
	newErr "forum/utils/internal/error"
	"forum/utils/models"
)

func (s *Store) InsertCategories(categories []string) error {
	for _, category := range categories {
		categoryRow, err := s.db.Prepare(`INSERT INTO category (name) VALUES (?)
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

func (s *Store) GetCategories() ([]models.Categories, error) {
	var categories []models.Categories

	rows, err := s.db.Query(`
		SELECT * FROM category
	`)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var category models.Categories

	for rows.Next() {
		if err = rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (s *Store) CheckCategoryID(categoryID int) error {
	if categoryID < 1 {
		return newErr.ErrWrongCategory
	}
	var id int
	if err := s.db.QueryRow(`SELECT id FROM category WHERE id = ?`, categoryID).Scan(&id); err != nil {
		return err
	}

	return nil
}
