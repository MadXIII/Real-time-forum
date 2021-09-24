package sqlite

import (
	"forum/models"
	"log"
)

// func (s *Store) GetUserById() (int, error) {

// }

//GetUserByNickname Searching User in database by Nickname
func (s *Store) GetUserByLogin(login string) (*models.User, error) {
	var user models.User

	rows, err := s.db.Query(`
		SELECT * FROM user WHERE nickname = ? OR email = ? 
	`, login, login)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Nickname, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Gender, &user.Age)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return &user, err
}
