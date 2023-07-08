package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
	"log"
)

// Users represents an users repositorie
type Users struct {
	db *sql.DB
}

// NewUsersRepositorie a new repositorie of users
func NewUsersRepositorie(db *sql.DB) *Users {
	return &Users{db}
}

// Create insert a user in database
func (repository Users) Create(user models.User) (ID uint, err error) {
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
	log.Println(userName)
	rows, err := repository.db.Query(
		"SELECT id, name, userName, email, password, creationDate FROM users WHERE userName LIKE $1;", userName,
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
			&user.Password,
			&user.CriationDate,
		); err != nil {
			return
		}
		users = append(users, user)
	}

	return
}
