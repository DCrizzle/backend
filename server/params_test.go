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
			err:         errParseReadConfigFile,
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

			if err == nil && p.CSRFKey != "key" {
				t.Errorf("params received: %+v", p)
			}
		})
	}
}

func Test_readPublicKey(t *testing.T) {
	tests := []struct {
		description string
		content     []byte
		output      string
		err         error
	}{
		{
			description: "no file at provided path",
			content:     nil,
			output:      "",
			err:         errParseReadPublicKeyFile,
		},
		{
			description: "successful config file parse to params",
			content:     []byte("test-public-key"),
			output:      "test-public-key",
			err:         nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			fileName := "readPublicKey*.pem"
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

			keyString, err := readPublicKey(fileName)
			if err != test.err {
				t.Errorf("error received: %v, expected: %v", err, test.err)
			}

			if keyString != test.output {
				t.Errorf("key received: %s", keyString)
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
			description: "auth0 api audience key not set",
			params:      params{},
			err:         errValidateAuth0APIAudience,
		},
		{
			description: "auth0 domain key not set",
			params: params{
				Auth0APIAudience: "audience",
			},
			err: errValidateAuth0Domain,
		},
		{
			description: "csrf key not set",
			params: params{
				Auth0APIAudience: "audience",
				Auth0Domain:      "domain",
			},
			err: errValidateCSRFKey,
		},
		{
			description: "dgraph url not set",
			params: params{
				Auth0APIAudience: "audience",
				Auth0Domain:      "domain",
				CSRFKey:          "key",
			},
			err: errValidateDgraphURL,
		},
		{
			description: "all fields set",
			params: params{
				Auth0APIAudience: "audience",
				Auth0Domain:      "domain",
				CSRFKey:          "key",
				DgraphURL:        "url",
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
