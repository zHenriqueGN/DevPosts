package controller

import (
	"api/internal/database"
	"api/internal/messages"
	"api/internal/models"
	"api/internal/repositories"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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
		return
	}
	if err = user.Prepare("register"); err != nil {
		messages.Error(w, http.StatusBadRequest, err)
		return
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
	userName := strings.ToLower(r.URL.Query().Get("username"))

	db, err := database.ConnectToDB()
	if err != nil {
		messages.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepositorie(db)
	users, err := repository.FilterByUserName(userName)
	if err != nil {
		messages.Error(w, http.StatusInternalServerError, err)
		return
	}

	messages.JSON(w, http.StatusOK, users)
}

// FetchUser fetch an user in database
func FetchUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		messages.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.ConnectToDB()
	if err != nil {
		messages.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepositorie(db)
	user, err := repository.GetById(id)
	if err != nil {
		messages.Error(w, http.StatusInternalServerError, err)
		return
	}

	if user.ID != id {
		messages.JSON(w, http.StatusNotFound, nil)
		return
	}

	messages.JSON(w, http.StatusOK, user)
}

// UpdateUser update an user in database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		messages.Error(w, http.StatusBadRequest, err)
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		messages.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err = json.Unmarshal(requestBody, &user); err != nil {
		messages.Error(w, http.StatusBadRequest, err)
		return
	}
	user.ID = id

	if err = user.Prepare("update"); err != nil {
		messages.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.ConnectToDB()
	if err != nil {
		messages.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepositorie(db)

	user, err = repository.GetById(id)
	if err != nil {
		messages.Error(w, http.StatusInternalServerError, err)
		return
	}

	if user.ID != id {
		messages.JSON(w, http.StatusNotFound, nil)
		return
	}

	if err := repository.Update(user); err != nil {
		messages.Error(w, http.StatusInternalServerError, err)
		return
	}

	messages.JSON(w, http.StatusNoContent, nil)
}

// DeleteUser delete an user from database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		messages.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.ConnectToDB()
	if err != nil {
		messages.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepositorie(db)
	user, err := repository.GetById(id)
	if err != nil {
		messages.Error(w, http.StatusInternalServerError, err)
		return
	}

	if user.ID != id {
		messages.JSON(w, http.StatusNotFound, nil)
		return
	}

	err = repository.Delete(id)
	if err != nil {
		messages.Error(w, http.StatusInternalServerError, err)
		return
	}

	messages.JSON(w, http.StatusNoContent, nil)
}
