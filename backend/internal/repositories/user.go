package repositories

import (
	"api/internal/models"
	"database/sql"
	"fmt"
)

// Users represents an users repositorie
type Users struct {
	db *sql.DB
}

// NewUsersRepositorie a new repositorie of users
func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

// Create insert a user in database
func (repository Users) Create(user models.User) (ID int, err error) {
	stmt, err := repository.db.Prepare(
		"INSERT INTO users (name, userName, email, password) VALUES ($1, $2, $3, $4) RETURNING id",
	)

	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(user.Name, user.UserName, user.Email, user.Password).Scan(&ID)
	if err != nil {
		return
	}

	return
}

// FilterByUserName fetch all the users that match the userName
func (repository Users) FilterByUserName(userName string) (users []models.User, err error) {
	userName = fmt.Sprintf("%%%s%%", userName)

	rows, err := repository.db.Query(
		"SELECT id, name, userName, email FROM users WHERE userName LIKE $1;", userName,
	)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.UserName,
			&user.Email,
		); err != nil {
			return
		}
		users = append(users, user)
	}

	return
}

// GetById fetch an user by id
func (repository Users) GetById(id int) (user models.User, err error) {
	row, err := repository.db.Query(
		"SELECT id, name, userName, email FROM users WHERE id=$1", id,
	)
	if err != nil {
		return
	}
	defer row.Close()

	if row.Next() {
		err = row.Scan(
			&user.ID,
			&user.Name,
			&user.UserName,
			&user.Email,
		)
		if err != nil {
			return
		}
	}

	return
}

// Update updates an user in database
func (repository Users) Update(user models.User) (err error) {
	stmt, err := repository.db.Prepare(
		"UPDATE users SET name=$1, userName=$2, email=$3 WHERE id=$4",
	)
	if err != nil {
		return
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Name, user.UserName, user.Email, user.ID); err != nil {
		return
	}

	return
}

// Delete deletes an user in database
func (repository Users) Delete(id int) (err error) {
	stmt, err := repository.db.Prepare("DELETE FROM users WHERE id=$1")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return
	}

	return
}

// SearchByEmail searchs an user using the email and returns ID and password of the user
func (repository Users) SearchByEmail(email string) (user models.User, err error) {
	row, err := repository.db.Query(
		"SELECT id, password FROM users WHERE email=$1", email,
	)
	if err != nil {
		return
	}
	defer row.Close()

	if row.Next() {
		err = row.Scan(&user.ID, &user.Password)
		if err != nil {
			return
		}
	}

	return
}

// Follow insert the user and the follower ID in the database, which represents the following action
func (repository Users) Follow(userID, followerID int) (err error) {
	stmt, err := repository.db.Prepare(
		"INSERT INTO followers (user_id, follower_id) VALUES ($1, $2)",
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, followerID)
	if err != nil {
		return
	}

	return
}

// Unfollow delete the user and the follower ID in the databasem, which represents the unfollowing action
func (repository Users) Unfollow(userID, followerID int) (err error) {
	stmt, err := repository.db.Prepare(
		"DELETE FROM followers WHERE user_id=$1 AND follower_id=$2",
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, followerID)
	if err != nil {
		return err
	}

	return
}

// GetFollowers get all the followers of a giver user in the database
func (repository Users) GetFollowers(userID int) (followers []models.User, err error) {
	rows, err := repository.db.Query(
		`
		SELECT U.id, U.name, U.username, U.email, U.creationdate
		FROM users U INNER JOIN followers F
		ON U.id = F.follower_id
		WHERE F.user_id = $1
		`, userID,
	)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var follower models.User
		if err = rows.Scan(
			&follower.ID,
			&follower.Name,
			&follower.UserName,
			&follower.Email,
			&follower.CreationDate,
		); err != nil {
			return
		}
		followers = append(followers, follower)
	}

	return
}
