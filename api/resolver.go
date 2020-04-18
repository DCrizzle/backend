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

type dgrapher interface {
	mutate(c context.Context, m string, v map[string]interface{}, t interface{}) (interface{}, error)
	query(c context.Context, q string, v map[string]interface{}, t interface{}) (interface{}, error)
}

type dgraph struct {
	graphqlClient *graphql.Client // TEMP
}

func (d *dgraph) mutate(ctx context.Context, mutation string, variables map[string]interface{}, input interface{}) (interface{}, error) {
	// NOTE: This is a quick-and-dirty implementation that will likely be swapped out for
	// a hand-written version (very similar to this client library) but with allowance for
	// GET requests to be made
	request := graphql.NewRequest(mutation)
	for key, value := range variables {
		request.Var(key, value)
	}

	if err := d.graphqlClient.Run(ctx, request, &input); err != nil {
		return nil, fmt.Errorf("%s: %w", errRunClient, err)
	}

	return input, nil
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
			graphqlClient: graphql.NewClient(dgraphURL), // TEMP
		},
	}
}
