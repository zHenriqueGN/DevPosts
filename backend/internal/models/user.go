package models

import (
	"api/internal/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User represents a user
type User struct {
	ID           int        `json:"id,omitempty"`
	Name         string     `json:"name,omitempty"`
	UserName     string     `json:"userName,omitempty"`
	Email        string     `json:"email,omitempty"`
	Password     string     `json:"password,omitempty"`
	CreationDate *time.Time `json:"creationDate,omitempty"`
}

// Prepare call the methods to validate and format the User
func (user *User) Prepare(stage string) (err error) {
	if err = user.validate(stage); err != nil {
		return
	}
	if err = user.format(stage); err != nil {
		return
	}
	return
}

func (user *User) validate(stage string) (err error) {
	if user.Name == "" {
		return errors.New("the field Name can't be empty")
	}
	if user.UserName == "" {
		return errors.New("the field UserName can't be empty")
	}
	if user.Email == "" {
		return errors.New("the field Email can't be empty")
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("invalid e-mail")
	}
	if stage == "register" && user.Password == "" {
		return errors.New("the field Password can't be empty")
	}
	return
}

func (user *User) format(stage string) (err error) {
	user.Name = strings.TrimSpace(user.Name)
	user.UserName = strings.TrimSpace(user.UserName)
	user.Email = strings.TrimSpace(user.Email)

	if stage == "register" {
		passwordWithHash, err := security.HashPassword(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(passwordWithHash)
	}
	return
}
