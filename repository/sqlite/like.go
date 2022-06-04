package sqlite

import (
	"fmt"
	"forum/internal/models"

	"github.com/jmoiron/sqlx"
)

type Vote struct {
	db *sqlx.DB
}

func NewVote(db *sqlx.DB) *Vote {
	return &Vote{db: db}
}

// InsertVote - Insert Vote in db by user in Post
func (v *Vote) InsertVote(like *models.PostLike) error {
	createRow, err := v.db.Prepare(`
		INSERT INTO postlike (user_id, post_id, like) VALUES (?, ?, ?)
	`)

	defer createRow.Close()

	if err != nil {
		return fmt.Errorf("InsertVote, Prepare: %w", err)
	}

	_, err = createRow.Exec(
		like.UserID,
		like.PostID,
		like.VoteState,
	)
	if err != nil {
		return fmt.Errorf("InsertVote, Exec: %w", err)
	}

	return nil
}

// GetVoteState - get state of vote of post from db
func (v *Vote) GetVoteState(like *models.PostLike) (bool, error) {
	var vote bool

	if err := v.db.QueryRow(`SELECT like FROM postlike WHERE post_id = ? AND user_id = ?
	`, like.PostID, like.UserID).Scan(&vote); err != nil {
		return false, fmt.Errorf("GetVoteState, QueryRow: %w", err)
	}

	return vote, nil
}

// UpdateVoteState - update state of vote of post in db
func (v *Vote) UpdateVoteState(like *models.PostLike) error {
	_, err := v.db.Exec(`UPDATE postlike SET like = ? WHERE post_id = ? AND user_id = ?
	`, like.VoteState, like.PostID, like.UserID)
	if err != nil {
		return fmt.Errorf("UpdateVoteState, Exec: %w", err)
	}
	return nil
}
