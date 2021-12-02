package sqlite

import (
	"database/sql"
	"fmt"
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
		return fmt.Errorf("InitDB, sql.Open: %w", err)
	}

	log.Println("DB Creating...")

	userTable, err := s.db.Prepare(`CREATE TABLE IF NOT EXISTS user (
		id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
		nickname VARCHAR(100) NOT NULL UNIQUE,
		email VARCHAR(100) NOT NULL UNIQUE,
		password BLOB NOT NULL,
		first_name VARCHAR(100),
		last_name VARCHAR(100),
		gender VARCHAR(100),
		age VARCHAR(50)
	);`)
	if err != nil {
		return fmt.Errorf("InitDB, userTable.Prepare: %w", err)
	}

	_, err = userTable.Exec()
	if err != nil {
		return fmt.Errorf("InitDB, userTable.Exec: %w", err)
	}

	defer userTable.Close()

	signerTable, err := s.db.Prepare(`CREATE TABLE IF NOT EXISTS signer (
		id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
		nickname VARCHAR(100) NOT NULL,
		password VARCHAR(100) NOT NULL,
		FOREIGN KEY (nickname) REFERENCES user(nickname),
		FOREIGN KEY (password) REFERENCES user(password)
	);`)
	if err != nil {
		return fmt.Errorf("InitDB, signerTable.Prepare: %w", err)
	}

	_, err = signerTable.Exec()
	if err != nil {
		return fmt.Errorf("InitDB, signerTable.Exec: %w", err)
	}

	defer signerTable.Close()

	postTable, err := s.db.Prepare(`CREATE TABLE IF NOT EXISTS post (
		id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
		username VARCHAR(100) NOT NULL,
		title VARCHAR(50) NOT NULL,
		content TEXT NOT NULL,
		timestamp TEXT NOT NULL,
		like_count integer NOT NULL,
		dis_count integer NOT NULL,
		FOREIGN KEY (username) REFERENCES user(nickname)
	);`)
	if err != nil {
		return fmt.Errorf("InitDB, postTable.Prepare: %w", err)
	}

	_, err = postTable.Exec()
	if err != nil {
		return fmt.Errorf("InitDB, postTable.Exec: %w", err)
	}

	defer postTable.Close()

	postLikeTable, err := s.db.Prepare(`CREATE TABLE IF NOT EXISTS postLike (
		user_id integer NOT NULL,
		post_id integer NOT NULL,
		like integer NOT NULL,
		type VARCHAR(50) NOT NULL,
		FOREIGN KEY(user_id) REFERENCES user(id),
		FOREIGN KEY(post_id) REFERENCES post(id)
	);`)
	if err != nil {
		return fmt.Errorf("InitDB, postLikeTable.Prepare: %w", err)
	}

	_, err = postLikeTable.Exec()
	if err != nil {
		return fmt.Errorf("InitDB, postLikeTable.Exec: %w", err)
	}

	defer postLikeTable.Close()

	commentTable, err := s.db.Prepare(`CREATE TABLE IF NOT EXISTS comment (
		id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
		post_id integer NOT NULL,
		username VARCHAR(100) NOT NULL,
		content TEXT NOT NULL,
		timestamp TEXT NOT NULL,
		FOREIGN KEY(post_id) REFERENCES post(id),
		FOREIGN KEY (username) REFERENCES user(nickname)
		);`)
	if err != nil {
		return fmt.Errorf("InitDB, commentTable.Prepare: %w", err)
	}

	_, err = commentTable.Exec()
	if err != nil {
		return fmt.Errorf("InitDB, commentTable.Exec: %w", err)
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
		return fmt.Errorf("InitDB, commentLikeTable.Prepare: %w", err)
	}

	_, err = commentLikeTable.Exec()
	if err != nil {
		return fmt.Errorf("InitDB, commentLikeTable.Exec: %w", err)
	}

	defer commentLikeTable.Close()

	log.Println("DB Created")
	return
}
