package backend

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {}

func NewServer(gql graphql) *Server {

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

func middleware(http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})

	// outline:
	// [ ] return wrapped handler function
	// [ ] retrieve token from request ("authorization" header)
	// [ ] validate/parse token
	// [ ] server request with context (?)

}

type graphql interface {
	query(query string, variables, headers map[string]string) io.ReadCloser
	mutation(mutation string, variables, headers map[string]string) io.ReadCloser
}
