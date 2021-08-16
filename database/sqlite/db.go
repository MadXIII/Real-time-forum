package sqlite

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	db *sql.DB
}

func Init(dbname string) (s *Store, err error) {
	fmt.Println(1)
	s = &Store{}

	s.db, err = sql.Open("sqlite3", dbname)
	if err != nil {
		return
	}

	userTable, err := s.db.Prepare(`CREATE TABLE IF NOT EXISTS user (
		id integer PRIMARY KEY NOT NULL,
		age integer NOT NULL,
		nickname VARCHAR(20) NOT NULL,
		gender VARCHAR(5),
		first_name VARCHAR(20),
		last_name VARCHAR(30),
		email VARCHAR(100) NOT NULL,
		password VARCHAR(255) NOT NULL
	);`)

	if err != nil {
		return
	}

	_, err = userTable.Exec()
	if err != nil {
		return
	}

	defer userTable.Close()

	postTable, err := s.db.Prepare(`CREATE TABLE IF NOT EXISTS post (
		id integer PRIMARY KEY NOT NULL,
		user_id integer NOT NULL,
		tittle VARCHAR(50) NOT NULL,
		content TEXT NOT NULL,
		timestamp TEXT NOT NULL,
		FOREIGN KEY (user_id) REFERENCES user(id)
	);`)

	if err != nil {
		return
	}

	_, err = postTable.Exec()
	if err != nil {
		return
	}

	defer postTable.Close()

	postLikeTable, err := s.db.Prepare(`CREATE TABLE IF NOT EXISTS postLike (
		user_id integer NOT NULL,
		post_id integer NOT NULL,
		like integer NOT NULL,
		FOREIGN KEY(user_id) REFERENCES user(id),
		FOREIGN KEY(post_id) REFERENCES post(id)
	);`)
	if err != nil {
		return
	}

	_, err = postLikeTable.Exec()
	if err != nil {
		return
	}

	defer postLikeTable.Close()

	commentTable, err := s.db.Prepare(`CREATE TABLE IF NOT EXISTS comment (
		id integer PRIMARY KEY NOT NULL,
		user_id integer NOT NULL,
		post_id integer NOT NULL,
		comment TEXT NOT NULL,
		timestamp TEXT,
		FOREIGN KEY(user_id) REFERENCES user(id),
		FOREIGN KEY(post_id) REFERENCES post(id)
		);`)

	_, err = commentTable.Exec()
	if err != nil {
		log.Println(23)
		return
	}

	defer commentTable.Close()

	commentLikeTable, err := s.db.Prepare(`CREATE TABLE IF NOT EXISTS commentLike (
		user_id integer NOT NULL,
		comment_id integer NOT NULL,
		like integer NOT NULL,
		FOREIGN KEY(user_id) REFERENCES user(id),
		FOREIGN KEY(comment_id) REFERENCES comment(id)
	);`)
	if err != nil {
		return
	}

	_, err = commentLikeTable.Exec()
	if err != nil {
		return
	}

	defer commentLikeTable.Close()

	fmt.Println(2)

	return
}
