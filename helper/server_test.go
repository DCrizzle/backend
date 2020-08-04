package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_getAuth0ManagementAPIToken(t *testing.T) {
	tests := []struct {
		description string
		response    string
		token       string
		err         string
	}{
		{
			description: "successful helper function invocation",
			response:    `{"access_token": "test-token"}`,
			token:       "test-token",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			mux := http.NewServeMux()
			mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, test.response)
			})

			server := httptest.NewServer(mux)
			token, err := getAuth0ManagementAPIToken(server.URL)

			if token != test.token {
				t.Errorf("token received: %s, expected: %s", token, test.token)
			}

			if err != nil && test.err != "" {
				t.Errorf("error received: %s, expected: %s", err.Error(), test.err)
			}
		})
	}
}

func Test_newServer(t *testing.T) {
	testHandler := func(w http.ResponseWriter, r *http.Request) {}
	server := newServer("TOKEN", testHandler)
	if server == nil {
		t.Error("error creating server")
	}
}

func Test_usersHandler(t *testing.T) {

	// outline:
	// [ ] create mock auth0 server/handlers
	// [ ] pass in mock client
	// - [ ] include url to mock auth0 server
	// [ ] test scenarios/data retrieved

}
