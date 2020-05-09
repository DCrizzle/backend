package backend

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	errNoPOSTRequestBody   = "no body received in payload"
	errParsingPOSTJSONBody = "error parsing json body"
	errNoQueryInURL        = "no query variable recieved in url params"
	errParsingGETQueryURL  = "error parsing query url variable"
)

// Server hosts the backend server with login/logout and GraphQL endpoints.
type Server struct {
	httpServer *http.Server
}

// NewServer generates a pointer to an inactive Server instance.
func NewServer(gql graphql) *Server {
	router := mux.NewRouter()

	router.Use(middleware)
	router.HandleFunc("/graphql", graphQLHandler(gql))

	return &Server{
		httpServer: &http.Server{
			Addr:         ":8888",
			Handler:      router,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
}

// Start begins serving the http.Server.
func (s *Server) Start() {
	s.httpServer.ListenAndServe()
}

// Stop ends serving the http.Server.
func (s *Server) Stop(ctx context.Context) {
	s.httpServer.Shutdown(ctx)
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

func graphQLHandler(gql graphql) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var output io.ReadCloser

		if r.Method == http.MethodPost {
			mut := mutation{}
			if r.Body == nil {
				http.Error(w, fmt.Sprintf(`{"message":"%s"}`, errNoPOSTRequestBody), http.StatusBadRequest)
				return
			}

			err := json.NewDecoder(r.Body).Decode(&mut)
			if err != nil {
				http.Error(w, fmt.Sprintf(`{"message":"%s"}`, errParsingPOSTJSONBody), http.StatusBadRequest)
				return
			}

			mutationResponse := gql.mutation(mut.mutation, mut.variables, r.Header)
			output = mutationResponse

		} else if r.Method == http.MethodGet {
			params := r.URL.Query()
			queryParam, ok := params["query"]
			if !ok {
				http.Error(w, fmt.Sprintf(`{"message":"%s"}`, errNoQueryInURL), http.StatusBadRequest)
				return
			}

			var variablesParam map[string]string
			variablesString, ok := params["variables"]
			if ok {
				err := json.Unmarshal([]byte(variablesString[0]), &variablesParam)
				if err != nil {
					http.Error(w, fmt.Sprintf(`{"message":"%s"}`, errParsingGETQueryURL), http.StatusBadRequest)
					return
				}
			}

			queryResponse := gql.query(queryParam[0], variablesParam, r.Header)
			output = queryResponse

		}

		buffer := new(bytes.Buffer)
		buffer.ReadFrom(output)
		fmt.Fprint(w, buffer.String())
	})
}

type mutation struct {
	mutation  string            `json:"query"`
	variables map[string]string `json:"variables"`
}

type graphql interface {
	mutation(mutation string, variables map[string]string, headers map[string][]string) io.ReadCloser
	query(query string, variables map[string]string, headers map[string][]string) io.ReadCloser
}

// todo: add implementation of graphql interface utilizing http client for graphql
