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
