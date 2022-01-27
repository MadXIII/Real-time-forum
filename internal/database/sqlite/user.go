package sqlite

import (
	"fmt"
	"forum/internal/models"
)

//InsertUser - Insert NewUser in db
func (s *Store) InsertUser(user *models.User) error {
	createTable, err := s.db.Prepare(`
	INSERT INTO user
	(nickname, email, password, first_name, last_name, gender, age)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return fmt.Errorf("InsertUser, Prepare: %w", err)
	}

	res, err := createTable.Exec(
		user.Nickname,
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
		user.Gender,
		user.Age,
	)

	if err != nil {
		return fmt.Errorf("InsertUser, Exec: %w", err)
	}

	userid, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("InsertUser, LastInsertId: %w", err)
	}
	user.ID = int(userid)

	return nil
}

//GetUserByLogin - Searching User in db by Nickname
func (s *Store) GetUserByLogin(login string) (models.User, error) {
	var user models.User

	err := s.db.QueryRow(`
	SELECT * FROM user WHERE nickname = ? OR email = ?
	`, login, login).
		Scan(&user.ID, &user.Nickname, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Gender, &user.Age)

	if err != nil {
		return user, fmt.Errorf("GetUserByLogin, Scan: %w", err)
	}

	return user, nil
}

//GetUsernameByID - Get username from db by userID
func (s *Store) GetUsernameByID(id int) (string, error) {
	var username string

	err := s.db.QueryRow(`
	SELECT nickname FROM user WHERE id = ?
	`, id).
		Scan(&username)

	if err != nil {
		return "", fmt.Errorf("GetUsernameByID, Scan: %w", err)
	}

	return username, nil
}
