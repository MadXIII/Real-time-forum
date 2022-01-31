package testdb

import (
	"forum/internal/models"

	"github.com/stretchr/testify/mock"
)

type TestDB struct {
	mock.Mock
}

func (db *TestDB) Init(dbname string) error {
	return nil
}

func (db *TestDB) InsertUser(user *models.User) error {
	args := db.Called(user)
	return args.Error(0)
}

func (db *TestDB) GetUserByLogin(login string) (models.User, error) {
	args := db.Called(login)
	return args.Get(0).(models.User), args.Error(1)
}

func (db *TestDB) GetUsernameByID(id int) (string, error) {
	args := db.Called(id)
	return args.Get(0).(string), args.Error(1)
}

//output problem - int
func (db *TestDB) InsertPost(post *models.Post) (int, error) {
	args := db.Called(post)
	return args.Get(0).(int), args.Error(1)
}

func (db *TestDB) GetPostByID(id string) (models.Post, error) {
	args := db.Called(id)
	return args.Get(0).(models.Post), args.Error(1)
}

func (db *TestDB) GetAllPostsByCategoryID(id int) ([]models.Post, error) {
	args := db.Called(id)
	return args.Get(0).([]models.Post), args.Error(1)
}

func (db *TestDB) InsertComment(Comment *models.Comment) error {
	args := db.Called(Comment)
	return args.Error(0)
}

func (db *TestDB) GetCommentsByPostID(id string) ([]models.Comment, error) {
	args := db.Called(id)
	return args.Get(0).([]models.Comment), args.Error(1)
}

func (db *TestDB) InsertVote(vote *models.PostLike) error {
	args := db.Called(vote)
	return args.Error(0)
}

func (db *TestDB) GetVoteState(vote *models.PostLike) (bool, error) {
	args := db.Called(vote)
	return args.Get(0).(bool), args.Error(1)
}

func (db *TestDB) UpdateVoteState(vote *models.PostLike) error {
	args := db.Called(vote)
	return args.Error(0)
}

func (db *TestDB) UpdateVotes(vote *models.PostLike) error {
	args := db.Called(vote)
	return args.Error(0)
}

func (db *TestDB) InsertCategories(Categories []string) error {
	args := db.Called(Categories)
	return args.Error(0)
}

func (db *TestDB) GetCategories() ([]models.Categories, error) {
	args := db.Called()
	return args.Get(0).([]models.Categories), args.Error(1)
}

func (db *TestDB) CheckCategoryID(id int) error {
	args := db.Called(id)
	return args.Error(0)
}

func (db *TestDB) Close() {
}
