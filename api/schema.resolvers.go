package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql/introspection"
)

func (r *mutationResolver) CreateOrg(ctx context.Context, name string) (*Org, error) {
	mutation := `
	mutation ($name: String!) {
		createOrg(name: $name) {
			id
			name
		}
	}
	`

	variables := map[string]interface{}{
		"name": name,
	}

	input := Org{}

	org, err := r.Resolver.dgraph.mutate(ctx, mutation, variables, input)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errCreateOrg, err)
	}

	return org.(*Org), nil
}

func (r *mutationResolver) UpdateOrg(ctx context.Context, orgID string, name string) (*Org, error) {
	mutation := `
	mutation ($orgID: ID!, $name: String!) {
		updateOrg(orgID: $orgID, name: $name) {
			id
			name
		}
	}
	`

	variables := map[string]interface{}{
		"orgID": orgID,
		"name":  name,
	}

	input := Org{}

	org, err := r.Resolver.dgraph.mutate(ctx, mutation, variables, input)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUpdateOrg, err)
	}

	return org.(*Org), nil
}

func (r *mutationResolver) DeleteOrg(ctx context.Context, orgID string) (*Org, error) {
	mutation := `
	mutation ($orgID: ID!) {
		deleteOrg(orgID: $orgID) {
			id
			name
		}
	}
	`

	variables := map[string]interface{}{
		"orgID": orgID,
	}

	input := Org{}

	org, err := r.Resolver.dgraph.mutate(ctx, mutation, variables, input)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errDeleteOrg, err)
	}

	return org.(*Org), nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, user CreateUserInput) (*User, error) {
	mutation := `
	mutation ($user: CreateUserInput!) {
		createUser(user: $user) {
			id
			firstName
			lastName
			email
			role
		}
	}
	`

	variables := map[string]interface{}{
		"user": user,
	}

	input := User{}

	org, err := r.Resolver.dgraph.mutate(ctx, mutation, variables, input)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errCreateUser, err)
	}

	return org.(*User), nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, userID string, user UpdateUserInput) (*User, error) {
	mutation := `
	mutation ($userID: ID!, $user: UpdateUserInput!) {
		updateUser(userID: $userID, user: $user) {
			id
			firstName
			lastName
			email
			role
		}
	}
	`

	variables := map[string]interface{}{
		"userID": userID,
		"user":   user,
	}

	input := User{}

	org, err := r.Resolver.dgraph.mutate(ctx, mutation, variables, input)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUpdateUser, err)
	}

	return org.(*User), nil
}

func (r *mutationResolver) AddUser(ctx context.Context, orgID string, userID string) (*User, error) {
	mutation := `
	mutation ($orgID: ID!, $userID: ID!) {
		addUser(orgID: $orgID, userID: $userID) {
			id
			firstName
			lastName
			email
			orgs
			role
		}
	}
	`

	variables := map[string]interface{}{
		"orgID":  orgID,
		"userID": userID,
	}

	input := User{}

	org, err := r.Resolver.dgraph.mutate(ctx, mutation, variables, input)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errAddUser, err)
	}

	return org.(*User), nil
}

func (r *mutationResolver) RemoveUser(ctx context.Context, orgID string, userID string) (*User, error) {
	mutation := `
	mutation ($orgID: ID!, $userID: ID!) {
		removeUser(orgID: $orgID, userID: $userID) {
			id
			firstName
			lastName
			email
			orgs
			role
		}
	}
	`

	variables := map[string]interface{}{
		"orgID":  orgID,
		"userID": userID,
	}

	input := User{}

	org, err := r.Resolver.dgraph.mutate(ctx, mutation, variables, input)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errRemoveUser, err)
	}

	return org.(*User), nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, userID string) (*User, error) {
	mutation := `
	mutation ($userID: ID!) {
		deleteUser(userID: $userID) {
			id
			firstName
			lastName
			email
			orgs
			role
		}
	}
	`

	variables := map[string]interface{}{
		"userID": userID,
	}

	input := User{}

	org, err := r.Resolver.dgraph.mutate(ctx, mutation, variables, input)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errDeleteUser, err)
	}

	return org.(*User), nil
}

func (r *mutationResolver) CreateItem(ctx context.Context, orgID string, item CreateItemInput) (*Item, error) {
	mutation := `
	mutation ($orgID: ID!, $item: CreateItemInput) {
		createItem(orgID: $orgID, item: $item) {
			id
			description
			parent
			children
		}
	}
	`

	variables := map[string]interface{}{
		"orgID": orgID,
		"item":  item,
	}

	input := Item{}

	org, err := r.Resolver.dgraph.mutate(ctx, mutation, variables, input)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errCreateItem, err)
	}

	return org.(*Item), nil
}

func (r *mutationResolver) CreateItems(ctx context.Context, orgID string, items []*CreateItemInput) ([]*Item, error) {
	mutation := `
	mutation ($orgID: ID!, $items: CreateItemInput) {
		createItems(orgID: $orgID, items: $items) {
			id
			description
			parent
			children
		}
	}
	`

	variables := map[string]interface{}{
		"orgID": orgID,
		"items": items,
	}

	input := []Item{}

	org, err := r.Resolver.dgraph.mutate(ctx, mutation, variables, input)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errCreateItems, err)
	}

	return org.([]*Item), nil
}

func (r *queryResolver) ReadOrg(ctx context.Context, orgID string) (*Org, error) {
	query := `
	query ($orgID: ID!) {
		readOrg(orgID: $orgID) {
			id
			name
			users
		}
	}
	`

	variables := map[string]interface{}{
		"orgID": orgID,
	}

	input := Org{}

	org, err := r.Resolver.dgraph.mutate(ctx, query, variables, input)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errReadOrg, err)
	}

	return org.(*Org), nil
}

func (r *queryResolver) ReadUser(ctx context.Context, userID string) (*User, error) {
	query := `
	query ($userID: ID!) {
		readUser(userID: $userID) {
			id
			firstName
			lastName
			email
			orgs
			role
		}
	}
	`

	variables := map[string]interface{}{
		"userID": userID,
	}

	input := User{}

	org, err := r.Resolver.dgraph.mutate(ctx, query, variables, input)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errReadUser, err)
	}

	return org.(*User), nil
}

func (r *queryResolver) ReadUsers(ctx context.Context, orgID string) ([]*User, error) {
	query := `
	query ($orgID: ID!) {
		readUsers(orgID: $orgID) {
			id
			firstName
			lastName
			email
			orgs
			role
		}
	}
	`

	variables := map[string]interface{}{
		"orgID": orgID,
	}

	input := []User{}

	org, err := r.Resolver.dgraph.mutate(ctx, query, variables, input)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errReadUsers, err)
	}

	return org.([]*User), nil
}

func (r *queryResolver) FilterItems(ctx context.Context, name string) (*introspection.Type, error) {
	// NOTE: https://graphql.org/learn/introspection/
	query := `
	query ($name: String!) {
		__type(name: $name) {
			name
			fields {
				name
				type
			}
		}
	}
	`

	variables := map[string]interface{}{
		"name": name,
	}

	input := introspection.Type{}

	org, err := r.Resolver.dgraph.mutate(ctx, query, variables, input)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errFilterItems, err)
	}

	return org.(*introspection.Type), nil
}

func (r *queryResolver) ReadItems(ctx context.Context, items ReadItemsInput) ([]*Item, error) {
	query := `
	query ($items: ReadItemsInput!) {
		readItems(orgID: $orgID) {
			id
			description
			parent
			children
		}
	}
	`

	variables := map[string]interface{}{
		"items": items,
	}

	input := []Item{}

	org, err := r.Resolver.dgraph.mutate(ctx, query, variables, input)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errReadItems, err)
	}

	return org.([]*Item), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
