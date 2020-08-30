package auth0

import (
	"net/http"

	"github.com/forstmeier/backend/config"
)

// Auth0 is the client struct hosting helper methods for interacting with
// the Auth0 API specifically for token management and metadata updates.
type Auth0 struct {
	client *http.Client
	config *config.Config
}

// New generates a pointer instance of the Auth0 client.
func New(cfg *config.Config) *Auth0 {
	return &Auth0{
		client: &http.Client{},
		config: cfg,
	}
}
