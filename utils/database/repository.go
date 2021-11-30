package database

import "forum/utils/models"

//Repository - interface to work with DB
type Repository interface {
	Init(dbname string) error
	InsertUser(*models.User) error
	GetUserByLogin(string) (models.User, error)
	InsertPost(*models.Post) (int, error)
	GetPostByID(string) (models.Post, error)
	GetAllPosts() ([]models.Post, error)
	GetUsernameByID(int) (string, error)
	InsertComment(*models.Comment) error
	GetCommentsByPostID(string) ([]models.Comment, error)
	InsertLike(*models.PostLike) error
	GetVoteState(int, int) (bool, error)
	UpdateVoteState(*models.PostLike)
	ChangeLikeDislikeDiff(int, bool)
	Close()
}
