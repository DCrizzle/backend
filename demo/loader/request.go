package loader

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

func (h *helper) sendRequest(mutation string, input interface{}) ([]string, error) {
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

	req, err := http.NewRequest("POST", h.dgraphURL, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Auth0-Token", h.token)

	resp, err := h.httpClient.Do(req)
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
