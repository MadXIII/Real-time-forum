package repository

import (
	"github.com/madxiii/real-time-forum/model"
	"github.com/madxiii/real-time-forum/repository/sqlite"

	"github.com/jmoiron/sqlx"
)

// Repository - interface to work with DB
type Repository struct {
	User
	Post
	Comment
	Vote
	// Category
}

type User interface {
	CreateUser(user *model.User) error
	GetUserByLogin(login string) (model.User, error)
	GetUsernameByID(id int) (string, error)
}

type Post interface {
	CreatePost(post *model.Post) (int, error)
	GetPostByID(id string) (model.Post, error)
	// GetAllPostsByCategory(category int) ([]model.Post, error)
}

type Comment interface {
	InsertComment(comment *model.Comment) error
	GetCommentsByPostID(id string) ([]model.Comment, error)
}

type Vote interface {
	CreateVote(vote *model.PostLike) error
	VoteState(vote *model.PostLike) (bool, error)
	UpdateVoteState(vote *model.PostLike) error
	UpdateVotes(vote *model.PostLike) error
}

type Category interface {
	InsertCategories(categories []string) error
	GetCategories() ([]model.Categories, error)
	CheckCategoryID(id int) error
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		User:    sqlite.NewUser(db),
		Post:    sqlite.NewPost(db),
		Comment: sqlite.NewComment(db),
		Vote:    sqlite.NewVote(db),
		// Category: sqlite.NewCategory(db),
	}
}
