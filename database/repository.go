package database

import "forum/models"

type Repository interface {
	Init(dbname string) error
	// 	User() UserRepository
}

type UserRepository interface {
	// TableUser()
	// GetUserByID()
	InsertUser(models.User, Repository) error
}
