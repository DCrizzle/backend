package app

import "github.com/google/uuid"

// outline:
// [ ] possibly create operation interface
// - [ ] sub interfaces:
// - - [ ] mutation
// - - [ ] query

// ID is the application-wide unique object identifier
type ID uuid.UUID

// Mutation is the input object to the database Mutate method
type Mutation struct {
	ID      *ID    `json:"id,omitempty"`
	Content []byte `json:"content,omitempty"`
}

// Query is the input object to the database Query method
type Query struct {
	ID        *ID               `json:"id,omitempty"`
	Content   *string           `json:"content,omitempty"`
	Variables map[string]string `json:"content,omitempty"`
}
