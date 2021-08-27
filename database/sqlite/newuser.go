package sqlite

import (
	"forum/models"
	"log"
)

func (s *Store) InsertUser(user models.User) (err error) {
	createTable, err := s.db.Prepare(`
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
		log.Fatal(err)
		return
	}

	userid, err := res.LastInsertId()
	user.ID = int(userid)

	if err != nil {
		log.Fatal(err)
		return
	}
	return
}
