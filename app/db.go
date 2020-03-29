package app

import (
	"context"
	"fmt"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
)

const ( // NOTE: error generating/processing may need updating
	errNewClient = "new client error"
	errNewDB     = "new db instance error"
	erralter     = "set schema error"
	errMutate    = "mutate execution error"
	errQuery     = "query execution error"
)

type transaction interface {
	Mutate(context.Context, *api.Mutation) (*api.Response, error)
	QueryWithVars(context.Context, string, map[string]string) (*api.Response, error)
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

type dgraph interface {
	getTransaction() transaction
	setSchema(string) error
}

type dg struct {
	client *dgo.Dgraph
}

func (dg *dg) getTransaction() transaction {
	return dg.client.NewTxn()
}

func (dg *dg) setSchema(schema string) error {
	ctx := context.Background()
	op := &api.Operation{
		Schema: schema,
	}

	return dg.client.Alter(ctx, op) // NOTE: add logic to wait/block until updated completes
}

type logger interface{}

type db struct {
	dgraph dgraph
	logger logger
}

func newDB(host string) (*db, error) {
	client, err := newClient(host)
	if err != nil {
		return nil, err
	}

	// outline:
	// [ ] set schema w/ required indices

	return &db{
		dgraph: &dg{
			client: client,
		},
	}, nil
}

func (db *db) alter(schema string) error {
	if err := db.dgraph.setSchema(schema); err != nil {
		return fmt.Errorf("%s: %w", erralter, err)
	}

	return nil
}

func (db *db) mutate(input *mutation) error {
	ctx := context.Background()
	m := &api.Mutation{
		CommitNow: true,
		SetJson:   input.Content,
	}

	resp, err := db.dgraph.getTransaction().Mutate(ctx, m)
	if err != nil {
		return fmt.Errorf("%s: %w", errMutate, err)
	}
	_ = resp.Uids

	return nil
}

func (db *db) query(input *query) ([]byte, error) {
	ctx := context.Background()
	resp, err := db.dgraph.getTransaction().QueryWithVars(ctx, *input.Content, input.Variables)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errQuery, err)
	}

	return resp.Json, nil
}
