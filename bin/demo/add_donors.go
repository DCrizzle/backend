package main

import (
	"math/rand"

	"github.com/tidwall/gjson"

	"github.com/forstmeier/backend/graphql"
)

func (dc *dgraphClient) addDonor(ownerID string) ([]string, error) {
	donorCount := rand.Intn(100) + 50
	donorInputs := []map[string]interface{}{}
	owner := map[string]string{
		"id": ownerID,
	}
	for donorCount > 0 {
		dob, age := randomDOBAndAge()

		input := map[string]interface{}{
			"street":    randomString(streets),
			"city":      randomString(cities),
			"county":    randomString(counties),
			"state":     randomString(states),
			"zip":       randomInt(zips),
			"country":   randomString(countries),
			"owner":     owner,
			"dob":       dob,
			"age":       age,
			"sex":       randomString(graphql.Sex),
			"race":      randomString(graphql.Race),
			"specimens": []string{},
			"consents":  []string{},
		}

		donorInputs = append(donorInputs, input)
		donorCount--
	}

	data, err := dc.SendRequest(graphql.AddDonorsMutation, donorInputs)
	if err != nil {
		return nil, err
	}

	idValues := gjson.Get(data, "data.addDonor.donor.#.id").Array()
	ids := []string{}
	for _, id := range idValues {
		ids = append(ids, id.String())
	}

	return ids, nil
}
