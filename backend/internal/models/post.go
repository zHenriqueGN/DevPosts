package models

import (
	"errors"
	"strings"
	"time"
)

// Post represents a post created by an user
type Post struct {
	ID             int        `json:"id,omitempty"`
	Title          string     `json:"title,omitempty"`
	Content        string     `json:"content,omitempty"`
	AuthorID       int        `json:"authorId,omitempty"`
	AuthorUserName string     `json:"authorUserName,omitempty"`
	Likes          int        `json:"likes"`
	CreationDate   *time.Time `json:"creationDate,omitempty"`
}

// Prepare call the methods to validate and format the User
func (post *Post) Prepare() (err error) {
	if err = post.validate(); err != nil {
		return
	}

	post.format()

	return
}

func (post *Post) validate() (err error) {
	if post.Title == "" {
		return errors.New("the field title can't be empty")
	}
	if post.Content == "" {
		return errors.New("the field content can't be empty")
	}
	return
}

func (post *Post) format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}
