package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// //sqlite3 User struct
// type User struct {
// 	ID        int    `json:"id"`
// 	Nickname  string `json:"nickname"`
// 	Email     string `json:"email"`
// 	Password  string `json:"password"`
// 	Confirm   string `json:"confirm"`
// 	FirstName string `json:"first_name,omitempty"`
// 	LastName  string `json:"last_name,omitempty"`
// 	Gender    string `json:"gender,omitempty"`
// 	Age       string `json:"age,omitempty"`
// }

type User struct {
	ID        primitive.ObjectID `json:"id, omitempty" bson:"_id, omitempty`
	Nickname  string             `json:"nickname, omitempty" bson:"nickname, omitempty`
	Email     string             `json:"email, omitempty" bson:"email, omitempty`
	Password  string             `json:"password, omitempty" bson:"password, omitempty`
	Confirm   string             `json:"confirm, omitempty" bson:"confirm, omitempty`
	FirstName string             `json:"first_name, omitempty" bson:"first_name, omitempty`
	LastName  string             `json:"last_name, omitempty" bson:"last_name, omitempty`
	Gender    string             `json:"gender, omitempty" bson:"gender, omitempty`
	Age       string             `json:"age, omitempty" bson:"age, omitempty`
}

type Categories struct {
	ID   int    `json:"id"`
	Name string `json:"category_name,omitempty"`
}

type Post struct {
	ID         int    `json:"id"`
	CategoryID int    `json:"category_id"`
	Username   string `json:"username"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Timestamp  string `json:"timestamp"`
	LikeCount  int    `json:"likes"`
}

type PostLike struct {
	UserID    int  `json:"user_id,omitempty"`
	PostID    int  `json:"post_id,omitempty"`
	VoteState bool `json:"vote_state,omitempty"`
}

type Comment struct {
	ID        int    `json:"id,omitempty"`
	PostID    int    `json:"cpost_id,omitempty"`
	Username  string `json:"username,omitempty"`
	Content   string `json:"content,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
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
