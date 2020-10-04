package server

import (
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	testMiddleware := func(http.Handler) http.Handler {
		return nil
	}
	testHandler := func(w http.ResponseWriter, r *http.Request) {}
	server := New(testMiddleware, testHandler, testHandler)
	if server == nil {
		t.Error("error creating server")
	}
}
