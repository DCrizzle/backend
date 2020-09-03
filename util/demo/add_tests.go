package main

import (
	"github.com/tidwall/gjson"

	"github.com/forstmeier/backend/graphql"
)

func (dc *dgraphClient) addTest(ownerID, labID string, specimenIDs []string) (string, error) {
	owner := map[string]string{
		"id": ownerID,
	}
	lab := map[string]string{
		"id": labID,
	}

	specimens := make([]map[string]interface{}, len(specimenIDs))
	for i, specimenID := range specimenIDs {
		specimens[i] = map[string]interface{}{
			"id": specimenID,
		}
	}

	input := []map[string]interface{}{
		map[string]interface{}{
			"description": randomString(descriptions),
			"owner":       owner,
			"lab":         lab,
			"specimens":   specimens,
			"results":     []string{},
		},
	}

	data, err := dc.SendRequest(graphql.AddTestsMutation, input)
	if err != nil {
		return "", err
	}

	idValues := gjson.Get(data, "data.addTest.test.#.id").Array()
	ids := []string{}
	for _, id := range idValues {
		ids = append(ids, id.String())
	}

	return ids[0], nil
}
