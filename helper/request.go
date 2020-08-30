package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func sendDgraphRequest(dgraphURL, token, mutation string, input interface{}) error {
	variables := map[string]interface{}{
		"input": input,
	}

	payload := struct {
		Query     string
		Variables interface{}
	}{
		Query:     mutation,
		Variables: variables,
	}

	payloadByets, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", dgraphURL, bytes.NewBuffer(payloadByets))
	if err != nil {
		return err
	}
	req.Header.Set("X-Auth0-Token", token)
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("non-200 status received: %d", resp.StatusCode)
	}

	return nil
}
