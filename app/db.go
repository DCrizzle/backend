package app

import (
	"fmt"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
)

const (
	errNewClient = "new client error"
	errNewDB     = "new db instance error"
)

func newClient(target string) (*dgo.Dgraph, error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("%s: %s", errNewClient, err.Error())
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(conn),
	), nil
}

type db struct {
	dgraph *dgo.Dgraph
}

func newDB() (*db, error) {
	client, err := newClient("localhost:9080")
	if err != nil {
		return nil, fmt.Errorf("%s: %s", errNewDB, err.Error())
	}

	return &db{
		dgraph: client,
	}, nil
}

func (db *db) query() {
	// read
}

func (db *db) mutate() {
	// create/update/delete
}
