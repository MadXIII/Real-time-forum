package sqlite

import "forum/utils/models"

func (s *Store) InsertComment(newComment *models.Comment) error {
	createRow, err := s.db.Prepare(`
		INSERT INTO comment
		(post_id, username, content, timestamp)
		VALUES (?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}

	defer createRow.Close()

	res, err := createRow.Exec(
		newComment.PostID,
		newComment.Username,
		newComment.Content,
		newComment.Timestamp,
	)
	if err != nil {
		return err
	}

	commnetID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	newComment.ID = int(commnetID)

	return nil
}
