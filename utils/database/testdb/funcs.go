package testdb

import (
	"forum/utils/models"

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
func (db *TestDB) InsertUser(user *models.User) error {
	args := db.Called(user)
	return args.Error(0)
}
func (db *TestDB) GetUserByLogin(login string) (*models.User, error) {
	args := db.Called(login)
	return args.Get(0).(*models.User), args.Error(1)
}

//output problem - int
func (db *TestDB) InsertPost(post *models.Post) (int, error) {
	args := db.Called(post)
	return args.Get(0).(int), args.Error(1)
}

func (db *TestDB) GetPostByID(id string) (*models.Post, error) {
	args := db.Called(id)
	return args.Get(0).(*models.Post), args.Error(1)
}

func (db *TestDB) GetAllPosts() (*[]models.Post, error) {
	args := db.Called()
	return args.Get(0).(*[]models.Post), args.Error(1)
}
func (db *TestDB) GetUsernameByID(id int) (string, error) {
	args := db.Called(id)
	return args.Get(0).(string), args.Error(1)
}
func (db *TestDB) InsertComment(Comment *models.Comment) (int, error) {
	args := db.Called(Comment)
	return args.Get(0).(int), args.Error(1)
}
func (db *TestDB) Close() {

}
