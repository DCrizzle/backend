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

func sendMutation(input payload) ([]string, error) {
	b, err := json.Marshal(input)
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
