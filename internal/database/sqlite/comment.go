package sqlite

import (
	"fmt"
	"forum/internal/models"
)

// InsertComment - Insert new comment in db
func (s *Store) InsertComment(newComment *models.Comment) error {
	createRow, err := s.db.Prepare(`
		INSERT INTO comment
		(post_id, username, content, timestamp)
		VALUES (?, ?, ?, ?)
		`)
	if err != nil {
		return fmt.Errorf("InsertComment, Prepare: %w", err)
	}
	defer createRow.Close()

	res, err := createRow.Exec(
		newComment.PostID,
		newComment.Username,
		newComment.Content,
		newComment.Timestamp,
	)
	if err != nil {
		return fmt.Errorf("InsertComment, Exec: %w", err)
	}

	commnetID, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("InsertComment, LastInsertId: %w", err)
	}
	newComment.ID = int(commnetID)

	return nil
}

// GetCommentsByPostID - Get slice of all comments by postID
func (s *Store) GetCommentsByPostID(pid string) ([]models.Comment, error) {
	var comments []models.Comment

	rows, err := s.db.Query(`
		SELECT * FROM comment WHERE post_id = ?
	`, pid)
	if err != nil {
		return nil, fmt.Errorf("GetCommentsByPostID, Query: %w", err)
	}
	defer rows.Close()

	var comment models.Comment
	for rows.Next() {
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.Username, &comment.Content, &comment.Timestamp); err != nil {
			return nil, fmt.Errorf("GetCommentsByPostID, Scan: %w", err)
		}
		comments = append(comments, comment)
	}

	return comments, nil
}
