package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

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
		desc string
		code int
	}{
		// {
		// 	desc: "token not in database",
		// 	code: 403,
		// },
		{
			desc: "successful invocation",
			code: 200,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			h := &helper{}

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

func Test_newServer(t *testing.T) {
	s, err := newServer("127.0.0.1:8080")
	if s == nil || err != nil {
		t.Errorf("error received: %+v", s)
	}
}
