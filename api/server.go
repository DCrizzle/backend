package api

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

func handleOrg(w http.ResponseWriter, r *http.Request) {}

func handleMutation(w http.ResponseWriter, r *http.Request) {}

func handleQuery(w http.ResponseWriter, r *http.Request) {}

type server struct {
	httpServer http.Server
}

func newServer(addr string) *server {

	router := mux.NewRouter()

	subrouter := router.Host(addr).Subrouter()

	// NOTE: probably refactor these paths/subrouters
	subrouter.HandleFunc("/org/{id}", handleOrg).Methods("GET")

	subrouter.HandleFunc("/org/{id}/db", handleMutation).Methods("POST")

	subrouter.HandleFunc("/org/{id}/db", handleQuery).Methods("GET")

	return &server{
		httpServer: http.Server{
			Addr:    addr,
			Handler: router,
		},
	}
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
