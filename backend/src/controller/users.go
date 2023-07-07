package controller

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// CreateUser create a user in database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err := json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repository := repositories.NewUsersRepositorie(db)
	ID, err := repository.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("New user created with ID %d", ID)))
}

// FetchUsers fetch all the users in database
func FetchUsers(w http.ResponseWriter, r *http.Request) {

}

// FetchUser fetch an user in database
func FetchUser(w http.ResponseWriter, r *http.Request) {

}

// UpdateUser update an user in database
func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

// DeleteUser delete an user from database
func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
