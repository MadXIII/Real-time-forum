package sqlite

import "forum/utils/models"

func (s *Store) GetPostByID(id int) (*models.Post, error) {
	var post models.Post
	err := s.db.QueryRow(`
		SELECT * FROM post WHERE id = ?
	`, id).Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.Timestamp)
	if err != nil {
		return nil, err
	}

	return &models.Post{}, err
}
