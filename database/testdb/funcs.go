package testdb

import (
	"forum/models"

	"github.com/stretchr/testify/mock"
)

type TestDB struct {
	mock.Mock
	// name string
	// pass string
}

func (db *TestDB) Init(dbname string) error {
	return nil
}
func (db *TestDB) InsertUser(user models.User) error {
	args := db.Called(user)
	return args.Error(0)
}
func (db *TestDB) GetUserByLogin(login string) (*models.User, error) {
	args := db.Called(login)
	return args.Get(0).(*models.User), args.Error(1)
}
func (db *TestDB) InsertPost(post models.Post) error {
	args := db.Called(post)
	return args.Error(0)
}
func (db *TestDB) Close() {

}
