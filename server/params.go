package server

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

var (
	errValidateCSRFKey   = errors.New("server: param csrf key not set")
	errValidateDgraphURL = errors.New("server: param dgraph url not set")

	errParseReadConfigFile    = errors.New("server: error reading config file")
	errParseUnmarshal         = errors.New("server: error unmarshalling config file content")
	errParseReadPublicKeyFile = errors.New("server: error reading public key file")
)

type params struct {
	CSRFKey   string `json:"csrf_key"`
	DgraphURL string `json:"dgraph_url"`
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
	if p.CSRFKey == "" {
		return errValidateCSRFKey
	}
	if p.DgraphURL == "" {
		return errValidateDgraphURL
	}

	return nil
}
