package main

import (
	"github.com/tidwall/gjson"

	"github.com/forstmeier/backend/graphql"
)

func (dc *dgraphClient) addProtocolForms(ownerID string, protocolIDs, protocolExternalIDs []string) ([]string, error) {
	protocolFormInputs := []map[string]interface{}{}
	owner := map[string]string{
		"id": ownerID,
	}

	for k, protocolID := range protocolIDs {
		protocol := map[string]string{
			"id": protocolID,
		}

		input := map[string]interface{}{
			"owner": owner,
			"title": randomString(titles),
			"body":  randomString(bodies),
			"protocols": []map[string]string{
				protocol,
			},
			"protocolID": protocolExternalIDs[k],
		}

		protocolFormInputs = append(protocolFormInputs, input)
	}

	data, err := dc.SendRequest(graphql.AddProtocolFormsMutation, protocolFormInputs)
	if err != nil {
		return nil, err
	}

	idValues := gjson.Get(data, "data.addProtocolForm.protocolForm.#.id").Array()
	ids := []string{}
	for _, id := range idValues {
		ids = append(ids, id.String())
	}

	return ids, nil
}
