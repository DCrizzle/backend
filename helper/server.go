package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func getAuth0ManagementAPIToken(client *http.Client) string {

	// outline:
	// [ ] invoke auth0 api
	// [ ] parse returned data
	// [ ] return jwt string value

	return ""
}

type server struct {
	auth0ManagementAPIToken string
	httpServer              *http.Server
}

func newServer(auth0ManagementAPIToken string, usersHandler http.HandlerFunc) *server {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/auth0").Subrouter()
	subrouter.HandleFunc("/users", usersHandler)

	helperServer := &http.Server{
		Addr:         "127.0.0.1:8888",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return &server{
		auth0ManagementAPIToken: auth0ManagementAPIToken,
		httpServer:              helperServer,
	}

}

func (s *server) start() {
	s.httpServer.ListenAndServe()
}

func (s *server) stop(ctx context.Context) {
	s.httpServer.Shutdown(ctx)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {

}
