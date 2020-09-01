package main

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/tidwall/gjson"

	"github.com/forstmeier/backend/graphql"
)

// NOTE: this could be improved to return the input specimenInputs variable
// which would be submitted in bulk by the calling scope
func (dc *dgraphClient) addBloodSpecimens(ownerID, donorID, consentID, protocolID string) ([]string, error) {
	owner := map[string]string{
		"id": ownerID,
	}
	donor := map[string]string{
		"id": donorID,
	}
	consent := map[string]string{
		"id": consentID,
	}
	protocol := map[string]string{
		"id": protocolID,
	}

	specimenCount := rand.Intn(10) + 1

	specimenInputs := []map[string]interface{}{}

	year := time.Now().Year()
	month := rand.Intn(12) + 1
	day := rand.Intn(25) + 1
	collectionDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC).String()

	status := randomString(graphql.Status)

	for specimenCount > 0 {
		input := map[string]interface{}{
			"externalID":     uuid.New().String(),
			"type":           graphql.SpecimenType[0],
			"collectionDate": collectionDate,
			"donor":          donor,
			"container":      graphql.Container[0],
			"status":         status,
			"description":    randomString(descriptions),
			"consent":        consent,
			"owner":          owner,
			"protocol":       protocol,
			"tests":          []string{},
			"bloodType":      randomString(graphql.BloodType),
			"volume":         1.0,
		}

		specimenInputs = append(specimenInputs, input)

		specimenCount--
	}

	data, err := dc.SendRequest(graphql.AddBloodSpecimensMutation, specimenInputs)
	if err != nil {
		return nil, err
	}

	idValues := gjson.Get(data, "data.addBloodSpecimen.bloodSpecimen.#.id").Array()
	ids := []string{}
	for _, id := range idValues {
		ids = append(ids, id.String())
	}

	return ids, nil
}
