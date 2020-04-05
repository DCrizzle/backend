package api

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func Test_helper(t *testing.T) {
	successPath := "org/1/graphql"
	errorPath := "org/2/graphql"

	tests := []struct {
		desc string
		path string
		code int
	}{
		{
			desc: "error invoking graphql database endpoint",
			path: errorPath,
			code: http.StatusBadRequest,
		},
		{
			desc: "successful invocation",
			path: successPath,
			code: http.StatusOK,
		},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/"+successPath, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data":{"test":"data"}}`))
	})
	mux.HandleFunc("/"+errorPath, func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "test error", http.StatusBadRequest)
	})

	h := new(httpHelp)

	server := httptest.NewServer(mux)
	serverURL, _ := url.Parse(server.URL + "/")

	payload := strings.NewReader("test payload")

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			resp, _ := h.post(serverURL.String()+test.path, "application/json", payload)

			if resp.StatusCode != test.code {
				t.Errorf("status received: %d, expected: %d", resp.StatusCode, test.code)
			}
		})

		t.Run(test.desc, func(t *testing.T) {
			resp, _ := h.get(serverURL.String() + test.path)

			if resp.StatusCode != test.code {
				t.Errorf("status received: %d, expected: %d", resp.StatusCode, test.code)
			}
		})
	}
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
		desc string
		code int
	}{
		// {
		// 	desc: "token not in database",
		// 	code: 403,
		// },
		{
			desc: "successful invocation",
			code: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			h := new(help)

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

type testClient struct {
	clientResp *http.Response
	clientErr  error
}

func (tc *testClient) post(url, contentType string, payload io.Reader) (*http.Response, error) {
	return tc.clientResp, tc.clientErr
}

func (tc *testClient) get(url string) (*http.Response, error) {
	return tc.clientResp, tc.clientErr
}

func Test_mutate(t *testing.T) {
	tests := []struct {
		desc       string
		path       string
		clientResp *http.Response
		clientErr  error
		addr       string
		code       int
		resp       string
	}{
		{
			desc:       "error invoking database endpoint",
			path:       "/graphql",
			clientResp: nil,
			clientErr:  errors.New("mock post error"),
			addr:       "test.com",
			code:       http.StatusInternalServerError,
			resp:       errMutateDB + "\n",
		},
		{
			desc: "successful invocation",
			path: "/graphql",
			clientResp: &http.Response{
				Body: ioutil.NopCloser(strings.NewReader(`{"data":{"test":"data"}}`)),
			},
			clientErr: nil,
			addr:      "test.com",
			code:      http.StatusOK,
			resp:      `{"data":{"test":"data"}}`,
		},
	}

	for _, test := range tests {
		h := &help{
			client: &testClient{
				clientResp: test.clientResp,
				clientErr:  test.clientErr,
			},
			addr: test.addr,
		}

		t.Run(test.desc, func(t *testing.T) {
			req, err := http.NewRequest("POST", test.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rec := httptest.NewRecorder()
			handler := http.HandlerFunc(h.mutate())

			handler.ServeHTTP(rec, req)

			if rec.Code != test.code {
				t.Errorf("code received: %d, expected: %d", rec.Code, test.code)
			}

			if rec.Body.String() != test.resp {
				t.Errorf("body received: %s, expected: %s", rec.Body.String(), test.resp)
			}
		})
	}
}

func Test_query(t *testing.T) {
	tests := []struct {
		desc       string
		path       string
		clientResp *http.Response
		clientErr  error
		addr       string
		code       int
		resp       string
	}{
		{
			desc:       "error invoking database endpoint",
			path:       "/graphql?query=test",
			clientResp: nil,
			clientErr:  errors.New("mock get error"),
			addr:       "test.com",
			code:       http.StatusInternalServerError,
			resp:       errQueryDB + "\n",
		},
		{
			desc: "successful invocation",
			path: "/graphql?query=test",
			clientResp: &http.Response{
				Body: ioutil.NopCloser(strings.NewReader(`{"data":{"test":"data"}}`)),
			},
			addr: "test.com",
			code: http.StatusOK,
			resp: `{"data":{"test":"data"}}`,
		},
	}

	for _, test := range tests {
		h := &help{
			client: &testClient{
				clientResp: test.clientResp,
				clientErr:  test.clientErr,
			},
			addr: test.addr,
		}

		t.Run(test.desc, func(t *testing.T) {
			req, err := http.NewRequest("GET", test.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rec := httptest.NewRecorder()
			handler := http.HandlerFunc(h.query())

			handler.ServeHTTP(rec, req)

			if rec.Code != test.code {
				t.Errorf("code received: %d, expected: %d", rec.Code, test.code)
			}

			if rec.Body.String() != test.resp {
				t.Errorf("body received: %s, expected: %s", rec.Body.String(), test.resp)
			}
		})
	}
}

func TestNewServer(t *testing.T) {
	s, err := NewServer("127.0.0.1:8080")
	if s == nil || err != nil {
		t.Errorf("error received: %+v", s)
	}
}
