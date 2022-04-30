package database

import "forum/internal/models"

// Repository - interface to work with DB
type Repository interface {
	Init(dbname string) error
	User
	Post
	Comment
	Vote
	Category
	Close()
}

type User interface {
	InsertUser(*models.User) error
	GetUserByLogin(string) (models.User, error)
	GetUsernameByID(int) (string, error)
	GetAllUsernames() ([]string, error)
}

type Post interface {
	InsertPost(*models.Post) (int, error)
	GetPostByID(string) (models.Post, error)
	GetAllPostsByCategoryID(int) ([]models.Post, error)
}

type Comment interface {
	InsertComment(*models.Comment) error
	GetCommentsByPostID(string) ([]models.Comment, error)
}

type Vote interface {
	InsertVote(*models.PostLike) error
	GetVoteState(*models.PostLike) (bool, error)
	UpdateVoteState(*models.PostLike) error
	UpdateVotes(*models.PostLike) error
}

type Category interface {
	InsertCategories([]string) error
	GetCategories() ([]models.Categories, error)
	CheckCategoryID(int) error
}
