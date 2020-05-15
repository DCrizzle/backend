package server

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	auth0 "github.com/auth0-community/go-auth0"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	jose "gopkg.in/square/go-jose.v2"
)

const (
	errMutationRequest = "server: error invoking backend database mutation"
	errQueryRequest    = "server: error invoking backend database query"
)

var (
	errValidateAuth0APIAudience     = errors.New("server: param auth0 api audience not set")
	errValidateAuth0APIClientSecret = errors.New("server: param auth0 api client secret not set")
	errValidateAuth0Domain          = errors.New("server: param auth0 domain not set")
	errValidateCSRFKey              = errors.New("server: param csrf key not set")
	errValidateDgraphURL            = errors.New("server: param dgraph url not set")

	errParseReadFile  = errors.New("server: error reading config file")
	errParseUnmarshal = errors.New("server: error unmarshalling config file content")
)

type params struct {
	Auth0APIAudience     string `json:"auth0_api_audience"`
	Auth0APIClientSecret string `json:"auth0_api_client_secret"`
	Auth0Domain          string `json:"auth0_domain"`
	CSRFKey              string `json:"csrf_key"`
	DgraphURL            string `json:"dgraph_url"`
}

func parseParams(configPath string) (*params, error) {
	configContent, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, errParseReadFile
	}

	p := &params{}
	if err := json.Unmarshal(configContent, p); err != nil {
		return nil, errParseUnmarshal
	}

	return p, nil
}

func (p *params) validate() error {
	if p.Auth0APIAudience == "" {
		return errValidateAuth0APIAudience
	}
	if p.Auth0APIClientSecret == "" {
		return errValidateAuth0APIClientSecret
	}
	if p.Auth0Domain == "" {
		return errValidateAuth0Domain
	}
	if p.CSRFKey == "" {
		return errValidateCSRFKey
	}
	if p.DgraphURL == "" {
		return errValidateDgraphURL
	}

	return nil
}

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

type GraphQLClient struct{}

func (gqlc *GraphQLClient) mutation(url string, body io.Reader) (*http.Response, error) {
	return http.Post(url, "application/json", body)
}

func (gqlc *GraphQLClient) query(url string) (*http.Response, error) {
	return http.Get(url)
}
