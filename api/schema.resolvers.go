package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	// "encoding/json"
	"fmt"
	// "io/ioutil"
	// "net/http"
	// "strings"

	"github.com/99designs/gqlgen/graphql/introspection"
)

func (r *mutationResolver) CreateOrg(ctx context.Context, name string) (*Org, error) {
	// mutation := fmt.Sprintf(`
	// mutation {
	// 	createOrg(name: %s) {
	// 		id
	// 		name
	// 	}
	// }
	// `, name)
	//
	// resp, err := r.httpClient.Post(r.dbURL, "application/json", strings.NewReader(mutation))
	// if err != nil || resp.StatusCode != http.StatusOK {
	// 	return nil, fmt.Errorf("%s: %w", errCreateOrg, err)
	// }
	//
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, fmt.Errorf("%s: %w", errReadBytes, err)
	// }
	//
	// org := &Org{}
	// if err := json.Unmarshal(body, org); err != nil {
	// 	return nil, fmt.Errorf("%s: %w", errUnmarshalJSON, err)
	// }
	//
	// return org, nil
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateOrg(ctx context.Context, orgID string, name string) (*Org, error) {
	// mutation := fmt.Sprintf(`
	// mutation {
	// 	updateOrg(orgID: %s, name: %s) {
	// 		id
	// 		name
	// 	}
	// }
	// `, orgID, name)
	//
	// resp, err := r.httpClient.Post(r.dbURL, "application/json", strings.NewReader(mutation))
	// if err != nil || resp.StatusCode != http.StatusOK {
	// 	return nil, fmt.Errorf("%s: %w", errDeleteOrg, err)
	// }
	//
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, fmt.Errorf("%s: %w", errReadBytes, err)
	// }
	//
	// org := &Org{}
	// if err := json.Unmarshal(body, org); err != nil {
	// 	return nil, fmt.Errorf("%s: %w", errUnmarshalJSON, err)
	// }
	//
	// return org, nil
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteOrg(ctx context.Context, orgID string) (*Org, error) {
	// mutation := fmt.Sprintf(`
	// mutation {
	// 	deleteOrg(orgID: %s) {
	// 		id
	// 		name
	// 	}
	// }
	// `, orgID)
	//
	// resp, err := r.httpClient.Post(r.dbURL, "application/json", strings.NewReader(mutation))
	// if err != nil || resp.StatusCode != http.StatusOK {
	// 	return nil, fmt.Errorf("%s: %w", errDeleteOrg, err)
	// }
	//
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, fmt.Errorf("%s: %w", errReadBytes, err)
	// }
	//
	// org := &Org{}
	// if err := json.Unmarshal(body, org); err != nil {
	// 	return nil, fmt.Errorf("%s: %w", errUnmarshalJSON, err)
	// }
	//
	// return org, nil
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, user CreateUserInput) (*User, error) {
	// mutation := fmt.Sprintf(`
	// mutation {
	// 	createUser(user: %v) {
	// 		id
	// 		name
	// 	}
	// }
	// `, user)
	//
	// resp, err := r.httpClient.Post(r.dbURL, "application/json", strings.NewReader(mutation))
	// if err != nil || resp.StatusCode != http.StatusOK {
	// 	return nil, fmt.Errorf("%s: %w", errCreateUser, err)
	// }
	//
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, fmt.Errorf("%s: %w", errReadBytes, err)
	// }
	//
	// newUser := &User{}
	// if err := json.Unmarshal(body, newUser); err != nil {
	// 	return nil, fmt.Errorf("%s: %w", errUnmarshalJSON, err)
	// }
	//
	// return newUser, nil
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

func (r *queryResolver) ReadItems(ctx context.Context, items ReadItemsInput) ([]*Item, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
