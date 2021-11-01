package sqlite

import (
	"forum/utils/models"
)

// func (s *Store) GetUserById() (int, error) {

// }

//GetUserByLogin Searching User in database by Nickname
func (s *Store) GetUserByLogin(login string) (*models.User, error) {
	var user models.User
	err := s.db.QueryRow(`
	SELECT * FROM user WHERE nickname = ? OR email = ?
	`, login, login).
		Scan(&user.ID, &user.Nickname, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Gender, &user.Age)
	if err != nil {
		return nil, err
	}

	// defer rows.Close()

	// for rows.Next() {
	// 	err := rows.Scan(&user.ID, &user.Nickname, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Gender, &user.Age)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	// if err := rows.Err(); err != nil {
	// 	return nil, err
	// }

	return &user, err
}
