package api

import (
	"context"
	"errors"
	"testing"

	"github.com/machinebox/graphql"
)

func TestResolver_newRequest(t *testing.T) {
	request := newRequest("operation", make(map[string]interface{}))
	if request == nil {
		t.Errorf("error creating request, received: %v\n", request)
	}
}

type mockGraphQLClient struct {
	err error
}

func (m *mockGraphQLClient) Run(ctx context.Context, req *graphql.Request, resp interface{}) error {
	return m.err
}

func TestResolver_mutate(t *testing.T) {
	tests := []struct {
		desc string
		err  error
	}{
		{
			desc: "error running graphql client",
			err:  errors.New("mock graphql client error"),
		},
		{
			desc: "succesful invocation",
			err:  nil,
		},
	}

	for _, test := range tests {
		d := &dgraph{
			graphQLClient: &mockGraphQLClient{
				err: test.err,
			},
		}

		ctx := context.Background()
		_, err := d.mutate(ctx, "mutation", make(map[string]interface{}), &Org{})
		if errors.Unwrap(err) != test.err {
			t.Errorf("error received: %v, expected: %v", err, test.err)
		}
	}
}

func TestResolver_query(t *testing.T) {
	tests := []struct {
		desc string
		err  error
	}{
		{
			desc: "error running graphql client",
			err:  errors.New("mock graphql client error"),
		},
		{
			desc: "succesful invocation",
			err:  nil,
		},
	}

	for _, test := range tests {
		d := &dgraph{
			graphQLClient: &mockGraphQLClient{
				err: test.err,
			},
		}

		ctx := context.Background()
		_, err := d.query(ctx, "query", make(map[string]interface{}), &Org{})
		if errors.Unwrap(err) != test.err {
			t.Errorf("error received: %v, expected: %v", err, test.err)
		}
	}
}
