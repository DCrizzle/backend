package server

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

var (
	errValidateAuth0APIAudience = errors.New("server: param auth0 api audience not set")
	errValidateAuth0Domain      = errors.New("server: param auth0 domain not set")
	errValidateCSRFKey          = errors.New("server: param csrf key not set")
	errValidateDgraphURL        = errors.New("server: param dgraph url not set")

	errParseReadConfigFile    = errors.New("server: error reading config file")
	errParseUnmarshal         = errors.New("server: error unmarshalling config file content")
	errParseReadPublicKeyFile = errors.New("server: error reading public key file")
)

type params struct {
	Auth0APIAudience string `json:"auth0_api_audience"`
	Auth0Domain      string `json:"auth0_domain"`
	CSRFKey          string `json:"csrf_key"`
	DgraphURL        string `json:"dgraph_url"`
}

func parseParams(configPath string) (*params, error) {
	configContent, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, errParseReadConfigFile
	}

	p := &params{}
	if err := json.Unmarshal(configContent, p); err != nil {
		return nil, errParseUnmarshal
	}

	return p, nil
}

func readPublicKey(publicKeyPath string) (string, error) {
	publicKeyContent, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return "", errParseReadPublicKeyFile
	}

	return string(publicKeyContent), err
}

func (p *params) validate() error {
	if p.Auth0APIAudience == "" {
		return errValidateAuth0APIAudience
	}
	if p.Auth0Domain == "" {
		return errValidateAuth0Domain
	}
	if p.CSRFKey == "" {
		return errValidateCSRFKey
	}
	if p.DgraphURL == "" {
		return errValidateDgraphURL
	}

	return nil
}
