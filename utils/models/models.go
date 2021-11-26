package models

type User struct {
	ID        int    `json:"id"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Confirm   string `json:"confirm"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Age       string `json:"age"`
}

type Post struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
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
	ID        int    `json:"id"`
	PostID    int    `json:"cpost_id"`
	Username  string `json:"username"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}

type CommentLike struct {
	UserID    int `json:"user_id"`
	CommentID int `json:"comment_id"`
	VoteState int `json:"vote_state"`
}

type Sign struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
