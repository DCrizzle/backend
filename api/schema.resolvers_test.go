package api

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type mockHTTPClient struct {
	postResp *http.Response
	postErr  error
	getResp  *http.Response
	getErr   error
}

func (mhc *mockHTTPClient) Post(url, contentType string, body io.Reader) (*http.Response, error) {
	return mhc.postResp, mhc.postErr
}

func (mhc *mockHTTPClient) Get(url string) (*http.Response, error) {
	return mhc.getResp, mhc.getErr
}

func TestCreateOrg(t *testing.T) {
	tests := []struct {
		desc     string
		name     string
		postResp *http.Response
		postErr  error
		err      error
	}{
		{
			desc:     "error posting mutation to database",
			name:     "temple-archives",
			postResp: nil,
			postErr:  errors.New("mock create org error"),
			err:      fmt.Errorf("%s: %w", errCreateOrg, errors.New("mock create org error")),
		},
		{
			desc: "error parsing database response body",
			name: "temple-archives",
			postResp: &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(strings.NewReader("")),
			},
			postErr: nil,
			err:     fmt.Errorf("%s: %w", errUnmarshalJSON, errors.New("unexpected end of JSON input")),
		},
		{
			desc: "successful invocation",
			name: "temple-archives",
			postResp: &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(strings.NewReader(`{name: "lost-twenty"}`)),
			},
			postErr: nil,
			err:     nil,
		},
	}

	ctx := context.Background()
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			m := &mutationResolver{
				&Resolver{
					httpClient: &mockHTTPClient{
						postResp: test.postResp,
						postErr:  test.postErr,
					},
				},
			}

			org, err := m.CreateOrg(ctx, test.name)
			if test.err != nil && err.Error() != test.err.Error() {
				t.Errorf("error received: %v, expected: %v", err, test.err)
			}

			if org != nil && org.Name != test.name {
				t.Errorf("org named received: %s, expected: %s", org.Name, test.name)
			}
		})
	}
}
