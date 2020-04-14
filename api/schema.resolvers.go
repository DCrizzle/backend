package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/99designs/gqlgen/graphql/introspection"
)

const (
	errCreateOrg     = "error creating org"
	errReadBytes     = "error reading body"
	errUnmarshalJSON = "error unmarshalling json"
)

func (r *mutationResolver) CreateOrg(ctx context.Context, name string) (*Org, error) {
	mutation := fmt.Sprintf(`
	mutation {
		createOrg(name: %s) {
			id
		}
	}
	`, name)

	resp, err := r.httpClient.Post(r.dbURL, "application/json", strings.NewReader(mutation))
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: %w", errCreateOrg, err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errReadBytes, err)
	}

	org := &Org{}
	if err := json.Unmarshal(body, org); err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalJSON, err)
	}

	return org, nil
}

func (r *mutationResolver) DeleteOrg(ctx context.Context, orgID string) (*Org, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateOrg(ctx context.Context, orgID string, name string) (*Org, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, user CreateUserInput) (*User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUser(ctx context.Context, userID string, user UpdateUserInput) (*User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddUser(ctx context.Context, orgID string, userID string) (*User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveUser(ctx context.Context, orgID string, userID string) (*User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUser(ctx context.Context, userID string) (*User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateItem(ctx context.Context, orgID string, item CreateItemInput) (*Item, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateItems(ctx context.Context, orgID string, items []*CreateItemInput) ([]*Item, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ReadOrg(ctx context.Context, orgID string) (*Org, error) {
	// outline:
	// [ ] create url query string
	// [ ] pass id into string
	// [ ] reference http client field
	// [ ] execute get request
	// [ ] create org object
	// [ ] return org / nil

	panic(fmt.Errorf("not implemented")) // TEMP
}

func (r *queryResolver) ReadUser(ctx context.Context, userID string) (*User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ReadUsers(ctx context.Context, orgID string) ([]*User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) FilterItems(ctx context.Context, name string) (*introspection.Type, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ReadItems(ctx context.Context, items ListItemsInput) ([]*Item, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) ListUsers(ctx context.Context, orgID string) ([]*User, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) ListItems(ctx context.Context, items ListItemsInput) ([]*Item, error) {
	panic(fmt.Errorf("not implemented"))
}
