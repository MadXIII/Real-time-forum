package sqlite

import (
	"forum/utils/models"
)

//InsertPost ...
func (s *Store) InsertPost(newPost *models.Post) (int, error) {
	createRow, err := s.db.Prepare(`
		INSERT INTO post 
		(user_id, username, title, content, timestamp)
		VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		return 0, err
	}

	res, err := createRow.Exec(
		newPost.UserID,
		newPost.Username,
		newPost.Title,
		newPost.Content,
		newPost.Timestamp,
	)
	if err != nil {
		return 0, err
	}

	postid, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	newPost.ID = int(postid)

	return newPost.ID, nil
}
