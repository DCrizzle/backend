package main

import (
	"time"

	"github.com/tidwall/gjson"

	"github.com/forstmeier/backend/graphql"
)

func (dc *dgraphClient) addOwnerOrgs() ([]string, error) {
	inputs := []map[string]interface{}{}
	for i := 0; i < len(orgs); i++ {
		now := time.Now().Format(time.RFC3339)
		input := map[string]interface{}{
			"street":    randomString(streets),
			"city":      randomString(cities),
			"county":    randomString(counties),
			"state":     randomString(states),
			"zip":       randomInt(zips),
			"country":   randomString(countries),
			"name":      orgs[i],
			"users":     []string{},
			"createdOn": now,
			"updatedOn": now,
			"labs":      []string{},
			"storages":  []string{},
		}
		inputs = append(inputs, input)
	}

	data, err := dc.SendRequest(graphql.AddOwnerOrgsMutation, inputs)
	if err != nil {
		return nil, err
	}

	idValues := gjson.Get(data, "data.addOwnerOrg.ownerOrg.#.id").Array()
	ids := []string{}
	for _, id := range idValues {
		ids = append(ids, id.String())
	}

	return ids, nil
}
