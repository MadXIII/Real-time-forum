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

func (s *Store) GetCommentsByPostID(pid string) ([]models.Comment, error) {
	var comments []models.Comment

	rows, err := s.db.Query(`
		SELECT * FROM comment WHERE post_id = ?
	`, pid)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.Username, &comment.Content, &comment.Timestamp); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
