package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

// Server hosts the backend server with login/logout and GraphQL endpoints.
type Server struct {
	httpServer *http.Server
	params     *params
}

// NewServer generates a pointer to an inactive Server instance.
func NewServer(configPath string) (*Server, error) {
	router := mux.NewRouter()

	params, err := parseParams(configPath)
	if err != nil {
		return nil, err
	}

	if err := params.validate(); err != nil {
		return nil, err
	}

	gql := newGraphQLClient(&http.Client{
		Timeout: 10 * time.Second,
	})

	router.HandleFunc("/graphql", graphQLHandler(params.DgraphURL, gql))

	csrf.Protect([]byte(params.CSRFKey))(router)

	return &Server{
		httpServer: &http.Server{
			Addr:         "127.0.0.1:8888",
			Handler:      router,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		params: params,
	}, nil
}

// Start begins serving the http.Server.
func (s *Server) Start() {
	s.httpServer.ListenAndServe()
}

// Stop ends serving the http.Server.
func (s *Server) Stop(ctx context.Context) {
	s.httpServer.Shutdown(ctx)
}
