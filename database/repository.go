package database

import "forum/models"

type Repository interface {
	// Init(dbname string) error
}

type UserRepository interface {
	// TableUser()
	// GetUserByID()
	InsertUser(models.User) error
}
