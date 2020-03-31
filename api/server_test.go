package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/forstmeier/tbd/app"
)

type testDB struct {
	errAlter  error
	errMutate error
	respQuery []byte
	errQuery  error
}

func (db *testDB) Alter(schema string) error {
	return db.errAlter
}

func (db *testDB) Mutate(input *app.Mutation) error {
	return db.errMutate
}

func (db *testDB) Query(input *app.Query) ([]byte, error) {
	return db.respQuery, db.errQuery
}

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
		db := &testDB{}
		handler := http.HandlerFunc(handleOrg(db))

		handler.ServeHTTP(rec, req)

		if rec.Code != test.code {
			t.Errorf("description: %s, received: %d, expected: %d", test.desc, rec.Code, test.code)
		}
	}
}

func Test_handleMutation(t *testing.T) {
	tests := []struct {
		desc string
		path string
		code int
		resp string
	}{
		{
			desc: "successful invocation",
			path: "/org/123/db",
			code: 200,
			resp: `{"data":"test"}`,
		},
	}

	for _, test := range tests {
		req, err := http.NewRequest("POST", test.path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rec := httptest.NewRecorder()
		db := &testDB{}
		handler := http.HandlerFunc(handleMutation(db))

		handler.ServeHTTP(rec, req)

		if rec.Code != test.code {
			t.Errorf("description: %s, received: %d, expected: %d", test.desc, rec.Code, test.code)
		}
	}
}

func Test_handleQuery(t *testing.T) {
	tests := []struct {
		desc string
		path string
		code int
		resp string
	}{
		{
			desc: "successful invocation",
			path: "/org/123/db",
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
		db := &testDB{}
		handler := http.HandlerFunc(handleQuery(db))

		handler.ServeHTTP(rec, req)

		if rec.Code != test.code {
			t.Errorf("description: %s, received: %d, expected: %d", test.desc, rec.Code, test.code)
		}
	}
}

func Test_newServer(t *testing.T) {
	s, err := newServer("127.0.0.1:8080")
	if s == nil || err != nil {
		t.Errorf("description: error creating server, received: %+v", s)
	}
}
