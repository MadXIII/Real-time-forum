package database

import "forum/models"

type Repository interface {
	Init(dbname string) error
	InsertUser(models.User) error
}
