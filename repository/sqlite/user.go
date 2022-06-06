package sqlite

import (
	"fmt"

	"github.com/madxiii/real-time-forum/model"

	"github.com/jmoiron/sqlx"
)

type User struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *User {
	return &User{db: db}
}

// CreateUser - Insert NewUser in db
func (u *User) CreateUser(user *model.User) error {
	createTable, err := u.db.Prepare(`
	INSERT INTO user
	(nickname, email, password, first_name, last_name, gender, age)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return fmt.Errorf("CreateUser, Prepare: %w", err)
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
		return fmt.Errorf("CreateUser, Exec: %w", err)
	}

	userid, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("CreateUser, LastInsertId: %w", err)
	}
	user.ID = int(userid)

	return nil
}

// GetUserByLogin - Searching User in db by Nickname
func (u *User) GetUserByLogin(login string) (model.User, error) {
	var user model.User

	err := u.db.QueryRow(`
	SELECT * FROM user WHERE nickname = ? OR email = ?
	`, login, login).
		Scan(&user.ID, &user.Nickname, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Gender, &user.Age)
	if err != nil {
		return user, fmt.Errorf("GetUserByLogin, Scan: %w", err)
	}

	return user, nil
}

// GetUsernameByID - Get username from db by userID
func (u *User) GetUsernameByID(id int) (string, error) {
	var username string

	err := u.db.QueryRow(`
	SELECT nickname FROM user WHERE id = ?
	`, id).
		Scan(&username)
	if err != nil {
		return "", fmt.Errorf("GetUsernameByID, Scan: %w", err)
	}

	return username, nil
}
