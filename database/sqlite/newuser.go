package sqlite

import (
	"forum/database"
	"forum/models"
)

type User struct {
	store *Store
}

func (u User) InsertUser(user models.User, s database.Repository) (err error) {
	createTable, err := u.store.db.Prepare(`
	INSERT INTO user
	(age, nickname, gender, first_name, last_name, email, password)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return
	}

	defer createTable.Close()

	res, err := createTable.Exec(
		user.Age,
		user.Nickname,
		user.Gender,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
	)

	if err != nil {
		return
	}

	userid, err := res.LastInsertId()
	user.ID = int(userid)

	if err != nil {
		return
	}
	return
}
