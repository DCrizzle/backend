package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

const (
	errPOST     = "error invoking post request"
	errGET      = "error invoking get request"
	errMutateDB = "error invoking graphql database mutation"
	errQueryDB  = "error invoking graphql database query"
)

type httpHelper interface {
	post(string, string, io.Reader) (*http.Response, error)
	get(string) (*http.Response, error)
}

type httpHelp struct{}

func (h *httpHelp) post(url, contentType string, payload io.Reader) (*http.Response, error) {

	// outline:
	// [ ] set headers/values in http post request

	resp, err := http.Post(url, contentType, payload)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errPOST, err)
	}
	return resp, nil
}

func (h *httpHelp) get(url string) (*http.Response, error) {

	// outline:
	// [ ] set headers/values in http post request

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errGET, err)
	}
	return resp, nil
}

type help struct {
	client httpHelper
	addr   string
}

func (h *help) secure(hdlr http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		_ = token

		// outline:
		// [ ] query database for user
		// [ ] return success/error based on response

		hdlr.ServeHTTP(w, r)
	})
}

func (h *help) mutate() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		dbURL := h.addr + "/graphql"
		contentType := "application/json"
		payload := r.Body

		resp, err := h.client.post(dbURL, contentType, payload)
		if err != nil {
			http.Error(w, fmt.Errorf("%s: %w", errMutateDB, err).Error(), http.StatusInternalServerError)
			return
		}

		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)

		w.Write(buf.Bytes())
	}
}

func (h *help) query() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()["query"][0]
		dbURL := h.addr + "/graphql?query=" + url.QueryEscape(query)

		resp, err := h.client.get(dbURL)
		if err != nil {
			http.Error(w, fmt.Errorf("%s: %w", errQueryDB, err).Error(), http.StatusInternalServerError)
			return
		}

		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)

		w.Write(buf.Bytes())
	}
}

// Server hosts the database API endpoints and application logic
type Server struct {
	httpServer http.Server
}

// NewServer generates a new Server instance
func NewServer(addr string) (*Server, error) {
	router := mux.NewRouter()

	// // outline:
	// // [ ] add non-auth handlers here

	subrouter := router.Host(addr).Subrouter()

	h := &help{
		client: new(httpHelp),
		addr:   "http://" + addr + ":8080", // NOTE: change to environment variable (?)
	}

	subrouter.Use(h.secure)

	subrouter.HandleFunc("/org/{orgID}/graphql", h.mutate()).Methods("POST")

	subrouter.HandleFunc("/org/{orgID}/graphql", h.query()).Methods("GET")

	return &Server{
		httpServer: http.Server{
			Addr:    addr + ":8181",
			Handler: router,
		},
	}, nil
}

// Start serves the HTTP listener
func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

// Stop shutsdown the HTTP listener
func (s *Server) Stop() error {
	ctx := context.Background()
	return s.httpServer.Shutdown(ctx)
}

// outline:
// [ ] handlers
// - [ ] "/" -> redirect to "/login"
// - [ ] "/login" -> validate user login; redirect to "/org/{id}" (POST)
// - [ ] "/org/{id}" -> render org (GET)
// - [x] "/org/{id}/db" -> execute mutations (POST)
// - [x] "/org/{id}/db" -> execute queries (GET)
// - [ ] "/org/{id}/admin" -> org settings (GET)
// - [ ] "/org/{id}/admin/users" -> users settings (GET)
// - [ ] "/org/{id}/admin/user/{id}" -> user settings (GET)
// [x] new server method
// [x] start server method
// [x] stop server method
// [ ] NOTE: possibly implement graphql-go to control possible graphql operations
