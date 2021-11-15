package sqlite

import (
	"fmt"
	"forum/utils/models"
)

//InsertPost ...
func (s *Store) InsertPost(newPost models.Post) (int, error) {
	createTable, err := s.db.Prepare(`
		INSERT INTO post 
		(user_id, username, title, content, timestamp)
		VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		fmt.Println(1)

		return 0, err
	}

	res, err := createTable.Exec(
		newPost.UserID,
		newPost.Username,
		newPost.Title,
		newPost.Content,
		newPost.Timestamp,
	)
	if err != nil {
		fmt.Println(2)
		return 0, err
	}

	postid, err := res.LastInsertId()
	if err != nil {
		fmt.Println(3)
		return 0, err
	}
	newPost.ID = int(postid)

	return newPost.ID, nil
}
