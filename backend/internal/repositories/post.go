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

// Update updates a post in database
func (repository Posts) Update(post models.Post) (err error) {
	stmt, err := repository.db.Prepare(
		`UPDATE posts SET title=$1, content=$2 WHERE id=$3`,
	)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(post.Title, post.Content, post.ID)
	if err != nil {
		return
	}

	return
}

// Delete deletes a post in database
func (repository Posts) Delete(ID int) (err error) {
	stmt, err := repository.db.Prepare("DELETE FROM posts WHERE id=$1")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(ID)
	if err != nil {
		return
	}

	return
}
