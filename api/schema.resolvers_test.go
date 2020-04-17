package api

import (
	"context"
	"errors"
	"testing"

	"github.com/99designs/gqlgen/graphql/introspection"
)

type mockDgraph struct {
	mutateResp interface{}
	mutateErr  error
	queryResp  interface{}
	queryErr   error
}

func (m *mockDgraph) mutate(ctx context.Context, mutation string, variables map[string]interface{}, input interface{}) (interface{}, error) {
	if m.mutateErr != nil {
		return nil, m.mutateErr
	}

	switch v := input.(type) {
	case Org:
		return &Org{}, nil
	case User:
		return &User{}, nil
	case Item:
		return &Item{}, nil
	case []Item:
		return []*Item{}, nil
	default:
		return v, nil
	}
}

func (m *mockDgraph) query(ctx context.Context, query string, variables map[string]interface{}, input interface{}) (interface{}, error) {
	if m.queryErr != nil {
		return nil, m.queryErr
	}

	switch v := input.(type) {
	case Org:
		return &Org{}, nil
	case User:
		return &User{}, nil
	case []User:
		return []*User{}, nil
	case introspection.Type:
		return &introspection.Type{}, nil
	case []Item:
		return []*Item{}, nil
	default:
		return v, nil
	}
}

func Test_mutationResolver(t *testing.T) {
	tests := []struct {
		desc      string
		mutateErr error
	}{
		{
			desc:      "error executing dgraph mutation",
			mutateErr: errors.New("mock mutate error"),
		},
		{
			desc:      "successful invocation",
			mutateErr: nil,
		},
	}

	for _, test := range tests {
		ctx := context.Background()
		m := &mutationResolver{
			&Resolver{
				dgraph: &mockDgraph{
					mutateErr: test.mutateErr,
				},
			},
		}

		t.Run("create org:"+test.desc, func(t *testing.T) {
			_, err := m.CreateOrg(ctx, "orgName")
			if errors.Unwrap(err) != test.mutateErr {
				t.Errorf("error received: %v, expected: %v", err, test.mutateErr)
			}
		})

		t.Run("update org:"+test.desc, func(t *testing.T) {
			_, err := m.UpdateOrg(ctx, "orgID", "orgName")
			if errors.Unwrap(err) != test.mutateErr {
				t.Errorf("error received: %v, expected: %v", err, test.mutateErr)
			}
		})

		t.Run("delete org:"+test.desc, func(t *testing.T) {
			_, err := m.DeleteOrg(ctx, "orgID")
			if errors.Unwrap(err) != test.mutateErr {
				t.Errorf("error received: %v, expected: %v", err, test.mutateErr)
			}
		})

		t.Run("create user:"+test.desc, func(t *testing.T) {
			_, err := m.CreateUser(ctx, CreateUserInput{})
			if errors.Unwrap(err) != test.mutateErr {
				t.Errorf("error received: %v, expected: %v", err, test.mutateErr)
			}
		})

		t.Run("update user:"+test.desc, func(t *testing.T) {
			_, err := m.UpdateUser(ctx, "userID", UpdateUserInput{})
			if errors.Unwrap(err) != test.mutateErr {
				t.Errorf("error received: %v, expected: %v", err, test.mutateErr)
			}
		})

		t.Run("add user:"+test.desc, func(t *testing.T) {
			_, err := m.AddUser(ctx, "orgID", "userID")
			if errors.Unwrap(err) != test.mutateErr {
				t.Errorf("error received: %v, expected: %v", err, test.mutateErr)
			}
		})

		t.Run("remove user:"+test.desc, func(t *testing.T) {
			_, err := m.RemoveUser(ctx, "orgID", "userID")
			if errors.Unwrap(err) != test.mutateErr {
				t.Errorf("error received: %v, expected: %v", err, test.mutateErr)
			}
		})

		t.Run("delete user:"+test.desc, func(t *testing.T) {
			_, err := m.DeleteUser(ctx, "userID")
			if errors.Unwrap(err) != test.mutateErr {
				t.Errorf("error received: %v, expected: %v", err, test.mutateErr)
			}
		})

		t.Run("create item:"+test.desc, func(t *testing.T) {
			_, err := m.CreateItem(ctx, "orgID", CreateItemInput{})
			if errors.Unwrap(err) != test.mutateErr {
				t.Errorf("error received: %v, expected: %v", err, test.mutateErr)
			}
		})

		t.Run("create items:"+test.desc, func(t *testing.T) {
			_, err := m.CreateItems(ctx, "orgID", []*CreateItemInput{})
			if errors.Unwrap(err) != test.mutateErr {
				t.Errorf("error received: %v, expected: %v", err, test.mutateErr)
			}
		})
	}
}

func Test_queryResolver(t *testing.T) {
	tests := []struct {
		desc     string
		queryErr error
	}{
		{
			desc:     "error executing dgraph query",
			queryErr: errors.New("mock query error"),
		},
		{
			desc:     "successful invocation",
			queryErr: nil,
		},
	}

	for _, test := range tests {
		ctx := context.Background()
		q := &queryResolver{
			&Resolver{
				dgraph: &mockDgraph{
					queryErr: test.queryErr,
				},
			},
		}

		t.Run("read org:"+test.desc, func(t *testing.T) {
			_, err := q.ReadOrg(ctx, "orgID")
			if errors.Unwrap(err) != test.queryErr {
				t.Errorf("error received: %v, expected: %v", err, test.queryErr)
			}
		})

		t.Run("read user:"+test.desc, func(t *testing.T) {
			_, err := q.ReadUser(ctx, "userID")
			if errors.Unwrap(err) != test.queryErr {
				t.Errorf("error received: %v, expected: %v", err, test.queryErr)
			}
		})

		t.Run("read users:"+test.desc, func(t *testing.T) {
			_, err := q.ReadUsers(ctx, "orgID")
			if errors.Unwrap(err) != test.queryErr {
				t.Errorf("error received: %v, expected: %v", err, test.queryErr)
			}
		})

		t.Run("filter items:"+test.desc, func(t *testing.T) {
			_, err := q.FilterItems(ctx, "type")
			if errors.Unwrap(err) != test.queryErr {
				t.Errorf("error received: %v, expected: %v", err, test.queryErr)
			}
		})

		t.Run("read items:"+test.desc, func(t *testing.T) {
			_, err := q.ReadItems(ctx, ReadItemsInput{})
			if errors.Unwrap(err) != test.queryErr {
				t.Errorf("error received: %v, expected: %v", err, test.queryErr)
			}
		})
	}

}
