package api

import (
	// "bytes"
	"context"
	// "encoding/json"
	"fmt"
	// "io"
	"net/http"

	"github.com/machinebox/graphql"
)

const (
	errCreateOrg     = "error creating org"
	errCreateRequest = "error creating dgraph request"
	errCreateUser    = "error creating user"
	errDecodeJSON    = "error decoding json"
	errDeleteOrg     = "error deleting org"
	errEncodeJSON    = "error encoding json"
	errPOSTRequest   = "error sending post request"
	errReadBody      = "error reading response body"
	errRunClient     = "error running graphql client"
	errUpdateOrg     = "error updating org"
	errUpdateUser    = "error updating user"
)

const dgraphURL = "temporary"

// NOTE: These interfaces/structs will be used in the permanent implementation thats
// will eventually replace the quick-and-dirty version
type httpClienter interface {
	Do(req *http.Request) (*http.Response, error)
}

type dgrapher interface {
	mutate(c context.Context, m string, v map[string]interface{}, t interface{}) (interface{}, error)
	query(c context.Context, q string, v map[string]interface{}, t interface{}) (interface{}, error)
}

type dgraph struct {
	httpClient    httpClienter
	graphqlClient *graphql.Client // TEMP
}

type graphQLError struct {
	Message string
}

type graphQLResponse struct {
	Data   interface{} `json:"data"`
	Errors []graphQLError
}

func (d *dgraph) mutate(ctx context.Context, mutation string, variables map[string]interface{}, input interface{}) (interface{}, error) {
	// NOTE: This is a quick-and-dirty implementation that will likely be swapped out for
	// a hand-written version (very similar to this client library) but with allowance for
	// GET requests to be made
	// NOTE: DO NOT DELETE COMMENTED CODE
	request := graphql.NewRequest(mutation)
	for key, value := range variables {
		request.Var(key, value)
	}

	if err := d.graphqlClient.Run(ctx, request, &input); err != nil {
		return nil, fmt.Errorf("%s: %w", errRunClient, err)
	}

	return input, nil

	// NOTE: DO NOT DELETE
	// var bodyBuffer bytes.Buffer
	// body := struct {
	// 	Query     string                 `json:"query"`
	// 	Variables map[string]interface{} `json:"variables"`
	// }{
	// 	Query:     mutation,
	// 	Variables: variables,
	// }
	//
	// if err := json.NewEncoder(&bodyBuffer).Encode(body); err != nil {
	// 	return nil, fmt.Errorf("%s: %w", errEncodeJSON, err)
	// }
	//
	// request, err := http.NewRequest(http.MethodPost, dgraphURL, &bodyBuffer)
	// if err != nil {
	// 	return nil, fmt.Errorf("%s: %w", errCreateRequest, err)
	// }
	//
	// request.Header.Set("Content-Type", "application/json; charset=utf-8")
	// request.Header.Set("Accept", "application/json; charset=utf-8")
	//
	// response, err := d.httpClient.Do(request)
	// if err != nil || response.StatusCode != http.StatusOK {
	// 	return nil, fmt.Errorf("%s: %w", errPOSTRequest, err)
	// }
	// defer response.Body.Close()
	//
	// var buf bytes.Buffer
	// if _, err := io.Copy(&buf, response.Body); err != nil {
	// 	return nil, fmt.Errorf("%s: %w", errReadBody, err)
	// }
	//
	// output := &graphQLResponse{
	// 	Data: input,
	// }
	//
	// if err := json.NewDecoder(&buf).Decode(&output); err != nil {
	// 	return nil, fmt.Errorf("%s: %w", errDecodeJSON, err)
	// }
	//
	// return output.Data, nil
}

func (d *dgraph) query(ctx context.Context, query string, variables map[string]interface{}, input interface{}) (interface{}, error) {
	// NOTE: This is a quick-and-dirty implementation that will likely be swapped out for
	// a hand-written version (very similar to this client library) but with allowance for
	// GET requests to be made
	request := graphql.NewRequest(query)
	for key, value := range variables {
		request.Var(key, value)
	}

	if err := d.graphqlClient.Run(ctx, request, &input); err != nil {
		return nil, fmt.Errorf("%s: %w", errRunClient, err)
	}

	return input, nil
}

type Resolver struct {
	dgraph dgrapher
}

func NewResolver() *Resolver {
	return &Resolver{
		dgraph: &dgraph{
			httpClient:    &http.Client{},
			graphqlClient: graphql.NewClient(dgraphURL), // TEMP
		},
	}
}
