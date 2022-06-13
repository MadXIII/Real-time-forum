package sqlite

import (
	"fmt"

	"github.com/madxiii/real-time-forum/model"

	"github.com/jmoiron/sqlx"
)

type Vote struct {
	db *sqlx.DB
}

func NewVote(db *sqlx.DB) *Vote {
	return &Vote{db: db}
}

// CreateVote - Insert Vote in db by user in Post
func (v *Vote) CreateVote(like *model.PostLike) error {
	createRow, err := v.db.Prepare(`
		INSERT INTO postlike (user_id, post_id, like) VALUES (?, ?, ?)
	`)
	if err != nil {
		return fmt.Errorf("CreateVote, Prepare: %w", err)
	}
	defer createRow.Close()

	_, err = createRow.Exec(
		like.UserID,
		like.PostID,
		like.VoteState,
	)
	if err != nil {
		return fmt.Errorf("CreateVote, Exec: %w", err)
	}

	return nil
}

// GetVoteState - get state of vote of post from db
func (v *Vote) VoteState(like *model.PostLike) (bool, error) {
	var vote bool

	if err := v.db.QueryRow(`SELECT like FROM postlike WHERE post_id = ? AND user_id = ?
	`, like.PostID, like.UserID).Scan(&vote); err != nil {
		return false, fmt.Errorf("GetVoteState, QueryRow: %w", err)
	}

	return vote, nil
}

// UpdateVoteState - update state of vote of post in db
func (v *Vote) UpdateVoteState(like *model.PostLike) error {
	_, err := v.db.Exec(`UPDATE postlike SET like = ? WHERE post_id = ? AND user_id = ?
	`, like.VoteState, like.PostID, like.UserID)
	if err != nil {
		return fmt.Errorf("UpdateVoteState, Exec: %w", err)
	}
	return nil
}

// UpdateVotes - udate LikeCount in Posts data
func (p *Vote) UpdateVotes(like *model.PostLike) error {
	if like.VoteState {
		_, err := p.db.Exec(`
			UPDATE post SET like_count = like_count + 1
			WHERE id = ?
		`, like.PostID)
		if err != nil {
			return fmt.Errorf("UpdateVotes, incrementLike: %w", err)
		}
	}
	if !like.VoteState {
		_, err := p.db.Exec(`
		UPDATE post SET like_count = like_count - 1
		WHERE id = ?
		`, like.PostID)
		if err != nil {
			return fmt.Errorf("UpdateVotes, decrementLike: %w", err)
		}
	}
	return nil
}
