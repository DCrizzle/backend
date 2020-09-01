package main

import (
	"time"

	"github.com/tidwall/gjson"

	"github.com/forstmeier/backend/graphql"
)

func (dc *dgraphClient) addConsent(ownerID, donorID, formID, protocolID string) (string, error) {
	owner := map[string]string{
		"id": ownerID,
	}
	donor := map[string]string{
		"id": donorID,
	}
	protocol := map[string]string{
		"id": protocolID,
	}
	form := map[string]string{
		"id": formID,
	}

	now := time.Now()

	input := []map[string]interface{}{
		map[string]interface{}{
			"owner":           owner,
			"donor":           donor,
			"protocol":        protocol,
			"form":            form,
			"consentedDate":   now.Format(time.RFC3339),
			"retentionPeriod": 360,
			"destructionDate": now.AddDate(0, 0, 360).Format(time.RFC3339),
		},
	}

	data, err := dc.SendRequest(graphql.AddConsentsMutation, input)
	if err != nil {
		return "", err
	}

	idValues := gjson.Get(data, "data.addConsent.consent.#.id").Array()
	ids := []string{}
	for _, id := range idValues {
		ids = append(ids, id.String())
	}

	return ids[0], nil
}
