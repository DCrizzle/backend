package server

import (
	"io/ioutil"
	"os"
	"testing"
)

func Test_parseParams(t *testing.T) {
	tests := []struct {
		description string
		content     []byte
		err         error
	}{
		{
			description: "no file at provided path",
			content:     nil,
			err:         errParseReadFile,
		},
		{
			description: "non-json content in config file",
			content:     []byte(`-------`),
			err:         errParseUnmarshal,
		},
		{
			description: "successful config file parse to params",
			content: []byte(`
				{
					"auth0_api_audience":"audience",
					"auth0_api_client_secret": "secret",
					"auth0_domain": "domain",
					"csrf_key": "key",
					"dgraph_url": "url"
				}
			`),
			err: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			fileName := "parseParamsConfig*.json"
			if test.content != nil {

				configFile, err := ioutil.TempFile("", fileName)
				if err != nil {
					t.Fatal(err)
				}
				defer os.Remove(configFile.Name())

				fileName = configFile.Name()

				_, err = configFile.Write(test.content)
				if err != nil {
					t.Fatal(err)
				}
			}

			p, err := parseParams(fileName)
			if err != test.err {
				t.Errorf("error received: %v, expected: %v", err, test.err)
			}

			if err == nil && p.Auth0APIAudience != "audience" {
				t.Errorf("params received: %+v", p)
			}
		})
	}
}

func Test_validate(t *testing.T) {
	tests := []struct {
		description string
		params      params
		err         error
	}{
		{
			description: "auth0 api audience not set",
			params:      params{},
			err:         errValidateAuth0APIAudience,
		},
		{
			description: "auth0 api client secret not set",
			params: params{
				Auth0APIAudience: "audience",
			},
			err: errValidateAuth0APIClientSecret,
		},
		{
			description: "auth0 domain not set",
			params: params{
				Auth0APIAudience:     "audience",
				Auth0APIClientSecret: "secret",
			},
			err: errValidateAuth0Domain,
		},
		{
			description: "csrf key not set",
			params: params{
				Auth0APIAudience:     "audience",
				Auth0APIClientSecret: "secret",
				Auth0Domain:          "domain",
			},
			err: errValidateCSRFKey,
		},
		{
			description: "dgraph url not set",
			params: params{
				Auth0APIAudience:     "audience",
				Auth0APIClientSecret: "secret",
				Auth0Domain:          "domain",
				CSRFKey:              "key",
			},
			err: errValidateDgraphURL,
		},
		{
			description: "all fields set",
			params: params{
				Auth0APIAudience:     "audience",
				Auth0APIClientSecret: "secret",
				Auth0Domain:          "domain",
				CSRFKey:              "key",
				DgraphURL:            "url",
			},
			err: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			err := test.params.validate()
			if err != test.err {
				t.Errorf("error received: %v, expected: %v", err, test.err)
			}
		})
	}
}
