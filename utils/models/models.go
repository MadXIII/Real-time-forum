package models

import "time"

type User struct {
	ID        int    `json:"id"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Confirm   string `json:"confirm"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Age       int    `json:"age"`
}

type Post struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
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

//nolint
type CommentLike struct {
	UserID    int `json:"user_id"`
	CommentID int `json:"comment_id"`
	VoteState int `json:"vote_state"`
}

type Sign struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
