package models

import "time"

// User represents a user
type User struct {
	ID           uint      `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	UserName     string    `json:"userName,omitempty"`
	Email        string    `json:"email,omitempty"`
	Password     string    `json:"password,omitempty"`
	CriationDate time.Time `json:"creationDate,omitempty"`
}
