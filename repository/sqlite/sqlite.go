package sqlite

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/madxiii/real-time-forum/config"
)

func New(cfg config.DB) (*sqlx.DB, error) {
	db, err := sqlx.Open(cfg.Driver, cfg.Filename)
	if err != nil {
		return nil, fmt.Errorf("InitDB, sql.Open: %w", err)
	}

	log.Println("DB Creating...")

	userTable, err := db.Prepare(`CREATE TABLE IF NOT EXISTS user (
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
		return nil, fmt.Errorf("InitDB, userTable.Prepare: %w", err)
	}

	_, err = userTable.Exec()
	if err != nil {
		return nil, fmt.Errorf("InitDB, userTable.Exec: %w", err)
	}

	defer userTable.Close()

	signerTable, err := db.Prepare(`CREATE TABLE IF NOT EXISTS signer (
		id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
		nickname VARCHAR(100) NOT NULL,
		password VARCHAR(100) NOT NULL,
		FOREIGN KEY (nickname) REFERENCES user(nickname),
		FOREIGN KEY (password) REFERENCES user(password)
	);`)
	if err != nil {
		return nil, fmt.Errorf("InitDB, signerTable.Prepare: %w", err)
	}

	_, err = signerTable.Exec()
	if err != nil {
		return nil, fmt.Errorf("InitDB, signerTable.Exec: %w", err)
	}

	defer signerTable.Close()

	// categoryTable, err := db.Prepare(`CREATE TABLE IF NOT EXISTS category (
	// 	id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
	// 	name VARCHAR(50) NOT NULL UNIQUE
	// );`)
	// if err != nil {
	// 	return nil, fmt.Errorf("InitDB, categoryTable.Prepare: %w", err)
	// }

	// _, err = categoryTable.Exec()
	// if err != nil {
	// 	return nil, fmt.Errorf("InitDB, categoryTabel.Exec: %w", err)
	// }

	// categories := []string{"All", "UFC", "Anime", "Other"}

	// if err = InsertCategories(categories); err != nil {
	// 	return nil, fmt.Errorf("InitDB, InsertCategories: %w", err)
	// }

	postTable, err := db.Prepare(`CREATE TABLE IF NOT EXISTS post (
		id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
		category_id VARCHAR(50) NOT NULL,
		username VARCHAR(100) NOT NULL,
		title VARCHAR(50) NOT NULL,
		content TEXT NOT NULL,
		timestamp TEXT NOT NULL,
		like_count integer NOT NULL,
		FOREIGN KEY (category_id) REFERENCES category(id),
		FOREIGN KEY (username) REFERENCES user(nickname)
	);`)
	if err != nil {
		return nil, fmt.Errorf("InitDB, postTable.Prepare: %w", err)
	}

	_, err = postTable.Exec()
	if err != nil {
		return nil, fmt.Errorf("InitDB, postTable.Exec: %w", err)
	}

	defer postTable.Close()

	postLikeTable, err := db.Prepare(`CREATE TABLE IF NOT EXISTS postLike (
		user_id integer NOT NULL,
		post_id integer NOT NULL,
		like integer NOT NULL,
		FOREIGN KEY(user_id) REFERENCES user(id),
		FOREIGN KEY(post_id) REFERENCES post(id)
	);`)
	if err != nil {
		return nil, fmt.Errorf("InitDB, postLikeTable.Prepare: %w", err)
	}

	_, err = postLikeTable.Exec()
	if err != nil {
		return nil, fmt.Errorf("InitDB, postLikeTable.Exec: %w", err)
	}

	defer postLikeTable.Close()

	commentTable, err := db.Prepare(`CREATE TABLE IF NOT EXISTS comment (
		id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
		post_id integer NOT NULL,
		username VARCHAR(100) NOT NULL,
		content TEXT NOT NULL,
		timestamp TEXT NOT NULL,
		FOREIGN KEY(post_id) REFERENCES post(id),
		FOREIGN KEY (username) REFERENCES user(nickname)
		);`)
	if err != nil {
		return nil, fmt.Errorf("InitDB, commentTable.Prepare: %w", err)
	}

	_, err = commentTable.Exec()
	if err != nil {
		return nil, fmt.Errorf("InitDB, commentTable.Exec: %w", err)
	}

	defer commentTable.Close()

	commentLikeTable, err := db.Prepare(`CREATE TABLE IF NOT EXISTS commentLike (
		user_id integer NOT NULL,
		comment_id integer NOT NULL,
		like integer NOT NULL,
		FOREIGN KEY(user_id) REFERENCES user(id),
		FOREIGN KEY(comment_id) REFERENCES comment(id)
	);`)
	if err != nil {
		return nil, fmt.Errorf("InitDB, commentLikeTable.Prepare: %w", err)
	}

	_, err = commentLikeTable.Exec()
	if err != nil {
		return nil, fmt.Errorf("InitDB, commentLikeTable.Exec: %w", err)
	}

	defer commentLikeTable.Close()

	log.Println("DB Created")
	return db, nil
}
