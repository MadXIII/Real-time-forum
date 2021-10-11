package sqlite

import (
	"forum/models"
	"log"
)

func (s *Store) InsertPost(newPost models.Post) error {
	createTable, err := s.db.Prepare(`
		INSERT INTO post 
		(tittle, content, timestamp)
		VALUES (?, ?, ?)
	`)
	if err != nil {
		log.Println(err)
		return err
	}

	defer s.db.Close()

	res, err := createTable.Exec(
		newPost.Title,
		newPost.Content,
		newPost.Timestamp,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	postid, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return err
	}
	newPost.ID = int(postid)

	return nil
}
