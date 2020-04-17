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
	if err != nil || org == nil {
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
	if err != nil || org == nil {
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
	if err != nil || org == nil {
		return nil, fmt.Errorf("%s: %w", errDeleteOrg, err)
	}

	return org.(*Org), nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, userInput CreateUserInput) (*User, error) {
	mutation := `
	mutation ($userInput: CreateUserInput!) {
		createUser(user: $userInput) {
			id
			firstName
			lastName
			email
			role
		}
	}
	`

	variables := map[string]interface{}{
		"userInput": userInput,
	}

	input := User{}

	user, err := r.Resolver.dgraph.mutate(ctx, mutation, variables, input)
	if err != nil || user == nil {
		return nil, fmt.Errorf("%s: %w", errCreateUser, err)
	}

	return user.(*User), nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, userID string, userInput UpdateUserInput) (*User, error) {
	mutation := `
	mutation ($userID: ID!, $userInput: UpdateUserInput!) {
		updateUser(userID: $userID, user: $userInput) {
			id
			firstName
			lastName
			email
			role
		}
	}
	`

	variables := map[string]interface{}{
		"userID":    userID,
		"userInput": userInput,
	}

	input := User{}

	user, err := r.Resolver.dgraph.mutate(ctx, mutation, variables, input)
	if err != nil || user == nil {
		return nil, fmt.Errorf("%s: %w", errUpdateUser, err)
	}

	return user.(*User), nil
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

	user, err := r.Resolver.dgraph.mutate(ctx, mutation, variables, input)
	if err != nil || user == nil {
		return nil, fmt.Errorf("%s: %w", errAddUser, err)
	}

	return user.(*User), nil
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
	user, err := r.Resolver.dgraph.mutate(ctx, mutation, variables, input)
	if err != nil || user != nil {
		return nil, fmt.Errorf("%s: %w", errRemoveUser, err)
	}

	return user.(*User), nil
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

	user, err := r.Resolver.dgraph.mutate(ctx, mutation, variables, input)
	if err != nil || user == nil {
		return nil, fmt.Errorf("%s: %w", errDeleteUser, err)
	}

	return user.(*User), nil
}

func (r *mutationResolver) CreateItem(ctx context.Context, orgID string, itemInput CreateItemInput) (*Item, error) {
	mutation := `
	mutation ($orgID: ID!, $itemInput: CreateItemInput) {
		createItem(orgID: $orgID, item: $itemInput) {
			id
			description
			parent
			children
		}
	}
	`

	variables := map[string]interface{}{
		"orgID":     orgID,
		"itemInput": itemInput,
	}

	input := Item{}

	item, err := r.Resolver.dgraph.mutate(ctx, mutation, variables, input)
	if err != nil || item == nil {
		return nil, fmt.Errorf("%s: %w", errCreateItem, err)
	}

	return item.(*Item), nil
}

func (r *mutationResolver) CreateItems(ctx context.Context, orgID string, itemsInput []*CreateItemInput) ([]*Item, error) {
	mutation := `
	mutation ($orgID: ID!, $itemsInput: CreateItemInput) {
		createItems(orgID: $orgID, items: $itemsInput) {
			id
			description
			parent
			children
		}
	}
	`

	variables := map[string]interface{}{
		"orgID":      orgID,
		"itemsInput": itemsInput,
	}

	input := []Item{}

	items, err := r.Resolver.dgraph.mutate(ctx, mutation, variables, input)
	if err != nil || items == nil {
		return nil, fmt.Errorf("%s: %w", errCreateItems, err)
	}

	return items.([]*Item), nil
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

	org, err := r.Resolver.dgraph.query(ctx, query, variables, input)
	if err != nil || org == nil {
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

	user, err := r.Resolver.dgraph.query(ctx, query, variables, input)
	if err != nil || user == nil {
		return nil, fmt.Errorf("%s: %w", errReadUser, err)
	}

	return user.(*User), nil
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

	users, err := r.Resolver.dgraph.query(ctx, query, variables, input)
	if err != nil || users == nil {
		return nil, fmt.Errorf("%s: %w", errReadUsers, err)
	}

	return users.([]*User), nil
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

	introType, err := r.Resolver.dgraph.query(ctx, query, variables, input)
	if err != nil || introType == nil {
		return nil, fmt.Errorf("%s: %w", errFilterItems, err)
	}

	return introType.(*introspection.Type), nil
}

func (r *queryResolver) ReadItems(ctx context.Context, itemsInput ReadItemsInput) ([]*Item, error) {
	query := `
	query ($itemsInput: ReadItemsInput!) {
		readItems(items: $itemsInput) {
			id
			description
			parent
			children
		}
	}
	`

	variables := map[string]interface{}{
		"itemsInput": itemsInput,
	}

	input := []Item{}

	items, err := r.Resolver.dgraph.query(ctx, query, variables, input)
	if err != nil || items == nil {
		return nil, fmt.Errorf("%s: %w", errReadItems, err)
	}

	return items.([]*Item), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
