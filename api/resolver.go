package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	errCreateRequest = "error creating dgraph request"
	errDecodeJSON    = "error decoding json"
	errEncodeJSON    = "error encoding json"
	errPOSTRequest   = "error sending post request"
	errReadBody      = "error reading response body"

	// errCreateOrg     = "error creating org"
	// errCreateUser    = "error creating user"
	// errDeleteOrg     = "error deleting org"
	// errReadBytes     = "error reading body bytes"
)

const dgraphURL = "temporary"

type httpClienter interface {
	Do(req *http.Request) (*http.Response, error)
}

type dgrapher interface {
	mutate(m string, v map[string]interface{}, t interface{}) (interface{}, error)
	query(q string, v map[string]interface{}, t interface{}) (interface{}, error)
}

type dgraph struct {
	httpClient httpClienter
}

type graphQLError struct {
	Message string
}

type graphQLResponse struct {
	Data   interface{} `json:"data"`
	Errors []graphQLError
}

func (d *dgraph) mutate(mutation string, variables map[string]interface{}, input interface{}) (interface{}, error) {
	var bodyBuffer bytes.Buffer
	body := struct {
		Query     string                 `json:"query"`
		Variables map[string]interface{} `json:"variables"`
	}{
		Query:     mutation,
		Variables: variables,
	}

	if err := json.NewEncoder(&bodyBuffer).Encode(body); err != nil {
		return nil, fmt.Errorf("%s: %w", errEncodeJSON, err)
	}

	request, err := http.NewRequest(http.MethodPost, dgraphURL, &bodyBuffer)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errCreateRequest, err)
	}

	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Set("Accept", "application/json; charset=utf-8")

	response, err := d.httpClient.Do(request)
	if err != nil || response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: %w", errPOSTRequest, err)
	}
	defer response.Body.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, response.Body); err != nil {
		return nil, fmt.Errorf("%s: %w", errReadBody, err)
	}

	output := &graphQLResponse{
		Data: input,
	}

	if err := json.NewDecoder(&buf).Decode(&output); err != nil {
		return nil, fmt.Errorf("%s: %w", errDecodeJSON, err)
	}

	return output.Data, nil
}

func (d *dgraph) query(q string, v map[string]interface{}, t interface{}) (interface{}, error) {
	return nil, nil
}

type Resolver struct {
	dgraph dgrapher
}

func NewResolver() *Resolver {
	return &Resolver{
		dgraph: &dgraph{
			httpClient: &http.Client{},
		},
	}
}
