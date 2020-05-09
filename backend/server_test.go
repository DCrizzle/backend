package backend

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mockGraphQL struct {
	mutationResponse io.ReadCloser
	queryResponse    io.ReadCloser
}

func (mgql *mockGraphQL) mutation(mutation string, variables map[string]string, headers map[string][]string) io.ReadCloser {
	return mgql.mutationResponse
}

func (mgql *mockGraphQL) query(query string, variables map[string]string, headers map[string][]string) io.ReadCloser {
	return mgql.queryResponse
}

func TestNewServer(t *testing.T) {
	mgql := &mockGraphQL{}

	s := NewServer(mgql)

	if s == nil {
		t.Error("nil received creating new server")
	}
}

func Test_graphQLHandler(t *testing.T) {
	tests := []struct {
		description      string
		requestMethod    string
		requestBody      io.Reader
		mutationResponse io.ReadCloser
		queryResponse    io.ReadCloser
		status           int
		body             string
	}{
		{
			description:      "no json body received in post request",
			requestMethod:    http.MethodPost,
			requestBody:      nil,
			mutationResponse: nil,
			queryResponse:    nil,
			status:           http.StatusBadRequest,
			body:             fmt.Sprintf(`{"message":"%s"}`, errNoPOSTRequestBody) + "\n",
		},
		{
			description:      "non-json body received in post request",
			requestMethod:    http.MethodPost,
			requestBody:      strings.NewReader("---------"),
			mutationResponse: nil,
			queryResponse:    nil,
			status:           http.StatusBadRequest,
			body:             fmt.Sprintf(`{"message":"%s"}`, errParsingPOSTJSONBody) + "\n",
		},
		{
			description:      "successful mutation post request invocation",
			requestMethod:    http.MethodPost,
			requestBody:      strings.NewReader(`{"query":"test query","variables":{"key":"value"}}`),
			mutationResponse: ioutil.NopCloser(strings.NewReader(`{"data":{"key":"value"}}`)),
			queryResponse:    nil,
			status:           http.StatusOK,
			body:             `{"data":{"key":"value"}}`,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			mgql := &mockGraphQL{
				mutationResponse: test.mutationResponse,
				queryResponse:    test.queryResponse,
			}

			req, err := http.NewRequest(test.requestMethod, "/graphql", test.requestBody)
			if err != nil {
				t.Fatal(err)
			}

			rec := httptest.NewRecorder()

			handler := http.HandlerFunc(graphQLHandler(mgql))

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

func Test_middleware(t *testing.T) {

}
