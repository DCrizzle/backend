package demo

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

const dgraphURL = ""

func generateDonors(count int, owner string) ([]string, error) {
	donors := []donor{}
	for i := 0; i < count; i++ {
		d := donor{
			Street:    randomString(streets),
			City:      randomString(cities),
			County:    randomString(counties),
			State:     randomString(states),
			ZIP:       randomInt(zips),
			Owner:     owner,
			Age:       randomInt(ages),
			DOB:       dob(),
			Sex:       randomString(SEX),
			Race:      randomString(RACE),
			Specimens: []string{},
			Consents:  []string{},
		}

		donors = append(donors, d)
	}

	variables := map[string][]donor{
		"input": donors,
	}

	input := payload{
		Query:     "",
		Variables: variables,
	}

	return sendMutation(input)
}

func generateOwnerOrg(name string) (string, error) {
	now := time.Now().Format(time.RFC3339)
	ownerOrg := ownerOrg{
		org: org{
			Street:    randomString(streets),
			City:      randomString(cities),
			County:    randomString(counties),
			State:     randomString(states),
			ZIP:       randomInt(zips),
			Name:      name,
			Users:     []string{},
			CreatedOn: now,
			UpdatedOn: now,
		},
		Labs:     []string{},
		Storages: []string{},
	}

	variables := map[string]interface{}{
		"input": ownerOrg,
	}

	input := payload{
		Query:     "",
		Variables: variables,
	}

	id, err := sendMutation(input)
	if err != nil {
		return "", err
	}

	return id[0], nil
}

func generateLabStorageOrgs(name, owner, orgType string) (string, error) {
	now := time.Now().Format(time.RFC3339)
	var o interface{}
	rootOrg := org{
		Street:    randomString(streets),
		City:      randomString(cities),
		County:    randomString(counties),
		State:     randomString(states),
		ZIP:       randomInt(zips),
		Name:      name,
		Users:     []string{},
		CreatedOn: now,
		UpdatedOn: now,
	}

	if orgType == "lab" {
		o = labOrg{
			org:       rootOrg,
			Owner:     owner,
			Specimens: []string{},
			Plans:     []string{},
		}
	} else if orgType == "storage" {
		o = storageOrg{
			org:       rootOrg,
			Owner:     owner,
			Specimens: []string{},
			Plans:     []string{},
		}
	}

	variables := map[string]interface{}{
		"input": o,
	}

	input := payload{
		Query:     "",
		Variables: variables,
	}

	id, err := sendMutation(input)
	if err != nil {
		return "", err
	}

	return id[0], nil
}

func generateProtocolConsentForm(owner, formType string) (string, error) {
	var form interface{}
	if formType == "protocol" {
		form = protocolForm{
			Owner:      owner,
			Title:      randomString(titles),
			Body:       randomString(bodies),
			ProtocolID: id(),
			Protocols:  []string{},
		}
	} else if formType == "consent" {
		form = consentForm{
			Owner:    owner,
			Title:    randomString(titles),
			Body:     randomString(bodies),
			Consents: []string{},
		}
	}

	variables := map[string]interface{}{
		"input": form,
	}

	input := payload{
		Query:     "",
		Variables: variables,
	}

	id, err := sendMutation(input)
	if err != nil {
		return "", err
	}

	return id[0], nil
}

func generateProtocol(owner, protocolID string) (string, error) {
	protocol := protocol{
		Street:      randomString(streets),
		City:        randomString(cities),
		County:      randomString(counties),
		State:       randomString(states),
		ZIP:         randomInt(zips),
		Owner:       owner,
		Name:        randomString(protocolNames),
		Description: randomString(descriptions),
		Form:        protocolID,
		Plan:        "non_id",
		Ages:        randomInts(5, ages),
		AgeStart:    25,
		AgeEnd:      45,
		DOBs:        []string{},
		DOBStart:    "",
		DOBEnd:      "",
		Race:        randomString(RACE),
		Sex:         randomString(SEX),
		Specimens:   []string{},
	}

	variables := map[string]interface{}{
		"input": protocol,
	}

	input := payload{
		Query:     "",
		Variables: variables,
	}

	id, err := sendMutation(input)
	if err != nil {
		return "", err
	}

	return id[0], nil
}

func generateConsent(owner, donor, protocol string) (string, error) {
	c := consent{
		Owner: owner,
		Donor: donor,
		Specimen: "",
		Protocol: protocol,
		ConsentDate: "",
		RetentionPeriod: 0,
		DestructionDate: "",
	}

	variables := map[string]interface{}{
		"input": c,
	}

	input := payload{
		Query:     "",
		Variables: variables,
	}

	id, err := sendMutation(input)
	if err != nil {
		return "", err
	}

	return id[0], nil
}

func randomString(options []string) string {
	return options[rand.Intn(len(options))]
}

func randomInt(options []int) int {
	return options[rand.Intn(len(options))]
}

func randomInts(count int, options []int) []int {
	ints := []int{}
	for i := 0; i < count; i++ {
		ints = append(ints, options[rand.Intn(len(options))])
	}
	return ints
}

func dob() string {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	return t.Format(time.RFC3339)
}

func id() string {
	return uuid.New().String()
}
