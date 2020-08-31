package graphql

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type payload struct {
	Query     string      `json:"query"`
	Variables interface{} `json:"variables"`
}

// Dgraph is the client struct for communicating with the Dgraph database.
type Dgraph struct {
	client *http.Client
	url    string
	token  string
}

// New generates a pointer Dgraph client instance.
func New(client *http.Client, url, token string) *Dgraph {
	return &Dgraph{
		client: client,
		url:    url,
		token:  token,
	}
}

// SendRequest executes the provided query and variables against the
// Dgraph GraphQL database endpoint.
func (d *Dgraph) SendRequest(query string, input interface{}) (string, error) {
	variables := map[string]interface{}{
		"input": input,
	}

	payload := struct {
		Query     string      `json:"query"`
		Variables interface{} `json:"variables"`
	}{
		Query:     query,
		Variables: variables,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", d.url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}
	req.Header.Set("X-Auth0-Token", d.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := d.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
