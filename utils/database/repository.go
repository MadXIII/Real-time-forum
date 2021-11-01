package database

import "forum/utils/models"

//Repository - interface to work with DB
type Repository interface {
	Init(dbname string) error
	InsertUser(models.User) error
	GetUserByLogin(string) (*models.User, error)
	InsertPost(models.Post) error
	GetPostByID(int) (*models.Post, error)
	Close()
}
