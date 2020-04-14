package api

import (
	"io"
	"net/http"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type httpClient interface {
	Post(url, contentType string, body io.Reader) (*http.Response, error)
	Get(url string) (*http.Response, error)
}

type Resolver struct {
	httpClient httpClient
	dbURL      string
}

func NewResolver() *Resolver {
	client := &http.Client{}

	return &Resolver{
		httpClient: client,
		dbURL:      "temporary",
	}
}
