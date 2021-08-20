package models

import "time"

type User struct {
	ID        int    `json:"id"`
	Age       int    `json:"age"`
	Nickname  string `json:"nickname"`
	Gender    string `json:"gender"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Post struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Thread    string    `json:"thread"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

type PostLike struct {
	UserID    int `json:"user_id"`
	PostID    int `json:"post_id"`
	VoteState int `json:"vote_state"`
}

type Comment struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	PostID    int       `json:"post_id"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

type CommentLike struct {
	UserID    int `json:"user_id"`
	CommentID int `json:"comment_id"`
	VoteState int `json:"vote_state"`
}
