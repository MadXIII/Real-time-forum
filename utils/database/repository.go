package database

import "forum/utils/models"

//Repository - interface to work with DB
type Repository interface {
	Init(dbname string) error
	InsertUser(*models.User) error
	GetUserByLogin(string) (*models.User, error)
	InsertPost(models.Post) (int, error)
	GetPostByID(string) (*models.Post, error)
	GetAllPosts() (*[]models.Post, error)
	GetUsernameByID(int) (string, error)
	// GetMyPosts() (*[]models.Post, error)
	Close()
}
