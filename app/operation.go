package app

import "github.com/google/uuid"

// outline:
// [ ] possibly create operation interface
// - [ ] sub interfaces:
// - - [ ] mutation
// - - [ ] query

type id uuid.UUID

type mutation struct {
	ID      *id    `json:"id,omitempty"`
	Content []byte `json:"content,omitempty"`
}

type query struct {
	ID        *id               `json:"id,omitempty"`
	Content   *string           `json:"content,omitempty"`
	Variables map[string]string `json:"content,omitempty"`
}

type response struct{}
