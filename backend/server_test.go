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
	responseData := `{"data":{"key":"value"}}`

	tests := []struct {
		description      string
		requestMethod    string
		requestURL       string
		requestBody      io.Reader
		mutationResponse io.ReadCloser
		queryResponse    io.ReadCloser
		status           int
		body             string
	}{
		{
			description:      "no json body received in post request",
			requestMethod:    http.MethodPost,
			requestURL:       "/graphql",
			requestBody:      nil,
			mutationResponse: nil,
			queryResponse:    nil,
			status:           http.StatusBadRequest,
			body:             fmt.Sprintf(`{"message":"%s"}`, errNoPOSTRequestBody) + "\n",
		},
		{
			description:      "non-json body received in post request",
			requestMethod:    http.MethodPost,
			requestURL:       "/graphql",
			requestBody:      strings.NewReader("---------"),
			mutationResponse: nil,
			queryResponse:    nil,
			status:           http.StatusBadRequest,
			body:             fmt.Sprintf(`{"message":"%s"}`, errParsingPOSTJSONBody) + "\n",
		},
		{
			description:      "successful mutation post request invocation",
			requestMethod:    http.MethodPost,
			requestURL:       "/graphql",
			requestBody:      strings.NewReader(`{"query":"test query","variables":{"key":"value"}}`),
			mutationResponse: ioutil.NopCloser(strings.NewReader(responseData)),
			queryResponse:    nil,
			status:           http.StatusOK,
			body:             responseData,
		},
		{
			description:      "no query parameter received in get request",
			requestMethod:    http.MethodGet,
			requestURL:       "/graphql",
			requestBody:      nil,
			mutationResponse: nil,
			queryResponse:    nil,
			status:           http.StatusBadRequest,
			body:             fmt.Sprintf(`{"message":"%s"}`, errNoQueryInURL) + "\n",
		},
		{
			description:      "incorrect variables parameter received in get request",
			requestMethod:    http.MethodGet,
			requestURL:       "/graphql?query=%20getData%20%7B%20value%20%7B&variables=---",
			requestBody:      nil,
			mutationResponse: nil,
			queryResponse:    nil,
			status:           http.StatusBadRequest,
			body:             fmt.Sprintf(`{"message":"%s"}`, errParsingGETQueryURL) + "\n",
		},
		{
			description:      "successful query get request invocation",
			requestMethod:    http.MethodGet,
			requestURL:       "/graphql?query=query%20getData(%24arg%3A%20String!)%20%7B%20getData(arg%3A%20%24arg)%20%7B%20value%20%7D%20%7D&variables=%7B%0A%20%20%22arg%22%3A%22value%22%0A%7D",
			requestBody:      nil,
			mutationResponse: nil,
			queryResponse:    ioutil.NopCloser(strings.NewReader(responseData)),
			status:           http.StatusOK,
			body:             responseData,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			mgql := &mockGraphQL{
				mutationResponse: test.mutationResponse,
				queryResponse:    test.queryResponse,
			}

			req, err := http.NewRequest(test.requestMethod, test.requestURL, test.requestBody)
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
