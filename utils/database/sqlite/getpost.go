package sqlite

import (
	"fmt"
	"forum/utils/models"
)

func (s *Store) GetPostByID(id int) (*models.Post, error) {
	var post models.Post
	err := s.db.QueryRow(`
		SELECT * FROM post WHERE id = ?
	`, id).Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.Timestamp)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &post, err
}
