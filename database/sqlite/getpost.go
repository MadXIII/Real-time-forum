// package sqlite

// import "forum/models"

// func (s *Store) GetPostByID(id int) (*models.Post, error) {
// 	var post models.Post
// 	rows, err := s.db.Query(`
// 		SELECT * FROM post WHERE id = ?
// 	`, id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		err := rows.Scan()
// 	}

// 	return &models.Post{}, err
// }
