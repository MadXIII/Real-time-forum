package sqlite

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//Store - store hole of DB
type Store struct {
	db *sql.DB
}

//Init - Creat DB if not created
func (s *Store) Init(dbname string) (err error) {
	s.db, err = sql.Open("sqlite3", dbname)
	if err != nil {
		return
	}
	log.Println("DB creating...")
	userTable, err := s.db.Prepare(`CREATE TABLE IF NOT EXISTS user (
		id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
		nickname VARCHAR(20) NOT NULL UNIQUE,
		email VARCHAR(100) NOT NULL UNIQUE,
		password BLOB NOT NULL,
		first_name VARCHAR(20),
		last_name VARCHAR(30),
		gender VARCHAR(5),
		age integer
	);`)

	if err != nil {
		return
	}

	_, err = userTable.Exec()
	if err != nil {
		return
	}

	defer userTable.Close()

	signerTable, err := s.db.Prepare(`CREATE TABLE IF NOT EXISTS signer (
		id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
		nickname VARCHAR(20) NOT NULL,
		password VARCHAR(32) NOT NULL,
		FOREIGN KEY (nickname) REFERENCES user(nickname),
		FOREIGN KEY (password) REFERENCES user(password)
	);`)

	if err != nil {
		return
	}

	_, err = signerTable.Exec()
	if err != nil {
		return
	}

	defer signerTable.Close()

	postTable, err := s.db.Prepare(`CREATE TABLE IF NOT EXISTS post (
		id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
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
		id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
		user_id integer NOT NULL,
		post_id integer NOT NULL,
		comment TEXT NOT NULL,
		timestamp TEXT,
		FOREIGN KEY(user_id) REFERENCES user(id),
		FOREIGN KEY(post_id) REFERENCES post(id)
		);`)

	_, err = commentTable.Exec()
	if err != nil {
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
	log.Println("DB Created")
	return
}
