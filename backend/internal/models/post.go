package models

import "time"

// Post represents a post created by an user
type Post struct {
	ID             int        `json:"id,omitempty"`
	Title          int        `json:"title,omitempty"`
	Content        int        `json:"content,omitempty"`
	AuthorID       int        `json:"authorId,omitempty"`
	AuthorUserName int        `json:"authorUserName,omitempty"`
	Likes          int        `json:"likes"`
	CreationDate   *time.Time `json:"creationDate,omitempty"`
}
