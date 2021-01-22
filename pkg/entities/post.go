package entities

import "github.com/google/uuid"

// Post is model for working with posts
type Post struct {
	ID    uuid.UUID `json:"id,omitempty"`
	Title string    `json:"title,omitempty"`
	Body  string    `json:"body,omitempty"`
}
