package app

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/dgraph-io/dgo/v2/protos/api"
)

func stringPtr(input string) *string {
	return &input
}

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func Test_newClient(t *testing.T) {
	client, err := newClient("localhost:8000")
	if client == nil || err != nil {
		t.Errorf("description: error creating client, client: %+v, error: %s", client, err.Error())
	}
}

func Test_new(t *testing.T) {
	db, err := newDB("localhost:8000")
	if db == nil || err != nil {
		t.Errorf("description: error creating db, db: %+v, error: %s", db, err.Error())
	}
}

type mockTransaction struct {
	mutateResp *api.Response
	mutateErr  error
	queryResp  *api.Response
	queryErr   error
}

func (mt *mockTransaction) Mutate(context.Context, *api.Mutation) (*api.Response, error) {
	return mt.mutateResp, mt.mutateErr
}

func (mt *mockTransaction) QueryWithVars(context.Context, string, map[string]string) (*api.Response, error) {
	return mt.queryResp, mt.queryErr
}

type mockDgraph struct {
	mockTransaction *mockTransaction
	setSchemaErr    error
}

func (md *mockDgraph) getTransaction() transaction {
	return md.mockTransaction
}

func (md *mockDgraph) setSchema(string) error {
	return md.setSchemaErr
}

func Test_alter(t *testing.T) {
	tests := []struct {
		desc         string
		setSchemaErr error
	}{
		{
			desc:         "dgraph set schema error",
			setSchemaErr: errors.New("mock set schema error"),
		},
		{
			desc:         "successful invocation",
			setSchemaErr: nil,
		},
	}

	for _, test := range tests {
		db := &db{
			dgraph: &mockDgraph{
				setSchemaErr: test.setSchemaErr,
			},
		}

		testErr := fmt.Errorf("%s: %w", erralter, test.setSchemaErr)

		err := db.alter(`
			name: string

			type Specimen {
				name: string
			}
		`)
		if err != nil && errors.Is(err, testErr) {
			t.Errorf("description: %s, expected: %v, received: %v", test.desc, testErr, err)
		}
	}
}

func Test_mutate(t *testing.T) {
	tests := []struct {
		desc       string
		mutateResp *api.Response
		mutateErr  error
	}{
		{
			desc:       "dgraph mutate method error",
			mutateResp: nil,
			mutateErr:  errors.New("mock mutate error"),
		},
		{
			desc:       "successful invocation",
			mutateResp: &api.Response{},
			mutateErr:  nil,
		},
	}

	for _, test := range tests {
		m := &mutation{
			Content: []byte(`{"type":"blood"}`),
		}

		db := &db{
			dgraph: &mockDgraph{
				mockTransaction: &mockTransaction{
					mutateResp: test.mutateResp,
					mutateErr:  test.mutateErr,
				},
			},
		}

		testErr := fmt.Errorf("%s: %w", errMutate, test.mutateErr)

		err := db.mutate(m)
		if err != nil && errors.Is(err, testErr) {
			t.Errorf("description: %s, expected: %v, received: %v", test.desc, testErr, err)
		}
	}
}

func Test_query(t *testing.T) {
	tests := []struct {
		desc      string
		queryResp *api.Response
		queryErr  error
	}{
		{
			desc:      "dgraph query method error",
			queryResp: nil,
			queryErr:  errors.New("mock query error"),
		},
		{
			desc: "successful invocation",
			queryResp: &api.Response{
				Json: []byte(`{"response":"data"}`),
			},
			queryErr: nil,
		},
	}

	q := &query{
		Content: stringPtr(`
			specimens(func: has(name)) {
				name
			}
		`),
	}

	for _, test := range tests {
		db := &db{
			dgraph: &mockDgraph{
				mockTransaction: &mockTransaction{
					queryResp: test.queryResp,
					queryErr:  test.queryErr,
				},
			},
		}

		testErr := fmt.Errorf("%s: %w", errQuery, test.queryErr)

		_, err := db.query(q)
		if err != nil && errors.Is(err, testErr) {
			t.Errorf("description: %s, expected: %v, received: %v", test.desc, testErr, err)
		}
	}
}
