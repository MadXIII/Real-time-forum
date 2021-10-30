package testdb

import (
	newErr "forum/internal/error"
	"forum/models"
)

type TestDB struct {
	user []models.User
	// name string
	// pass string
}

func (t *TestDB) Init(dbname string) error {
	return nil
}
func (t *TestDB) InsertUser(user models.User) error {
	t.user = append(t.user, user)
	return nil
}
func (t *TestDB) GetUserByLogin(login string) (*models.User, error) {
	for _, r := range t.user {
		if login == r.Nickname {
			return &r, nil
		}
	}
	return &models.User{}, newErr.ErrWrongLogin
}
func (t *TestDB) InsertPost(models.Post) error {
	return nil
}
func (t *TestDB) Close() {

}
