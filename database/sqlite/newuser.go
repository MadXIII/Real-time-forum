package sqlite

import (
	"forum/models"
	"log"
)

func (s *Store) InsertUser(user models.User) (err error) {
	createTable, err := s.db.Prepare(`
	INSERT INTO user
	(nickname, email, password, first_name, last_name, gender, age)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		log.Println(err)
		return
	}
	defer createTable.Close()

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
		log.Println(err)
		return
	}

	userid, err := res.LastInsertId()
	user.ID = int(userid)

	if err != nil {
		log.Println(err)
		return
	}
	return
}
