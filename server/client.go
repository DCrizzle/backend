package server

import (
	"io"
	"net/http"
)

type graphQL interface {
	post(url, token string, body io.Reader) (*http.Response, error)
}

type graphQLClient struct {
	httpClient *http.Client
}

func newGraphQLClient(httpClient *http.Client) *graphQLClient {
	return &graphQLClient{
		httpClient: httpClient,
	}
}

func (gqlc *graphQLClient) post(url, token string, body io.Reader) (*http.Response, error) {
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("X-Auth0-Token", token)

	response, err := gqlc.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
