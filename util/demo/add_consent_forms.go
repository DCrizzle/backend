package main

import (
	"github.com/tidwall/gjson"

	"github.com/forstmeier/backend/graphql"
)

func (dc *dgraphClient) addConsentForms(ownerID string, count int) ([]string, error) {
	consentFormInputs := []map[string]interface{}{}
	owner := map[string]string{
		"id": ownerID,
	}

	for count > 0 {
		input := map[string]interface{}{
			"owner": owner,
			"title": randomString(titles),
			"body":  randomString(bodies),
		}

		consentFormInputs = append(consentFormInputs, input)
		count--
	}

	data, err := dc.SendRequest(graphql.AddConsentFormsMutation, consentFormInputs)
	if err != nil {
		return nil, err
	}

	idValues := gjson.Get(data, "data.addConsentForm.consentForm.#.id").Array()
	ids := []string{}
	for _, id := range idValues {
		ids = append(ids, id.String())
	}

	return ids, nil
}
