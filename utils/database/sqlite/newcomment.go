package sqlite

import "forum/utils/models"

func (s *Store) InsertComment(newComment *models.Comment) (int, error) {
	createRow, err := s.db.Prepare(`
		INSERT INTO comment
		(user_id, post_id, username, content, timestamp)
		VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		return 0, err
	}

	defer createRow.Close()

	res, err := createRow.Exec(
		newComment.UserID,
		newComment.PostID,
		newComment.Username,
		newComment.Content,
		newComment.Timestamp,
	)
	if err != nil {
		return 0, err
	}

	commnetID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	newComment.ID = int(commnetID)

	return newComment.ID, nil
}
