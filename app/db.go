package app

import (
	"context"
	// "errors"
	"fmt"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
)

const (
	errNewClient = "new client error"
	errNewDB     = "new db instance error"
	errSetSchema = "set schema error"
	errMutate    = "mutate execution error"
	errQuery     = "query execution error"
)

type transaction interface {
	Mutate(context.Context, *api.Mutation) (*api.Response, error)
	QueryWithVars(context.Context, string, map[string]string) (*api.Response, error)
}

type dgraph interface {
	transaction() transaction
	setSchema(string) error
}

func newClient(target string) (*dgo.Dgraph, error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("%s: %v", errNewClient, err)
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(conn),
	), nil
}

type dg struct {
	client *dgo.Dgraph
}

func (dg *dg) transaction() transaction {
	return dg.client.NewTxn()
}

func (dg *dg) setSchema(schema string) error {
	ctx := context.Background()
	op := &api.Operation{
		Schema: schema,
	}

	return dg.client.Alter(ctx, op)
}

type db struct {
	dgraph dgraph
}

func newDB() (*db, error) {
	client, err := newClient("localhost:9080")
	if err != nil {
		return nil, err
	}

	return &db{
		dgraph: &dg{
			client: client,
		},
	}, nil
}

func (db *db) setSchema(schema string) error {
	if err := db.dgraph.setSchema(schema); err != nil {
		return fmt.Errorf("%s: %w", errSetSchema, err)
	}

	return nil
}

func (db *db) mutate(input []byte) error {
	mu := &api.Mutation{
		CommitNow: true,
	}

	ctx := context.Background()
	mu.SetJson = input
	resp, err := db.dgraph.transaction().Mutate(ctx, mu)
	if err != nil {
		return fmt.Errorf("%s: %w", errMutate, err)
	}
	_ = resp

	return nil
}

func (db *db) query(query string, vars map[string]string) ([]byte, error) {
	ctx := context.Background()
	resp, err := db.dgraph.transaction().QueryWithVars(ctx, query, vars)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", errQuery, err)
	}

	return resp.Json, nil
}
