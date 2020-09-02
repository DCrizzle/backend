package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_usersHandler(t *testing.T) {
	auth0ID := "auth0|id"

	tests := []struct {
		description         string
		auth0RespStatusCode int
		auth0RespID         string
		auth0ReqReceived    string
		helperSecret        string
		requestSecret       string
		requestMethod       string
		requestBody         []byte
		responseStatusCode  int
		responseBody        string
	}{
		{
			description:         "incorrect secret provided in request to helper",
			auth0RespStatusCode: http.StatusTeapot,
			auth0RespID:         "",
			auth0ReqReceived:    "",
			helperSecret:        "correct_secret",
			requestSecret:       "incorrect_secret",
			requestMethod:       http.MethodPost,
			requestBody:         []byte{},
			responseStatusCode:  http.StatusBadRequest,
			responseBody:        errIncorrectSecret,
		},
		{
			description:         "invalid json body received in request to helper",
			auth0RespStatusCode: http.StatusTeapot,
			auth0RespID:         "",
			auth0ReqReceived:    "",
			helperSecret:        "correct_secret",
			requestSecret:       "correct_secret",
			requestMethod:       http.MethodPost,
			requestBody:         []byte("---------"),
			responseStatusCode:  http.StatusBadRequest,
			responseBody:        errIncorrectRequestBody,
		},
		{
			description:         "unsupported http method in request to helper",
			auth0RespStatusCode: http.StatusTeapot,
			auth0RespID:         "",
			auth0ReqReceived:    "",
			helperSecret:        "correct_secret",
			requestSecret:       "correct_secret",
			requestMethod:       http.MethodPut,
			requestBody:         []byte(`{"email": "grandmaster@jeditemple.edu"}`),
			responseStatusCode:  http.StatusBadRequest,
			responseBody:        errIncorrectHTTPMethod,
		},
		{
			description:         "error received in response from auth0 server",
			auth0RespStatusCode: http.StatusBadRequest,
			auth0RespID:         "",
			auth0ReqReceived:    `{"email":"masteroftheorder@jeditemple.edu","password":"may-the-force-be-with-you","app_metadata":{"role":"USER_ADMIN","orgID":"jedi"},"given_name":"mace","family_name":"windu","connection":"Username-Password-Authentication"}`,
			helperSecret:        "correct_secret",
			requestSecret:       "correct_secret",
			requestMethod:       http.MethodPost,
			requestBody:         []byte(`{"owner":"jedi","email":"masteroftheorder@jeditemple.edu","password":"may-the-force-be-with-you","firstName":"mace","lastName":"windu","role":"USER_ADMIN","org":"jedi"}`),
			responseStatusCode:  http.StatusInternalServerError,
			responseBody:        errExecutingAuth0Request,
		},
		{
			description:         "successful create user request to helper server",
			auth0RespStatusCode: http.StatusOK,
			auth0RespID:         "",
			auth0ReqReceived:    `{"email":"battlemaster@jeditemple.edu","password":"may-the-force-be-with-you","app_metadata":{"role":"USER_ADMIN","orgID":"jedi"},"given_name":"cin","family_name":"dralling","connection":"Username-Password-Authentication"}`,
			helperSecret:        "correct_secret",
			requestSecret:       "correct_secret",
			requestMethod:       http.MethodPost,
			requestBody:         []byte(`{"owner":"jedi","email":"battlemaster@jeditemple.edu","password":"may-the-force-be-with-you","firstName":"cin","lastName":"dralling","role":"USER_ADMIN","org":"jedi"}`),
			responseStatusCode:  http.StatusOK,
			responseBody:        fmt.Sprintf(`{"message": "success", "auth0ID": "%s"}`, auth0ID),
		},
		{
			description:         "successful update user request to helper server",
			auth0RespStatusCode: http.StatusOK,
			auth0RespID:         "/",
			auth0ReqReceived:    `{"app_metadata":{"role":"USER_LAB"}}`,
			helperSecret:        "correct_secret",
			requestSecret:       "correct_secret",
			requestMethod:       http.MethodPatch,
			requestBody:         []byte(fmt.Sprintf(`{"authZeroID":"%s","role":"USER_LAB"}`, auth0ID)),
			responseStatusCode:  http.StatusOK,
			responseBody:        fmt.Sprintf(`{"message": "success", "auth0ID": "%s"}`, auth0ID),
		},
		{
			description:         "successful delete user request to helper server",
			auth0RespStatusCode: http.StatusOK,
			auth0RespID:         "/",
			auth0ReqReceived:    "",
			helperSecret:        "correct_secret",
			requestSecret:       "correct_secret",
			requestMethod:       http.MethodDelete,
			requestBody:         []byte(fmt.Sprintf(`{"authZeroID":"%s"}`, auth0ID)),
			responseStatusCode:  http.StatusOK,
			responseBody:        fmt.Sprintf(`{"message": "success", "auth0ID": "%s"}`, auth0ID),
		},
	}

	for _, test := range tests {
		var apiBodyReceived []byte

		handlerURL := "/users" + test.auth0RespID
		auth0Mux := http.NewServeMux()
		auth0Mux.HandleFunc(handlerURL, func(w http.ResponseWriter, r *http.Request) {
			receivedBytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "error parsing received body", http.StatusBadRequest)
				return
			}

			apiBodyReceived = receivedBytes
			w.WriteHeader(test.auth0RespStatusCode)
			w.Write([]byte(fmt.Sprintf(`{"user_id": "%s"}`, auth0ID)))
		})

		auth0Server := httptest.NewServer(auth0Mux)
		auth0URL := auth0Server.URL + "/"

		dgraphMux := http.NewServeMux()
		dgraphMux.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		dgraphServer := httptest.NewServer(dgraphMux)
		dgraphURL := dgraphServer.URL + "/graphql"

		t.Run(test.description, func(t *testing.T) {
			// mock request from dgraph @custom directive to helper server/handler
			req, err := http.NewRequest(
				test.requestMethod,
				"/auth0/users",
				bytes.NewReader(test.requestBody),
			)
			if err != nil {
				t.Fatal("error creating request:", err.Error())
			}
			req.Header.Set("folivora-helper-secret", test.requestSecret)

			rec := httptest.NewRecorder()

			// helper users handler wrapper function
			handler := http.HandlerFunc(usersHandler(test.helperSecret, "test_token", auth0URL, dgraphURL))

			handler.ServeHTTP(rec, req)

			if status := rec.Code; status != test.responseStatusCode {
				t.Errorf("status received: %d, expected: %d", status, test.responseStatusCode)
			}

			if body := strings.TrimSuffix(rec.Body.String(), "\n"); body != test.responseBody {
				t.Errorf("body received: %s, expected: %s", body, test.responseBody)
			}

			if apiBodyReceived != nil {
				receivedString := string(apiBodyReceived)
				if test.auth0ReqReceived != receivedString {
					t.Errorf("api body received: %s, expected: %s", receivedString, test.auth0ReqReceived)
				}
			}
		})
	}
}