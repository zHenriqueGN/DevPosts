package repositories

import (
	"api/internal/models"
	"database/sql"
)

// Users represents a posts repository
type Posts struct {
	db *sql.DB
}

// NewUsersRepository create a new repository of posts
func NewPostsRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

// Create insert a post in database
func (repository Posts) Create(post models.Post) (ID int, err error) {
	stmt, err := repository.db.Prepare("INSERT INTO posts (title, content, author_id) VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(post.Title, post.Content, post.AuthorID).Scan(&ID)
	if err != nil {
		return
	}

	return
}

// GetById fetch a post by id
func (repository Posts) GetById(ID int) (post models.Post, err error) {
	row, err := repository.db.Query(
		`SELECT P.id, 
				P.title, 
				P.content, 
				P.author_id,
				U.username, 
				P.likes, 
				P.creation_date
		FROM posts P INNER JOIN users U
		ON P.author_id = U.id
		WHERE P.id = $1`, ID,
	)
	if err != nil {
		return
	}
	defer row.Close()

	if row.Next() {
		err = row.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.AuthorUserName,
			&post.Likes,
			&post.CreationDate,
		)
		if err != nil {
			return
		}
	}

	return
}
