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

func (s *Store) GetAllPosts() (*[]models.Post, error) {
	var posts []models.Post

	rows, err := s.db.Query(`
		SELECT * FROM post ORDER BY id DESC
	`)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post models.Post
		rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.Timestamp)
		posts = append(posts, post)
	}

	return &posts, nil
}
