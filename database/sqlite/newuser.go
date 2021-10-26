package sqlite

import (
	"forum/models"
	"log"
)

//InsertUser - Insert user in DB
func (s *Store) InsertUser(user models.User) error {
	createTable, err := s.db.Prepare(`
	INSERT INTO user
	(nickname, email, password, first_name, last_name, gender, age)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		log.Println(err)
		return err
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
		return err
	}

	userid, err := res.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int(userid)

	return nil
}
