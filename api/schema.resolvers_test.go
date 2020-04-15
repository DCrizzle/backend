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
			name:     "jedi-archives",
			postResp: nil,
			postErr:  errors.New("mock create org error"),
			err:      fmt.Errorf("%s: %w", errCreateOrg, errors.New("mock create org error")),
		},
		{
			desc: "error parsing database response body",
			name: "jedi-archives",
			postResp: &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(strings.NewReader("")),
			},
			postErr: nil,
			err:     fmt.Errorf("%s: %w", errUnmarshalJSON, errors.New("unexpected end of JSON input")),
		},
		{
			desc: "successful invocation",
			name: "jedi-archives",
			postResp: &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(strings.NewReader(`{"name": "jedi-archives"}`)),
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
				t.Errorf("org name received: %s, expected: %s", org.Name, test.name)
			}
		})
	}
}

func TestUpdateOrg(t *testing.T) {
	tests := []struct {
		desc     string
		id       string
		name     string
		postResp *http.Response
		postErr  error
		err      error
	}{
		{
			desc:     "error posting mutation to database",
			id:       "jedi-archives",
			name:     "jedi-temple-library",
			postResp: nil,
			postErr:  errors.New("mock create org error"),
			err:      fmt.Errorf("%s: %w", errDeleteOrg, errors.New("mock create org error")),
		},
		{
			desc: "error parsing database response body",
			id:   "jedi-archives",
			name: "jedi-temple-library",
			postResp: &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(strings.NewReader("")),
			},
			postErr: nil,
			err:     fmt.Errorf("%s: %w", errUnmarshalJSON, errors.New("unexpected end of JSON input")),
		},
		{
			desc: "successful invocation",
			id:   "jedi-archives",
			name: "jedi-temple-library",
			postResp: &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(strings.NewReader(`{"id": "jedi-archives","name":"jedi-temple-library"}`)),
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

			org, err := m.UpdateOrg(ctx, test.id, test.name)
			if test.err != nil && err.Error() != test.err.Error() {
				t.Errorf("error received: %v, expected: %v", err, test.err)
			}

			if org != nil && (org.ID != test.id || org.Name != test.name) {
				t.Errorf("org id/name received: %s/%s, expected: %s/%s", org.ID, org.Name, test.id, test.name)
			}
		})
	}
}

func TestDeleteOrg(t *testing.T) {
	tests := []struct {
		desc     string
		id       string
		postResp *http.Response
		postErr  error
		err      error
	}{
		{
			desc:     "error posting mutation to database",
			id:       "jedi-archives",
			postResp: nil,
			postErr:  errors.New("mock create org error"),
			err:      fmt.Errorf("%s: %w", errDeleteOrg, errors.New("mock create org error")),
		},
		{
			desc: "error parsing database response body",
			id:   "jedi-archives",
			postResp: &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(strings.NewReader("")),
			},
			postErr: nil,
			err:     fmt.Errorf("%s: %w", errUnmarshalJSON, errors.New("unexpected end of JSON input")),
		},
		{
			desc: "successful invocation",
			id:   "jedi-archives",
			postResp: &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(strings.NewReader(`{"id": "jedi-archives"}`)),
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

			org, err := m.DeleteOrg(ctx, test.id)
			if test.err != nil && err.Error() != test.err.Error() {
				t.Errorf("error received: %v, expected: %v", err, test.err)
			}

			if org != nil && org.ID != test.id {
				t.Errorf("org id received: %s, expected: %s", org.ID, test.id)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	tests := []struct {
		desc     string
		input    *CreateUserInput
		postResp *http.Response
		postErr  error
		err      error
	}{
		{
			desc:     "error posting mutation to database",
			input:    &CreateUserInput{},
			postResp: nil,
			postErr:  errors.New("mock create org error"),
			err:      fmt.Errorf("%s: %w", errCreateUser, errors.New("mock create org error")),
		},
		{
			desc:  "error parsing database response body",
			input: &CreateUserInput{},
			postResp: &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(strings.NewReader("")),
			},
			postErr: nil,
			err:     fmt.Errorf("%s: %w", errUnmarshalJSON, errors.New("unexpected end of JSON input")),
		},
		{
			desc: "successful invocation",
			input: &CreateUserInput{
				FirstName: "jocasta",
				LastName:  "nu",
			},
			postResp: &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(strings.NewReader(`{"firstName": "jocasta","lastName":"nu"}`)),
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

			user, err := m.CreateUser(ctx, *test.input)
			if test.err != nil && err.Error() != test.err.Error() {
				t.Errorf("error received: %v, expected: %v", err, test.err)
			}

			if user != nil && user.FirstName != test.input.FirstName {
				t.Errorf("user name received: %s, expected: %s", user.FirstName, test.input.FirstName)
			}
		})
	}
}
