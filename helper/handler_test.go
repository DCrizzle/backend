package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_usersHandler(t *testing.T) {
	tests := []struct {
		description        string
		mockAPIMethod      string
		mockAPIStatusCode  int
		helperSecret       string
		requestSecret      string
		requestBody        []byte
		responseStatusCode int
		responseBody       string
	}{
		{
			description:        "incorrect secret provided in request",
			mockAPIMethod:      http.MethodPost,
			mockAPIStatusCode:  http.StatusTeapot,
			helperSecret:       "correct_secret",
			requestSecret:      "incorrect_secret",
			requestBody:        []byte{},
			responseStatusCode: http.StatusBadRequest,
			responseBody:       errIncorrectSecret,
		},
	}

	// scenarios:
	// [x] incorrect secret
	// [ ] invalid request json
	// [ ] error response from mock server
	// [ ] successful create
	// [ ] successful update
	// [ ] successful delete

	for _, test := range tests {
		var apiBodyReceived io.ReadCloser

		mux := http.NewServeMux()
		mux.HandleFunc(test.mockAPIMethod, func(w http.ResponseWriter, r *http.Request) {
			apiBodyReceived = r.Body
			w.WriteHeader(test.mockAPIStatusCode)
			w.Write([]byte(`{"user_id": "auth0|123456"}`))
		})

		server := httptest.NewServer(mux)
		url := server.URL

		t.Run(test.description, func(t *testing.T) {
			req, err := http.NewRequest(test.mockAPIMethod, "/auth0/users", bytes.NewReader(test.requestBody))
			if err != nil {
				t.Fatal("error creating request:", err.Error())
			}
			req.Header.Set("folivora-helper-secret", test.requestSecret)

			rec := httptest.NewRecorder()

			handler := http.HandlerFunc(usersHandler(test.helperSecret, "test_token", url))

			handler.ServeHTTP(rec, req)

			if status := rec.Code; status != test.responseStatusCode {
				t.Errorf("status received: %d, expected: %d", status, test.responseStatusCode)
			}

			if body := rec.Body.String(); strings.TrimSuffix(body, "\n") != test.responseBody {
				t.Errorf("body received: %s, expected: %s", body, test.responseBody)
			}

			if apiBodyReceived != nil {
				receivedBytes, err := ioutil.ReadAll(apiBodyReceived)
				if err != nil {
					t.Fatal("error reading api body received:", err.Error())
				}

				receivedString := string(receivedBytes)
				requestString := string(test.requestBody)
				if receivedString != requestString {
					t.Errorf("api body received: %s, expected: %s", receivedString, requestString)
				}
			}
		})
	}
}
