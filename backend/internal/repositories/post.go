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

// Fetch search for all user related posts in database
func (repository Posts) Fetch(userID int) (posts []models.Post, err error) {
	rows, err := repository.db.Query(
		`
		SELECT DISTINCT P.id, P.title, P.content, P.author_id, U.username, P.likes, P.creation_date
		FROM posts P
		INNER JOIN users U
				ON
			P.author_id = U.id
		INNER JOIN followers f 
				ON
			P.author_id = F.user_id
		WHERE U.id = $1 OR F.follower_id = $2
		ORDER BY P.id DESC
		`, userID, userID,
	)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.AuthorUserName,
			&post.Likes,
			&post.CreationDate,
		); err != nil {
			return
		}
		posts = append(posts, post)
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

// GetByUserId get all the posts by an user ID
func (repository Posts) GetByUserId(userID int) (posts []models.Post, err error) {
	rows, err := repository.db.Query(
		`SELECT P.id, 
				P.title, 
				P.content, 
				P.author_id,
				U.username, 
				P.likes, 
				P.creation_date
		FROM posts P INNER JOIN users U
		ON P.author_id = U.id
		WHERE P.author_id = $1
		ORDER BY P.id DESC`, userID,
	)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.AuthorUserName,
			&post.Likes,
			&post.CreationDate,
		); err != nil {
			return
		}
		posts = append(posts, post)
	}

	return
}

// Like increments the like count of a post
func (repository Posts) Like(ID int) (err error) {
	stmt, err := repository.db.Prepare("UPDATE posts SET likes=likes+1 WHERE id=$1")
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
