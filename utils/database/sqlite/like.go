package sqlite

import "forum/utils/models"

func (s *Store) InsertLike(like *models.PostLike) error {
	createRow, err := s.db.Prepare(`
		INSERT INTO postlike
		(user_id, post_id, like)
		VALUES (?, ?, ?)
	`)

	defer createRow.Close()

	if err != nil {
		return err
	}

	res, err := createRow.Exec(
		like.UserID,
		like.PostID,
		like.VoteState,
	)
	if err != nil {
		return err
	}

	res = res

	return nil
}
