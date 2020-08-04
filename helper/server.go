package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type responseJSON struct {
	AccessToken string `json:"access_token"`
}

func getAuth0ManagementAPIToken(url string) (string, error) {
	c, err := readConfig()
	if err != nil {
		return "", err
	}

	payloadString := fmt.Sprintf(
		"grant_type=client_credentials&client_id=%s&client_secret=%s&audience=%s",
		c.Auth0.ClientID,
		c.Auth0.ClientSecret,
		c.Auth0.Audience,
	)

	payload := strings.NewReader(payloadString)
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	resJSON := responseJSON{}
	if err := json.Unmarshal(resBytes, &resJSON); err != nil {
		return "", err
	}

	return resJSON.AccessToken, nil
}

type server struct {
	auth0Token string
	httpServer *http.Server
}

func newServer(auth0Token string, usersHandler http.HandlerFunc) *server {
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
		auth0Token: auth0Token,
		httpServer: helperServer,
	}
}

func (s *server) start() {
	s.httpServer.ListenAndServe()
}

func (s *server) stop(ctx context.Context) {
	s.httpServer.Shutdown(ctx)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {

	// outline:
	// [ ] retrieve secret code
	// - [ ] validate against stored code
	// [ ] parse received data
	// [ ] construct auth0 request based on method
	// - [ ] post method
	// - [ ] patch method
	// - [ ] delete method

}
