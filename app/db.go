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
	errAlter     = "set schema error"
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

// Database is the public interface for the Dgraph wrapper
type Database interface {
	Alter(schema string) error
	Mutate(input *Mutation) error
	Query(input *Query) ([]byte, error)
}

type logger interface{}

// DB implements the public Database interface
type DB struct {
	dgraph dgraph
	logger logger
}

// NewDB generates an implementation of the Database interface
func NewDB(host string) (*DB, error) {
	client, err := newClient(host)
	if err != nil {
		return nil, err
	}

	// outline:
	// [ ] set schema w/ required indices

	return &DB{
		dgraph: &dg{
			client: client,
		},
	}, nil
}

// Alter is a method for updating the Dgraph schema
func (db *DB) Alter(schema string) error {
	if err := db.dgraph.setSchema(schema); err != nil {
		return fmt.Errorf("%s: %w", errAlter, err)
	}

	return nil
}

// Mutate is a method for executing mutation operations on the Dgraph database
func (db *DB) Mutate(input *Mutation) error {
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

// Query is a method for executing query operations on the Dgraph database
func (db *DB) Query(input *Query) ([]byte, error) {
	ctx := context.Background()
	resp, err := db.dgraph.getTransaction().QueryWithVars(ctx, *input.Content, input.Variables)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errQuery, err)
	}

	return resp.Json, nil
}
