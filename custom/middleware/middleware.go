package middleware

import (
	"net/http"

	"github.com/forstmeier/backend/custom/handlers"
)

// Root hosts middleware logic and passes variables in as fields
type Root struct {
	folivoraSecret string
}

// New generates a pointer instance of the Root object
func New(folivoraSecret string) *Root {
	return &Root{
		folivoraSecret: folivoraSecret,
	}
}

// Middleware performs a header security check on all requests
func (root *Root) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if root.folivoraSecret != r.Header.Get("folivora-custom-secret") {
			http.Error(w, handlers.ErrIncorrectSecret, http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
