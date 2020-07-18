package demo

import (
	"math/rand"
	"time"
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

func generateOwnerOrg(name, owner, orgType string) (string, error) {
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

	if orgType == "owner" {
		o = ownerOrg{
			org:      rootOrg,
			Labs:     []string{},
			Storages: []string{},
		}
	} else if orgType == "lab" {
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

func randomString(options []string) string {
	return options[rand.Intn(len(options))]
}

func randomInt(options []int) int {
	return options[rand.Intn(len(options))]
}

func dob() string {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	return t.Format(time.RFC3339)
}
