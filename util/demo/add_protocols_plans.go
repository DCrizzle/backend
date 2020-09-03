package main

import (
	"errors"
	"time"

	"github.com/tidwall/gjson"

	"github.com/forstmeier/backend/graphql"
)

func (dc *dgraphClient) addProtocolsAndPlans(ownerID string, labIDs, storageIDs []string) ([]string, []string, error) {
	dobStart := time.Date(1977, time.May, 25, 22, 0, 0, 0, time.UTC)
	dobEnd := time.Date(2005, time.May, 19, 22, 0, 0, 0, time.UTC)
	owner := map[string]string{
		"id": ownerID,
	}

	emptyForm := map[string]interface{}{
		"owner":      owner,
		"title":      "",
		"body":       "",
		"protocolID": "",
	}
	emptyPlan := map[string]interface{}{
		"owner":    owner,
		"name":     "",
		"labs":     []string{},
		"storages": []string{},
	}

	protocolInputs := []map[string]interface{}{}
	for _, protocolName := range protocolNames {
		input := map[string]interface{}{
			"street":      randomString(streets),
			"city":        randomString(cities),
			"county":      randomString(counties),
			"state":       randomString(states),
			"zip":         randomInt(zips),
			"country":     randomString(countries),
			"owner":       owner,
			"name":        protocolName,
			"description": randomString(descriptions),
			"form":        emptyForm,
			"plan":        emptyPlan,
			"dobStart":    dobStart.String(),
			"dobEnd":      dobEnd.String(),
			"race": []string{
				randomString(graphql.Race),
			},
			"sex": []string{
				randomString(graphql.Sex),
			},
			"specimens": make([]map[string]string, 0),
		}

		protocolInputs = append(protocolInputs, input)
	}

	protocols, err := dc.SendRequest(graphql.AddProtocolsMutation, protocolInputs)
	if err != nil {
		return nil, nil, err
	}

	protocolsIDValues := gjson.Get(protocols, "data.addProtocol.protocol.#.id").Array()
	protocolIDs := []string{}
	for _, id := range protocolsIDValues {
		protocolIDs = append(protocolIDs, id.String())
	}
	if len(protocolIDs) == 0 {
		return nil, nil, errors.New("no protocol ids returned from add mutation")
	}

	planInputs := []map[string]interface{}{}
	for j, planName := range planNames {
		lab := map[string]string{
			"id": randomString(labIDs),
		}
		storage := map[string]string{
			"id": randomString(storageIDs),
		}
		protocol := map[string]string{
			"id": protocolIDs[j],
		}

		input := map[string]interface{}{
			"owner": owner,
			"name":  planName,
			"labs": []map[string]string{
				lab,
			},
			"storages": []map[string]string{
				storage,
			},
			"protocol": protocol,
		}

		planInputs = append(planInputs, input)
	}

	plans, err := dc.SendRequest(graphql.AddPlansMutation, planInputs)
	if err != nil {
		return nil, nil, err
	}

	plansIDValues := gjson.Get(plans, "data.addPlan.plan.#.id").Array()
	planIDs := []string{}
	for _, id := range plansIDValues {
		planIDs = append(planIDs, id.String())
	}

	return protocolIDs, planIDs, nil
}
