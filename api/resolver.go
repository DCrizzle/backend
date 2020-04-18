package api

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

const (
	errAddUser     = "error adding user"
	errCreateItem  = "error creating item"
	errCreateItems = "error creating items"
	errCreateOrg   = "error creating org"
	errCreateUser  = "error creating user"
	errDeleteOrg   = "error deleting org"
	errDeleteUser  = "error deleting user"
	errFilterItems = "error filtering items"
	errReadItems   = "error reading items"
	errReadOrg     = "error reading org"
	errReadUser    = "error reading user"
	errReadUsers   = "error reading users"
	errRemoveUser  = "error removing user"
	errRunClient   = "error running graphql client"
	errUpdateOrg   = "error updating org"
	errUpdateUser  = "error updating user"
)

const dgraphURL = "temporary"

func newRequest(mutation string, variables map[string]interface{}) *graphql.Request {
	request := graphql.NewRequest(mutation)
	for key, value := range variables {
		request.Var(key, value)
	}

	return request
}

type graphQLClienter interface {
	Run(ctx context.Context, req *graphql.Request, resp interface{}) error
}

type dgrapher interface {
	mutate(c context.Context, m string, v map[string]interface{}, t interface{}) (interface{}, error)
	query(c context.Context, q string, v map[string]interface{}, t interface{}) (interface{}, error)
}

type dgraph struct {
	graphQLClient graphQLClienter
}

func (d *dgraph) mutate(ctx context.Context, mutation string, variables map[string]interface{}, input interface{}) (interface{}, error) {
	request := newRequest(mutation, variables)
	if err := d.graphQLClient.Run(ctx, request, &input); err != nil {
		return nil, fmt.Errorf("%s: %w", errRunClient, err)
	}

	return input, nil
}

func (d *dgraph) query(ctx context.Context, query string, variables map[string]interface{}, input interface{}) (interface{}, error) {
	request := newRequest(query, variables)
	if err := d.graphQLClient.Run(ctx, request, &input); err != nil {
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
			graphQLClient: graphql.NewClient(dgraphURL), // TEMP
		},
	}
}
