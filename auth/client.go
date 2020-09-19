package auth

import (
	"net/http"

	"github.com/forstmeier/backend/config"
)

// Auth is the client struct hosting helper methods for interacting with
// the Auth0 API specifically for token management and metadata updates.
type Auth struct {
	client *http.Client
	config *config.Config
}

// New generates a pointer instance of the Auth client.
func New(cfg *config.Config) *Auth {
	return &Auth{
		client: &http.Client{},
		config: cfg,
	}
}
