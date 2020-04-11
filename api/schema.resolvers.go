package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql/introspection"
)

func (r *mutationResolver) CreateOrg(ctx context.Context, name string, user CreateUserInput) (*Org, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, orgID string, user CreateUserInput) (*User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUser(ctx context.Context, orgID string, user UpdateUserInput) (*User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUser(ctx context.Context, orgID string, userID string) (*User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateItem(ctx context.Context, orgID string, item CreateItemInput) (*Item, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateItems(ctx context.Context, items []*CreateItemInput) ([]*Item, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ReadOrg(ctx context.Context, orgID string) (*Org, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ReadUser(ctx context.Context, userID string) (*User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListUsers(ctx context.Context, orgID string) ([]*User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) FilterItems(ctx context.Context, name string) (*introspection.Type, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListItems(ctx context.Context, items ListItemsInput) ([]*Item, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
