package sqlite

import (
	"forum/utils/models"
)

// func (s *Store) GetUserById() (int, error) {

// }

//GetUserByLogin Searching User in database by Nickname
func (s *Store) GetUserByLogin(login string) (models.User, error) {
	var user models.User
	err := s.db.QueryRow(`
	SELECT * FROM user WHERE nickname = ? OR email = ?
	`, login, login).
		Scan(&user.ID, &user.Nickname, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Gender, &user.Age)
	if err != nil {
		return user, err
	}
	return user, nil
}

//GetUsernameByID - ...
func (s *Store) GetUsernameByID(id int) (string, error) {
	var username string
	err := s.db.QueryRow(`
	SELECT nickname FROM user WHERE id = ?
	`, id).
		Scan(&username)

	if err != nil {
		return "", err
	}
	return username, nil
}
