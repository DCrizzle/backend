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

func Test_graphQLHandler(t *testing.T) {
	responseData := `{"data":{"key":"value"}}`

	tests := []struct {
		description      string
		requestMethod    string
		requestURL       string
		requestBody      io.Reader // note: remove (?)
		mutationResponse *http.Response
		mutationError    error
		queryResponse    *http.Response
		queryError       error
		status           int
		body             string
	}{
		{
			description:      "error returned from database mutation request",
			requestMethod:    http.MethodPost,
			requestURL:       "/graphql",
			requestBody:      nil,
			mutationResponse: nil,
			mutationError:    errors.New("mock mutation error"),
			queryResponse:    nil,
			queryError:       nil,
			status:           http.StatusInternalServerError,
			body:             fmt.Sprintf(`{"message":"%s"}`, errMutationRequest) + "\n",
		},
		{
			description:   "successful mutation post request invocation",
			requestMethod: http.MethodPost,
			requestURL:    "/graphql",
			requestBody:   nil,
			mutationResponse: &http.Response{
				Body: ioutil.NopCloser(strings.NewReader(responseData)),
			},
			mutationError: nil,
			queryResponse: nil,
			queryError:    nil,
			status:        http.StatusOK,
			body:          responseData,
		},
		{
			description:      "error returned from database query request",
			requestMethod:    http.MethodGet,
			requestURL:       "/graphql",
			requestBody:      nil,
			mutationResponse: nil,
			mutationError:    nil,
			queryResponse:    nil,
			queryError:       errors.New("mock query error"),
			status:           http.StatusInternalServerError,
			body:             fmt.Sprintf(`{"message":"%s"}`, errQueryRequest) + "\n",
		},
		{
			description:      "successful query get request invocation",
			requestMethod:    http.MethodGet,
			requestURL:       "/graphql?query=query%20getData(%24arg%3A%20String!)%20%7B%20getData(arg%3A%20%24arg)%20%7B%20value%20%7D%20%7D&variables=%7B%0A%20%20%22arg%22%3A%22value%22%0A%7D",
			requestBody:      nil,
			mutationResponse: nil,
			mutationError:    nil,
			queryResponse: &http.Response{
				Body: ioutil.NopCloser(strings.NewReader(responseData)),
			},
			queryError: nil,
			status:     http.StatusOK,
			body:       responseData,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			mgql := &mockGraphQL{
				mutationResponse: test.mutationResponse,
				mutationError:    test.mutationError,
				queryResponse:    test.queryResponse,
				queryError:       test.queryError,
			}

			req, err := http.NewRequest(test.requestMethod, test.requestURL, test.requestBody)
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