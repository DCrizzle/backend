package server

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mockGraphQL struct {
	postResponse *http.Response
	postError    error
}

func (mgql *mockGraphQL) post(url, token string, body io.Reader) (*http.Response, error) {
	return mgql.postResponse, mgql.postError
}

func Test_graphQLHandler(t *testing.T) {
	responseData := `{"data":{"key":"value"}}`

	tests := []struct {
		description  string
		requestURL   string
		requestBody  io.Reader // note: remove (?)
		postResponse *http.Response
		postError    error
		status       int
		body         string
	}{
		{
			description:  "error returned from database mutation request",
			requestURL:   "/graphql",
			requestBody:  nil,
			postResponse: nil,
			postError:    errors.New("mock mutation error"),
			status:       http.StatusInternalServerError,
			body:         fmt.Sprintf(`{"message":"%s"}`, errMutationRequest) + "\n",
		},
		{
			description: "successful mutation post request invocation",
			requestURL:  "/graphql",
			requestBody: nil,
			postResponse: &http.Response{
				Body: ioutil.NopCloser(strings.NewReader(responseData)),
			},
			postError: nil,
			status:    http.StatusOK,
			body:      responseData,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			mgql := &mockGraphQL{
				postResponse: test.postResponse,
				postError:    test.postError,
			}

			req, err := http.NewRequest(http.MethodPost, test.requestURL, test.requestBody)
			if err != nil {
				t.Fatal(err)
			}

			rec := httptest.NewRecorder()

			handler := http.HandlerFunc(graphQLHandler("testURL", mgql))

			handler.ServeHTTP(rec, req)

			if status := rec.Code; status != test.status {
				t.Errorf("status received: %d, expected: %d", status, test.status)
			}

			received := rec.Body.String()
			if received != test.body {
				t.Errorf("body received: %s, expected: %s", received, test.body)
			}
		})
	}
}
