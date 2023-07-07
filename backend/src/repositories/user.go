package repositories

import (
	"api/src/models"
	"database/sql"
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
