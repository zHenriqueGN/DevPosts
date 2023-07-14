package models

import (
	"errors"
	"strings"
	"time"
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
	user.format()
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
	if stage == "register" && user.Password == "" {
		return errors.New("the field Password can't be empty")
	}
	return
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.UserName = strings.TrimSpace(user.UserName)
	user.Email = strings.TrimSpace(user.Email)
}
