package app

import (
	"testing"
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
