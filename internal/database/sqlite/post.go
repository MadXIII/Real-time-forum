package sqlite

import (
	"database/sql"
	"fmt"
	"forum/internal/models"
)

//InsertPost - insert newpost in db
func (s *Store) InsertPost(newPost *models.Post) (int, error) {
	createRow, err := s.db.Prepare(`
		INSERT INTO post 
		(username, category_id, title, content, timestamp, like_count)
		VALUES (?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return 0, fmt.Errorf("InsertPost, Prepare: %w", err)
	}

	res, err := createRow.Exec(
		newPost.Username,
		newPost.CategoryID,
		newPost.Title,
		newPost.Content,
		newPost.Timestamp,
		newPost.LikeCount,
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
	`, id).Scan(&post.ID, &post.CategoryID, &post.Username, &post.Title, &post.Content, &post.Timestamp, &post.LikeCount)
	if err != nil {
		return post, fmt.Errorf("GetPostByID, Scan: %w", err)
	}

	return post, nil
}

//GetAllPostsByCategoryID - Get all posts by CategoryID to show in main page
func (s *Store) GetAllPostsByCategoryID(categoryID int) (posts []models.Post, err error) {
	var rows *sql.Rows

	if categoryID < 2 {
		rows, err = s.db.Query(`
		SELECT * FROM post ORDER BY id DESC
	`)
	} else {
		rows, err = s.db.Query(`
		SELECT * FROM post WHERE category_id = ? ORDER BY id DESC
		`, categoryID)
	}

	defer rows.Close()

	if err != nil {
		return nil, fmt.Errorf("GetAllPosts, Query: %w", err)
	}

	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.CategoryID, &post.Username, &post.Title, &post.Content, &post.Timestamp, &post.LikeCount); err != nil {
			return nil, fmt.Errorf("GetAllPosts, Scan: %w", err)
		}
		posts = append(posts, post)
	}

	return posts, nil
}

//UpdateVotes - udate LikeCount in Posts data
func (s *Store) UpdateVotes(like *models.PostLike) error {
	if like.VoteState {
		_, err := s.db.Exec(`
			UPDATE post SET like_count = like_count + 1
			WHERE id = ?
		`, like.PostID)
		if err != nil {
			return fmt.Errorf("UpdateVotes, incrementLike: %w", err)
		}
	}
	if !like.VoteState {
		_, err := s.db.Exec(`
		UPDATE post SET like_count = like_count - 1
		WHERE id = ?
		`, like.PostID)
		if err != nil {
			return fmt.Errorf("UpdateVotes, decrementLike: %w", err)
		}
	}
	return nil
}
