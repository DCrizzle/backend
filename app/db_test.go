package app

import (
	"context"
	"errors"
	"testing"

	"github.com/dgraph-io/dgo/v2/protos/api"
)

func Test_newClient(t *testing.T) {
	client, err := newClient("localhost:8000")
	if client == nil || err != nil {
		t.Errorf("description: error creating client, client: %+v, error: %s", client, err.Error())
	}
}

func Test_new(t *testing.T) {
	db, err := newDB()
	if db == nil || err != nil {
		t.Errorf("description: error creating db, db: %+v, error: %s", db, err.Error())
	}
}

type mockTransaction struct {
	mutateResp *api.Response
	mutateErr  error
}

func (mt *mockTransaction) Mutate(context.Context, *api.Mutation) (*api.Response, error) {
	return mt.mutateResp, mt.mutateErr
}

type mockDgraph struct {
	mockTransaction *mockTransaction
}

func (md *mockDgraph) transaction() transaction {
	return md.mockTransaction
}

func Test_mutate(t *testing.T) {
	tests := []struct {
		desc       string
		mutateResp *api.Response
		mutateErr  error
		err        error
	}{
		{
			desc:       "dgraph mutate method error",
			mutateResp: nil,
			mutateErr:  errors.New("mock mutate error"),
			err:        errors.New("mutate execution error: mock mutate error"),
		},
		{
			desc:       "successful invocation",
			mutateResp: &api.Response{},
			mutateErr:  nil,
			err:        nil,
		},
	}

	for _, test := range tests {
		specimen := []byte(`{"type":"blood"}`)

		db := &db{
			dgraph: &mockDgraph{
				mockTransaction: &mockTransaction{
					mutateResp: test.mutateResp,
					mutateErr:  test.mutateErr,
				},
			},
		}

		err := db.mutate(specimen)
		if err != nil && err.Error() != test.err.Error() {
			t.Errorf("description: %s, expected: %v, received: %v", test.desc, test.err, err)
		}
	}
}
