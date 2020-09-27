package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type server struct {
	httpServer *http.Server
}

func newServer(usersHandler http.HandlerFunc) *server {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/auth0").Subrouter()
	subrouter.HandleFunc("/users", usersHandler)

	customServer := &http.Server{
		Addr:         "127.0.0.1:4080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return &server{
		httpServer: customServer,
	}
}

func (s *server) start() {
	s.httpServer.ListenAndServe()
}

func (s *server) stop(ctx context.Context) {
	s.httpServer.Shutdown(ctx)
}
