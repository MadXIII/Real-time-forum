package sqlite

import "fmt"

func (s *Store) InsertCategories(categories []string) error {
	for _, category := range categories {
		categoryRow, err := s.db.Prepare(`
		INSERT INTO category
		(name) VALUES (?)
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
