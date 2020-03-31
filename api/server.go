package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/forstmeier/tbd/app"
)

const (
	errNewDB = "error calling new db function"
)

type apiHandler func(db app.Database) func(w http.ResponseWriter, r *http.Request)

func contextWrap(handler func(db app.Database) func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	// outline:
	// [ ] create context
	// [ ] validate user
	// [ ] insert into context with values
	// - [ ] handlers: handleOrg, handleMutation, handleQuery
	// [ ] pass into received handler
	// [ ] return/invoke handler
})

func handleOrg(db app.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pathSplit := strings.Split(r.URL.Path, "/")
		_ = pathSplit
	}

	// outline:
	// [ ] get user from request
	// - [ ] metadata (?)
	// - [ ] url parameters
	// [ ] validate user access
	// [ ] create database query + execute
	// [ ] create return object + return
}

func handleMutation(db app.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pathSplit := strings.Split(r.URL.Path, "/")
		_ = pathSplit
	}

	// outline:
	// [ ] get user from request
	// - [ ] metadata (?)
	// - [ ] url parameters
	// [ ] validate user access
	// [ ] create database query + execute
	// [ ] create return object + return
}

func handleQuery(db app.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pathSplit := strings.Split(r.URL.Path, "/")
		_ = pathSplit
	}

	// outline:
	// [ ] get user from request
	// - [ ] metadata (?)
	// - [ ] url parameters
	// [ ] validate user access
	// [ ] create database query + execute
	// [ ] create return object + return
}

type server struct {
	httpServer http.Server
	database   app.Database
}

func newServer(addr string) (*server, error) {

	router := mux.NewRouter()

	subrouter := router.Host(addr).Subrouter()

	db, err := app.NewDB(addr)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errNewDB, err)
	}

	// NOTE: probably refactor these paths/subrouters
	subrouter.HandleFunc("/org/{id}", handleOrg(db)).Methods("GET")

	subrouter.HandleFunc("/org/{id}/db", handleMutation(db)).Methods("POST")

	subrouter.HandleFunc("/org/{id}/db", handleQuery(db)).Methods("GET")

	return &server{
		httpServer: http.Server{
			Addr:    addr,
			Handler: router,
		},
		database: db,
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
