package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	errPOST   = "error invoking post request"
	errGET    = "error invoking get request"
	errPOSTDB = "error invoking graphql database"
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
			http.Error(w, errPOSTDB, http.StatusInternalServerError)
			return
		}

		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)

		w.Write(buf.Bytes())
	}
}

func (h *help) query() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// outline:
		// [ ] parse get request query string
		// [ ] create/clean query operation string
		// [ ] set headers/values in http get request
		// [ ] invoke get request
		// [ ] parse received response
		// [ ] assert/format type required
		// [ ] create/send api response

	}
}

type server struct {
	httpServer http.Server
}

func newServer(addr string) (*server, error) {
	router := mux.NewRouter()

	// // outline:
	// // [ ] add non-auth handlers here

	subrouter := router.Host(addr).Subrouter()

	h := new(help)

	subrouter.Use(h.secure)

	subrouter.HandleFunc("/org/{orgID}/graphql", h.mutate()).Methods("POST")

	subrouter.HandleFunc("/org/{orgID}/graphql", h.query()).Methods("GET")

	return &server{
		httpServer: http.Server{
			Addr:    addr,
			Handler: router,
		},
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
// - [x] "/org/{id}/db" -> execute mutations (POST)
// - [x] "/org/{id}/db" -> execute queries (GET)
// - [ ] "/org/{id}/admin" -> org settings (GET)
// - [ ] "/org/{id}/admin/users" -> users settings (GET)
// - [ ] "/org/{id}/admin/user/{id}" -> user settings (GET)
// [ ] new server method
// [ ] start server method
// [ ] stop server method
