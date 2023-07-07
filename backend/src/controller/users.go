package controller

import (
	"api/src/database"
	"api/src/messages"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io"
	"net/http"
)

// CreateUser create a user in database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		messages.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err := json.Unmarshal(requestBody, &user); err != nil {
		messages.Error(w, http.StatusBadRequest, err)
	}

	db, err := database.ConnectToDB()
	if err != nil {
		messages.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepositorie(db)
	ID, err := repository.Create(user)
	if err != nil {
		messages.Error(w, http.StatusInternalServerError, err)
		return
	}
	user.ID = ID

	messages.JSON(w, http.StatusCreated, user)
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
