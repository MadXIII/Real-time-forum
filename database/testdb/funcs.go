package testdb

import (
	"forum/models"
)

type TestDB struct {
	// name string
	// pass string
}

func (t *TestDB) Init(dbname string) error {
	return nil
}
func (t *TestDB) InsertUser(models.User) error {
	return nil
}
func (t *TestDB) GetUserByLogin(string) (*models.User, error) {

	return &models.User{Nickname: "User", Password: "Pass"}, nil
}
func (t *TestDB) InsertPost(models.Post) error {
	return nil
}
func (t *TestDB) Close() {

}
