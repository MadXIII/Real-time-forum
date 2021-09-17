package database

import "forum/models"

type Repository interface {
	Init(dbname string) error
	InsertUser(models.User) error
	// GetUserById
	GetUserByNickname(string) (models.User, error)
	GetUserByEmail(string) (models.User, error)
}
