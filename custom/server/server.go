package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Server holds the custom server router and exposes the
// required helper methods
type Server struct {
	httpServer *http.Server
}

// New generates a pointer instance of the Server object
func New(usersHandler http.HandlerFunc) *Server {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/auth0").Subrouter()
	subrouter.HandleFunc("/users", usersHandler)

	customServer := &http.Server{
		Addr:         "127.0.0.1:4080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return &Server{
		httpServer: customServer,
	}
}

// Start starts the configured server
func (s *Server) Start() {
	s.httpServer.ListenAndServe()
}

// Stop stops the configured server
func (s *Server) Stop(ctx context.Context) {
	s.httpServer.Shutdown(ctx)
}
