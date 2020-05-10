package backend

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	auth0 "github.com/auth0-community/go-auth0"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	jose "gopkg.in/square/go-jose.v2"
)

const (
	errMutationRequest = "error invoking backend database mutation"
	errQueryRequest    = "error invoking backend database query"
)

// Server hosts the backend server with login/logout and GraphQL endpoints.
type Server struct {
	httpServer *http.Server
}

// NewServer generates a pointer to an inactive Server instance.
func NewServer(dgraphURL string, gql graphQL) *Server {
	router := mux.NewRouter()

	router.Use(middleware)
	router.HandleFunc("/graphql", graphQLHandler(dgraphURL, gql))

	csrf.Protect([]byte("TEMP_32_BYTE_LONG_ARRAY"))(router)

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

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secret := []byte("AUTH0_API_CLIENT_SECRET")
		secretProvider := auth0.NewKeyProvider(secret)
		audience := []string{"AUTH0_API_AUDIENCE"}
		domain := "https://AUTH0_DOMAIN.auth0.com/"

		configuration := auth0.NewConfiguration(secretProvider, audience, domain, jose.RS256)
		validator := auth0.NewValidator(configuration, nil)

		_, err := validator.ValidateRequest(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("unauthorized"))
		}

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

type graphQLClient struct{}

func (gqlc *graphQLClient) mutation(url string, body io.Reader) (*http.Response, error) {
	return http.Post(url, "application/json", body)
}

func (gqlc *graphQLClient) query(url string) (*http.Response, error) {
	return http.Get(url)
}
