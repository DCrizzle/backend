package server

import (
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	testHandler := func(w http.ResponseWriter, r *http.Request) {}
	server := New(testHandler, testHandler)
	if server == nil {
		t.Error("error creating server")
	}
}
