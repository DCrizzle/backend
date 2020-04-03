package api

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type helper struct{}

func (h *helper) secure(hdlr http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		_ = token

		// outline:
		// [ ] query database for user
		// [ ] return success/error based on response

		hdlr.ServeHTTP(w, r)
	})
}

func mutate() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func query() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

type server struct {
	httpServer http.Server
}

func newServer(addr string) (*server, error) {
	router := mux.NewRouter()

	// // outline:
	// // [ ] add non-auth handlers here

	subrouter := router.Host(addr).Subrouter()

	h := &helper{}

	subrouter.Use(h.secure)

	subrouter.HandleFunc("/graphql", mutate()).Methods("POST")

	subrouter.HandleFunc("/graphql", query()).Methods("GET")

	return &server{
		httpServer: http.Server{
			Addr:    addr,
			Handler: router,
		},
	}, nil
}

func (s *server) start() error {
	return s.httpServer.ListenAndServe()
}

func (s *server) stop() error {
	ctx := context.Background()
	return s.httpServer.Shutdown(ctx)
}

// outline:
// [ ] handlers
// - [ ] "/" -> redirect to "/login"
// - [ ] "/login" -> validate user login; redirect to "/org/{id}" (POST)
// - [ ] "/org/{id}" -> render org (GET)
// - [ ] "/org/{id}/db" -> execute mutations (POST)
// - [ ] "/org/{id}/db" -> execute queries (GET)
// - [ ] "/org/{id}/admin" -> org settings (GET)
// - [ ] "/org/{id}/admin/users" -> users settings (GET)
// - [ ] "/org/{id}/admin/user/{id}" -> user settings (GET)
// [ ] new server method
// [ ] start server method
// [ ] stop server method
