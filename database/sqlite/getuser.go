package sqlite

import (
	"forum/models"
)

// func (s *Store) GetUserById() (int, error) {

// }

//GetUserByNickname Searching User in database by Nickname
func (s *Store) GetUserByNickname(nickname string) (models.User, error) {
	var user models.User

	rows, err := s.db.Query(`
		SELECT * FROM user WHERE nickname = ?
	`, nickname)

	defer rows.Close()

	if err != nil {
		return user, err
	}

	for rows.Next() {
		rows.Scan(&user.ID, &user.Nickname, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Gender, &user.Age)
	}

	return user, err
}

//GetUserByEmail Searching User in database by Email
func (s *Store) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	rows, err := s.db.Query(`
		SELECT * FROM user WHERE email = ?
	`, email)

	defer rows.Close()

	if err != nil {
		return user, err
	}

	for rows.Next() {
		rows.Scan(&user.ID, &user.Nickname, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Gender, &user.Age)
	}

	return user, err
}
