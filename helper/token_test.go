package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_getAuth0APIToken(t *testing.T) {
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
			c := &config{}
			token, err := getAuth0APIToken(server.URL, c.Auth0)

			if token != test.token {
				t.Errorf("token received: %s, expected: %s", token, test.token)
			}

			if err != nil && test.err != "" {
				t.Errorf("error received: %s, expected: %s", err.Error(), test.err)
			}
		})
	}
}
