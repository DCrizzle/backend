package main

import (
	"net/http"
	"testing"
)

func Test_newServer(t *testing.T) {
	testHandler := func(w http.ResponseWriter, r *http.Request) {}
	server := newServer(testHandler)
	if server == nil {
		t.Error("error creating server")
	}
}
