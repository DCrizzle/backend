package api

import (
	"errors"
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

type mockHandle struct {
	w http.ResponseWriter
	r *http.Request
}

func (mh *mockHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mh.w = w
	mh.r = r
}

func Test_secure(t *testing.T) {
	tests := []struct {
		desc      string
		respQuery []byte
		errQuery  error
		code      int
	}{
		{
			desc:      "token not in database",
			respQuery: nil,
			errQuery:  errors.New("mock query error"),
			code:      403,
		},
		{
			desc:      "successful invocation",
			respQuery: []byte(`"test":"data"`),
			errQuery:  nil,
			code:      200,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			h := &helper{
				db: &testDB{
					respQuery: test.respQuery,
					errQuery:  test.errQuery,
				},
			}

			mh := &mockHandle{}
			hdlr := h.secure(mh)

			req, err := http.NewRequest("GET", "/user/123", nil)
			if err != nil {
				t.Fatal(err)
			}

			rec := httptest.NewRecorder()

			hdlr.ServeHTTP(rec, req)

			if rec.Code != test.code {
				t.Errorf("code received: %d, expected: %d", rec.Code, test.code)
			}
		})
	}
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
		t.Run(test.desc, func(t *testing.T) {
			req, err := http.NewRequest("GET", test.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rec := httptest.NewRecorder()
			db := &testDB{}
			handler := http.HandlerFunc(handleOrg(db))

			handler.ServeHTTP(rec, req)

			if rec.Code != test.code {
				t.Errorf("code received: %d, expected: %d", rec.Code, test.code)
			}
		})
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
		t.Run(test.desc, func(t *testing.T) {
			req, err := http.NewRequest("POST", test.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rec := httptest.NewRecorder()
			db := &testDB{}
			handler := http.HandlerFunc(handleMutation(db))

			handler.ServeHTTP(rec, req)

			if rec.Code != test.code {
				t.Errorf("code received: %d, expected: %d", rec.Code, test.code)
			}
		})
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
		t.Run(test.desc, func(t *testing.T) {
			req, err := http.NewRequest("GET", test.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rec := httptest.NewRecorder()
			db := &testDB{}
			handler := http.HandlerFunc(handleQuery(db))

			handler.ServeHTTP(rec, req)

			if rec.Code != test.code {
				t.Errorf("code received: %d, expected: %d", rec.Code, test.code)
			}
		})
	}
}

func Test_newServer(t *testing.T) {
	s, err := newServer("127.0.0.1:8080")
	if s == nil || err != nil {
		t.Errorf("error received: %+v", s)
	}
}
