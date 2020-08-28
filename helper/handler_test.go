package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_usersHandler(t *testing.T) {
	tests := []struct {
		description        string
		mockAPIStatusCode  int
		mockAPIPath        string
		mockAPIAuth0ID     string
		mockAPIRequest     string
		helperSecret       string
		requestSecret      string
		requestMethod      string
		requestBody        []byte
		responseStatusCode int
		responseBody       string
	}{
		{
			description:        "incorrect secret provided in request to helper",
			mockAPIStatusCode:  http.StatusTeapot,
			mockAPIPath:        "/auth0/users",
			mockAPIAuth0ID:     "",
			mockAPIRequest:     "",
			helperSecret:       "correct_secret",
			requestSecret:      "incorrect_secret",
			requestMethod:      http.MethodPost,
			requestBody:        []byte{},
			responseStatusCode: http.StatusBadRequest,
			responseBody:       errIncorrectSecret,
		},
		{
			description:        "invalid json body received in request to helper",
			mockAPIStatusCode:  http.StatusTeapot,
			mockAPIPath:        "/auth0/users",
			mockAPIAuth0ID:     "",
			mockAPIRequest:     "",
			helperSecret:       "correct_secret",
			requestSecret:      "correct_secret",
			requestMethod:      http.MethodPost,
			requestBody:        []byte("---------"),
			responseStatusCode: http.StatusBadRequest,
			responseBody:       errIncorrectRequestBody,
		},
		{
			description:        "unsupported http method in request to helper",
			mockAPIStatusCode:  http.StatusTeapot,
			mockAPIPath:        "/auth0/users",
			mockAPIAuth0ID:     "",
			mockAPIRequest:     "",
			helperSecret:       "correct_secret",
			requestSecret:      "correct_secret",
			requestMethod:      http.MethodPut,
			requestBody:        []byte(`{"email": "grandmaster@jeditemple.edu"}`),
			responseStatusCode: http.StatusBadRequest,
			responseBody:       errIncorrectHTTPMethod,
		},
		{
			description:        "error received in response from auth0 server",
			mockAPIStatusCode:  http.StatusBadRequest,
			mockAPIPath:        "/auth0/users",
			mockAPIAuth0ID:     "",
			mockAPIRequest:     `{"email":"masteroftheorder@jeditemple.edu","app_metadata":{"role":"","orgID":""},"given_name":"","family_name":"","connection":"Username-Password-Authentication"}`,
			helperSecret:       "correct_secret",
			requestSecret:      "correct_secret",
			requestMethod:      http.MethodPost,
			requestBody:        []byte(`{"email":"masteroftheorder@jeditemple.edu","app_metadata":{"role":"","orgID":""},"given_name":"","family_name":"","connection":"Username-Password-Authentication"}`),
			responseStatusCode: http.StatusInternalServerError,
			responseBody:       errExecutingAuth0Request,
		},
		{
			description:        "successful create user request to helper server",
			mockAPIStatusCode:  http.StatusOK,
			mockAPIPath:        "/auth0/users",
			mockAPIAuth0ID:     "",
			mockAPIRequest:     `{"email":"battlemaster@jeditemple.edu","app_metadata":{"role":"","orgID":""},"given_name":"","family_name":"","connection":"Username-Password-Authentication"}`,
			helperSecret:       "correct_secret",
			requestSecret:      "correct_secret",
			requestMethod:      http.MethodPost,
			requestBody:        []byte(`{"email":"battlemaster@jeditemple.edu","app_metadata":{"role":"USER_ADMIN","orgID":"jedi-order"},"given_name":"","family_name":"","connection":"Username-Password-Authentication"}`),
			responseStatusCode: http.StatusOK,
			responseBody:       `{"message": "success", "auth0ID": "auth0|id"}`,
		},
		{
			description:        "successful update user request to helper server",
			mockAPIStatusCode:  http.StatusOK,
			mockAPIPath:        "/auth0/users",
			mockAPIAuth0ID:     "/auth0|id",
			mockAPIRequest:     `{"app_metadata":{"role":"USER_ADMIN","orgID":""}}`,
			helperSecret:       "correct_secret",
			requestSecret:      "correct_secret",
			requestMethod:      http.MethodPatch,
			requestBody:        []byte(`{"auth0ID":"auth0|id","role":"USER_ADMIN","orgID":""}`),
			responseStatusCode: http.StatusOK,
			responseBody:       `{"message": "success", "auth0ID": "auth0|id"}`,
		},
		{
			description:        "successful delete user request to helper server",
			mockAPIStatusCode:  http.StatusOK,
			mockAPIPath:        "/auth0/users",
			mockAPIAuth0ID:     "/auth0|id",
			mockAPIRequest:     "",
			helperSecret:       "correct_secret",
			requestSecret:      "correct_secret",
			requestMethod:      http.MethodDelete,
			requestBody:        []byte(`{"auth0ID":"auth0|id"}`),
			responseStatusCode: http.StatusOK,
			responseBody:       `{"message": "success", "auth0ID": "auth0|id"}`,
		},
	}

	for _, test := range tests {
		var apiBodyReceived []byte

		mux := http.NewServeMux()
		mux.HandleFunc(test.mockAPIPath+test.mockAPIAuth0ID, func(w http.ResponseWriter, r *http.Request) {
			receivedBytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "error parsing received body", http.StatusBadRequest)
				return
			}

			apiBodyReceived = receivedBytes
			w.WriteHeader(test.mockAPIStatusCode)
			w.Write([]byte(`{"auth0ID": "auth0|id"}`))
		})

		server := httptest.NewServer(mux)
		url := server.URL + test.mockAPIPath

		t.Run(test.description, func(t *testing.T) {
			// mock request from dgraph to helper server/handler
			req, err := http.NewRequest(
				test.requestMethod,
				test.mockAPIPath,
				bytes.NewReader(test.requestBody),
			)
			if err != nil {
				t.Fatal("error creating request:", err.Error())
			}
			req.Header.Set("folivora-helper-secret", test.requestSecret)

			rec := httptest.NewRecorder()

			// helper users handler wrapper function
			handler := http.HandlerFunc(usersHandler(test.helperSecret, "test_token", url))

			handler.ServeHTTP(rec, req)

			if status := rec.Code; status != test.responseStatusCode {
				t.Errorf("status received: %d, expected: %d", status, test.responseStatusCode)
			}

			if body := strings.TrimSuffix(rec.Body.String(), "\n"); body != test.responseBody {
				t.Errorf("body received: %s, expected: %s", body, test.responseBody)
			}

			if apiBodyReceived != nil {
				receivedString := string(apiBodyReceived)
				if test.mockAPIRequest != receivedString {
					t.Errorf("api body received: %s, expected: %s", receivedString, test.mockAPIRequest)
				}
			}
		})
	}
}
