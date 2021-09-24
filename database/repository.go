package database

import "forum/models"

//Repository ...
type Repository interface {
	Init(dbname string) error
	InsertUser(models.User) error
	GetUserByLogin(string) (*models.User, error)
}
