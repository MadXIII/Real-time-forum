package sqlite

import (
	"forum/utils/models"
)

func (s *Store) InsertLike(like *models.PostLike) error {
	// like.VoteState = true // got liked

	createRow, err := s.db.Prepare(`
		INSERT INTO postlike
		(user_id, post_id, like)
		VALUES (?, ?, ?)
	`)

	defer createRow.Close()

	if err != nil {
		return err
	}

	_, err = createRow.Exec(
		like.UserID,
		like.PostID,
		like.VoteState,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetVoteState(like *models.PostLike) (bool, error) {
	var vote bool

	if err := s.db.QueryRow(`
		SELECT like FROM postlike
		WHERE post_id = ? AND user_id = ?
	`, like.PostID, like.UserID).Scan(&vote); err != nil {
		return false, err
	}

	return vote, nil
}

func (s *Store) UpdateVoteState(like *models.PostLike) error {
	_, err := s.db.Exec(`
	UPDATE postlike SET like = ?
	WHERE post_id = ? AND user_id = ?
	`, like.VoteState, like.PostID, like.UserID)
	if err != nil {
		return err
	}
	return nil
}
