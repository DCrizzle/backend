package demo

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type payload struct {
	Query     string      `json:"query"`
	Variables interface{} `json:"variables"`
}

type response struct {
	Data []data `json:"data"`
}

type data struct {
	ID string `json:"id"`
}

const dgraphURL = "localhost:8080/graphql"

func sendRequest(mutation string, input interface{}) ([]string, error) {
	variables := map[string]interface{}{
		"input": input,
	}

	p := payload{
		Query:     mutation,
		Variables: variables,
	}

	b, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(dgraphURL, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	ids := []string{}
	var respJSON response
	json.NewDecoder(resp.Body).Decode(&respJSON)
	for _, item := range respJSON.Data {
		ids = append(ids, item.ID)
	}

	return ids, nil
}
