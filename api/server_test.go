package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_handleOrg(t *testing.T) {
	tests := []struct {
		desc string
		path string
		code int
		resp string
	}{
		{
			desc: "successful invocation",
			path: "/org/123",
			code: 200,
			resp: `{"data":"test"}`,
		},
	}

	for _, test := range tests {
		req, err := http.NewRequest("GET", test.path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rec := httptest.NewRecorder()
		handler := http.HandlerFunc(handleOrg)

		handler.ServeHTTP(rec, req)

		if rec.Code != test.code {
			t.Errorf("description: %s, received: %d, expected: %d", test.desc, rec.Code, test.code)
		}
	}
}

func Test_newServer(t *testing.T) {
	s := newServer("127.0.0.1:8080")
	if s == nil {
		t.Errorf("description: error creating server, received: %+v", s)
	}
}
