package api

import (
	"io"
	"net/http"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

const (
	errCreateOrg     = "error creating org"
	errCreateUser    = "error creating user"
	errDeleteOrg     = "error deleting org"
	errReadBytes     = "error reading body bytes"
	errUnmarshalJSON = "error unmarshalling json"
)

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
