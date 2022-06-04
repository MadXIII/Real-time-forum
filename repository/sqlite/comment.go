package sqlite

import (
	"fmt"
	"forum/internal/models"

	"github.com/jmoiron/sqlx"
)

type Comment struct {
	db *sqlx.DB
}

func NewComment(db *sqlx.DB) *Comment {
	return &Comment{db: db}
}

// InsertComment - Insert new comment in db
func (c *Comment) InsertComment(newComment *models.Comment) error {
	createRow, err := c.db.Prepare(`
		INSERT INTO comment
		(post_id, username, content, timestamp)
		VALUES (?, ?, ?, ?)
		`)

	defer createRow.Close()

	if err != nil {
		return fmt.Errorf("InsertComment, Prepare: %w", err)
	}

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
func (c *Comment) GetCommentsByPostID(pid string) ([]models.Comment, error) {
	var comments []models.Comment

	rows, err := c.db.Query(`
		SELECT * FROM comment WHERE post_id = ?
	`, pid)

	defer rows.Close()

	if err != nil {
		return nil, fmt.Errorf("GetCommentsByPostID, Query: %w", err)
	}

	var comment models.Comment
	for rows.Next() {
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.Username, &comment.Content, &comment.Timestamp); err != nil {
			return nil, fmt.Errorf("GetCommentsByPostID, Scan: %w", err)
		}
		comments = append(comments, comment)
	}

	return comments, nil
}
