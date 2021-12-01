package sqlite

import (
	"fmt"
	"forum/utils/models"
)

//InsertPost - insert newpost in db
func (s *Store) InsertPost(newPost *models.Post) (int, error) {
	createRow, err := s.db.Prepare(`
		INSERT INTO post 
		(username, title, content, timestamp, diffLikes)
		VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		return 0, fmt.Errorf("InsertPost, Prepare: %w", err)
	}

	res, err := createRow.Exec(
		newPost.Username,
		newPost.Title,
		newPost.Content,
		newPost.Timestamp,
		newPost.LikeDis,
	)
	if err != nil {
		return 0, fmt.Errorf("InsertPost, Exec: %w", err)
	}

	postid, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("InsertPost, LastInsertId: %w", err)
	}
	newPost.ID = int(postid)

	return newPost.ID, nil
}

//GetPostByID - Get post by postID from db
func (s *Store) GetPostByID(id string) (models.Post, error) {
	var post models.Post

	err := s.db.QueryRow(`
		SELECT * FROM post WHERE id = ?
	`, id).Scan(&post.ID, &post.Username, &post.Title, &post.Content, &post.Timestamp, &post.LikeDis)
	if err != nil {
		return post, fmt.Errorf("GetPostByID, Scan: %w", err)
	}

	return post, nil
}

//GetAllPosts - Get all posts to show in main page
func (s *Store) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post

	rows, err := s.db.Query(`
		SELECT * FROM post ORDER BY id DESC
	`)
	defer rows.Close()

	if err != nil {
		return nil, fmt.Errorf("GetAllPosts, Query: %w", err)
	}

	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Username, &post.Title, &post.Content, &post.Timestamp, &post.LikeDis); err != nil {
			return nil, fmt.Errorf("GetAllPosts, Scan: %w", err)
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (s *Store) ChangeLikeDislikeDiff(pid int, up bool) {
	if up {
		s.db.Exec(`
			UPDATE post SET diffLikes = diffLikes + 1
			WHERE id = ?
		`, pid)
	} else {
		s.db.Exec(`
			UPDATE post SET diffLikes = diffLikes - 1
			WHERE id = ?
		`, pid)
	}
}
