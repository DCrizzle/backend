package backend

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {}

func NewServer() *Server {

	router := mux.NewRouter()

	router.Use(middleware)

	// outline:
	// [x] declare mux router
	// [ ] add authentication middleware
	// [ ] add graphql handler
	// [ ] configure/add http server
	// [ ] return server pointer

	return &Server{}
}

func middleware(client dgrapher) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {





		})

	}

	// outline:
	// [ ] return wrapped handler function
	// [ ] retrieve token from request
	// [ ] validate/parse token
	// [ ] query database with dgraph client
	// [ ] hydrate full user object
	// [ ] add user object to request context
	// [ ] server request with context

}

type dgrapher interface {
// 	Query(query string, variables map[string]string) ([]byte, error)
}

// type dgraph struct{}
//
// func (d *dgraph) Query(query string, variables map[string]string) ([]byte, error) {
//
// }
