package main

import (
	"math/rand"
	"time"

	"github.com/tidwall/gjson"

	"github.com/forstmeier/backend/graphql"
)

func (dc *dgraphClient) addLabAndStorageOrgs(ownerID string) ([]string, []string, error) {
	labInputs := []map[string]interface{}{}
	labCount := rand.Intn(len(labs)) + 1

	owner := map[string]string{
		"id": ownerID,
	}
	now := time.Now().Format(time.RFC3339)

	for labCount > 0 {
		labInput := map[string]interface{}{
			"street":    randomString(streets),
			"city":      randomString(cities),
			"county":    randomString(counties),
			"state":     randomString(states),
			"zip":       randomInt(zips),
			"country":   randomString(countries),
			"name":      labs[labCount-1],
			"users":     []string{},
			"createdOn": now,
			"updatedOn": now,
			"owner":     owner,
			"specimens": []string{},
			"plans":     []string{},
		}

		labInputs = append(labInputs, labInput)
		labCount--
	}

	labs, err := dc.SendRequest(graphql.AddLabOrgsMutation, labInputs)
	if err != nil {
		return nil, nil, err
	}

	labIDValues := gjson.Get(labs, "data.addLabOrg.labOrg.#.id").Array()
	labIDs := []string{}
	for _, id := range labIDValues {
		labIDs = append(labIDs, id.String())
	}

	storageIndex := rand.Intn(len(storages))
	storageInput := []map[string]interface{}{
		map[string]interface{}{
			"street":    randomString(streets),
			"city":      randomString(cities),
			"county":    randomString(counties),
			"state":     randomString(states),
			"country":   randomString(countries),
			"zip":       randomInt(zips),
			"name":      storages[storageIndex],
			"users":     []string{},
			"createdOn": now,
			"updatedOn": now,
			"owner":     owner,
			"specimens": []string{},
			"plans":     []string{},
		},
	}

	storages, err := dc.SendRequest(graphql.AddStorageOrgsMutation, storageInput)
	if err != nil {
		return nil, nil, err
	}

	storageIDValues := gjson.Get(storages, "data.addStorageOrg.storageOrg.#.id").Array()
	storageIDs := []string{}
	for _, id := range storageIDValues {
		storageIDs = append(storageIDs, id.String())
	}

	return labIDs, storageIDs, nil
}
