package database

import (
	"context"
	"forum/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// Repository - interface to work with DB
type Repository interface {
	InitMongoStore(context.Context, string) (*mongo.Client, error)
	InitMainStore(dbname string) error
	// User
	Post
	Comment
	Vote
	Category
	Close()
}

type User interface {
	InsertUser(context.Context, *models.User) error
	GetUserByLogin(context.Context, string) (models.User, error)
	GetUsernameByID(context.Context, int) (string, error)
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
