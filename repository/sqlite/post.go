package sqlite

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/madxiii/real-time-forum/model"
)

type Post struct {
	db *sqlx.DB
}

func NewPost(db *sqlx.DB) *Post {
	return &Post{db: db}
}

// CreatePost - insert newpost in db
func (p *Post) CreatePost(newPost *model.Post) (int, error) {
	createRow, err := p.db.Prepare(`
		INSERT INTO post 
		(username, category_id, title, content, timestamp, like_count)
		VALUES (?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return 0, fmt.Errorf("CreatePost, Prepare: %w", err)
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
		return 0, fmt.Errorf("CreatePost, Exec: %w", err)
	}

	postid, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("CreatePost, LastInsertId: %w", err)
	}
	newPost.ID = int(postid)

	return newPost.ID, nil
}

// GetPostByID - Get post by postID from db
func (p *Post) GetPostByID(id string) (model.Post, error) {
	var post model.Post

	err := p.db.QueryRow(`
		SELECT * FROM post WHERE id = ?
	`, id).Scan(&post.ID, &post.CategoryID, &post.Username, &post.Title, &post.Content, &post.Timestamp, &post.LikeCount)
	if err != nil {
		return post, fmt.Errorf("GetPostByID, Scan: %w", err)
	}

	return post, nil
}

// GetAllPostsByCategoryID - Get all posts by CategoryID to show in main page
func (p *Post) GetAllPostsByCategoryID(categoryID int) (posts []model.Post, err error) {
	var rows *sqlx.Rows

	if categoryID < 2 {
		rows, err = p.db.Queryx(`
		SELECT * FROM post ORDER BY id DESC
	`)
	} else {
		rows, err = p.db.Queryx(`
		SELECT * FROM post WHERE category_id = ? ORDER BY id DESC
		`, categoryID)
	}

	defer rows.Close()

	if err != nil {
		return nil, fmt.Errorf("GetAllPosts, Query: %w", err)
	}

	for rows.Next() {
		var post model.Post
		if err := rows.Scan(&post.ID, &post.CategoryID, &post.Username, &post.Title, &post.Content, &post.Timestamp, &post.LikeCount); err != nil {
			return nil, fmt.Errorf("GetAllPosts, Scan: %w", err)
		}
		posts = append(posts, post)
	}

	return posts, nil
}
