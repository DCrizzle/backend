package main

import (
	"github.com/tidwall/gjson"

	"github.com/forstmeier/backend/graphql"
)

func (dc *dgraphClient) addResult(ownerID, testID string) (string, error) {
	owner := map[string]string{
		"id": ownerID,
	}
	test := map[string]string{
		"id": testID,
	}

	input := []map[string]interface{}{
		map[string]interface{}{
			"owner": owner,
			"notes": randomString(notes),
			"test":  test,
		},
	}

	data, err := dc.SendRequest(graphql.AddResultsMutation, input)
	if err != nil {
		return "", err
	}

	idValues := gjson.Get(data, "data.addResult.result.#.id").Array()
	ids := []string{}
	for _, id := range idValues {
		ids = append(ids, id.String())
	}

	return ids[0], nil
}
