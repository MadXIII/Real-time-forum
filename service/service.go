package service

import (
	"forum/model"
	"forum/repository"
)

type Service struct {
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
	GetVoteState(vote *model.PostLike) (bool, error)
	UpdateVoteState(vote *model.PostLike) error
	UpdateVotes(vote *model.PostLike) error
}

type Category interface {
	InsertCategories(categories []string) error
	GetCategories() ([]model.Categories, error)
	CheckCategoryID(id int) error
}

func New(repo repository.Repository) *Service {
	return &Service{
		User:    NewUser(repo.User),
		Post:    NewPost(repo.Post),
		Comment: NewComment(repo.Comment),
		Vote:    NewVote(repo.Vote),
		// Category: NewCategory(repo.Category),
	}
}
