package server

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

const (
	errMutationRequest = "server: error invoking backend database mutation"
	errQueryRequest    = "server: error invoking backend database query"
)

// Server hosts the backend server with login/logout and GraphQL endpoints.
type Server struct {
	httpServer *http.Server
	params     *params
}

// NewServer generates a pointer to an inactive Server instance.
func NewServer(configPath string, gql graphQL) (*Server, error) {
	router := mux.NewRouter()

	params, err := parseParams(configPath)
	if err != nil {
		return nil, err
	}

	if err := params.validate(); err != nil {
		return nil, err
	}

	// note: pass params struct into middleware wrapper
	// router.Use(middleware)
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

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func graphQLHandler(dgraphURL string, gql graphQL) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var output io.ReadCloser

		if r.Method == http.MethodPost {
			response, err := gql.mutation(dgraphURL, r.Body)
			if err != nil {
				http.Error(w, fmt.Sprintf(`{"message":"%s"}`, errMutationRequest), http.StatusInternalServerError)
				return
			}

			output = response.Body

		} else if r.Method == http.MethodGet {
			queryURL := dgraphURL + r.URL.Path + "?" + r.URL.RawQuery
			response, err := gql.query(queryURL)
			if err != nil {
				http.Error(w, fmt.Sprintf(`{"message":"%s"}`, errQueryRequest), http.StatusInternalServerError)
				return
			}

			output = response.Body
		}

		w.Header().Set("X-CSRF-Token", csrf.Token(r))

		buffer := new(bytes.Buffer)
		buffer.ReadFrom(output)
		fmt.Fprint(w, buffer.String())
	})
}

type graphQL interface {
	mutation(url string, body io.Reader) (*http.Response, error)
	query(url string) (*http.Response, error)
}

type GraphQLClient struct{}

func (gqlc *GraphQLClient) mutation(url string, body io.Reader) (*http.Response, error) {
	return http.Post(url, "application/json", body)
}

func (gqlc *GraphQLClient) query(url string) (*http.Response, error) {
	return http.Get(url)
}
