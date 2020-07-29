package demo

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func addOwnerOrgs() ([]string, error) {
	orgNames := []string{
		"owner_org_a",
		"owner_org_b",
		"owner_org_c",
	}

	inputs := []map[string]interface{}{}
	for i := 0; i < len(orgNames); i++ {
		input := map[string]interface{}{
			"street":    randomString(streets),
			"city":      randomString(cities),
			"county":    randomString(counties),
			"state":     randomString(states),
			"zip":       randomInt(zips),
			"name":      orgNames[i],
			"users":     []string{},
			"createdOn": "",
			"updatedOn": "",
		}
		inputs = append(inputs, input)
	}

	// outline:
	// [ ] create payload struct w/ populated fields
	// [ ] execute mutation
	// [ ] return ids / error values
}

// NOTE: reduce all input arguments to single owner IDs and call
// the generator iteratively over the input values
func addLabStorageOrgs(ownerIDs []string) (map[string]string, map[string]string, error) {
	labs := make(map[string]string)
	storages := make(map[string]string)

	// result := make(map[string]map[string]string)
	labNames := []string{
		"lab_org_a",
		"lab_org_b",
		"lab_org_c",
		"lab_org_d",
		"lab_org_e",
		"lab_org_f",
	}
	storageNames := []string{
		"storage_org_a",
		"storage_org_b",
	}

	for i, ownerID := range ownerIDs {
		labCount := rand.Intn(len(labNames))
		labInputs := []map[string]interface{}{}
		for labCount > 0 {
			labInput := map[string]interface{}{
				"street":    randomString(streets),
				"city":      randomString(cities),
				"county":    randomString(counties),
				"state":     randomString(states),
				"zip":       randomInt(zips),
				"name":      labNames[i],
				"users":     []string{},
				"createdOn": "",
				"updatedOn": "",
				"owner":     ownerID,
				"specimens": []string{},
				"plans":     []string{},
			}

			labInputs = append(labInputs, labInput)

			// outline:
			// [ ] create payload struct w/ populated fields
			// [ ] execute mutation
			// [ ] store ids in result map with key "owner id"

			labCount--
		}

		storageName := storageNames[rand.Intn(len(storageNames))]
		storageInput := map[string]interface{}{
			"street":    randomString(streets),
			"city":      randomString(cities),
			"county":    randomString(counties),
			"state":     randomString(states),
			"zip":       randomInt(zips),
			"name":      storageNames[i],
			"users":     []string{},
			"createdOn": "",
			"updatedOn": "",
			"owner":     ownerID,
			"specimens": []string{},
			"plans":     []string{},
		}

		// outline:
		// [ ] create payload struct w/ populated fields
		// [ ] execute mutation
		// [ ] store ids in result map with key "owner id"
	}

	return labs, storages, nil
}

func addUsers(ownerIDs, labIDs, storageIDs []string) ([]string, error) {
	inputs := []map[string]interface{}{}
	for i, user := range users {
		ownerID := ownerIDs[i%3]
		orgID := ""
		if user.role == "USER_STORAGE" {
			orgID = randomString(storageIDs)
		} else if user.role == "USER_LAB" {
			orgID = randomString(labIDs)
		} else {
			orgID = ownerID
		}

		input := map[string]interface{}{
			"owner":     ownerID,
			"email":     user.email,
			"firstName": user.first,
			"lastName":  user.last,
			"role":      user.role,
			"org":       orgID,
		}

		inputs = append(inputs, input)
	}

	// outline:
	// [ ] create payload struct w/ populated fields
	// [ ] execute mutation
	// [ ] return ids / error values
}

func addProtocolsFormsPlans(ownerIDs, labIDs, storageIDs []string) ([]string, []string, []string, error) {
	for i, ownerID := range ownerIDs {
		dobStart := time.Date(1977, time.May, 25, 22, 0, 0, 0, time.UTC)
		dobEnd := time.Date(2005, time.May, 19, 22, 0, 0, 0, time.UTC)

		protocolInputs := []map[string]interface{}{}
		for _, protocolName := range protocolNames {
			ageStart := randomInt(ages)
			ageEnd := ageStart + 20

			input := map[string]interface{}{
				"street":      randomString(streets),
				"city":        randomString(cities),
				"county":      randomString(counties),
				"state":       randomString(states),
				"zip":         randomInt(zips),
				"owner":       ownerID,
				"name":        protocolName,
				"description": randomString(descriptions),
				"form":        "",
				"plan":        "",
				"dobStart":    dobStart.String(),
				"dobEnd":      dobEnd.String(),
				"race":        randomString(races),
				"sex":         randomString(sexes),
				"specimens":   "",
			}

			protocolInputs = append(protocolInputs, input)
		}

		// outline:
		// [ ] create payload struct w/ populated fields
		// [ ] execute mutation
		// [ ] store ids in result map with key "owner id"

		protocolIDs := []string{"id_A", "id_B", "id_C"}
		protocolExternalIDs := []string{
			uuid.New().String(),
			uuid.New().String(),
			uuid.New().String(),
		}

		protocolFormInput := []map[string]interface{}{}
		for k, protocolID := range protocolIDs {
			input := map[string]interface{}{
				"owner":      ownerID,
				"title":      randomString(titles),
				"body":       randomString(bodies),
				"protocol":   protocolID,
				"protocolID": protocolExternalIDs[k],
			}

			protocolFormInput = append(protocolFormInput, input)
		}

		// outline:
		// [ ] create payload struct w/ populated fields
		// [ ] execute mutation
		// [ ] store ids in result map with key "owner id"

		planInputs := []map[string]interface{}{}
		for j, planName := range planNames {
			input := map[string]interface{}{
				"owner":    ownerID,
				"name":     planName,
				"labs":     randomString(labIDs),
				"storages": randomString(storageIDs),
				"protocol": protocolIDs[j],
			}

			planInputs = append(planInputs, input)
		}

		// outline:
		// [ ] create payload struct w/ populated fields
		// [ ] execute mutation
		// [ ] store ids in result map with key "owner id"
	}
}

func addConsentForms(ownerIDs []string) ([]string, error) {
	consentFormInput := []map[string]interface{}{}
	for _, ownerID := range ownerIDs {
		input := map[string]interface{}{
			"owner": ownerID,
			"title": randomString(titles),
			"body":  randomString(bodies),
		}

		consentFormInput = append(consentFormInput, input)
	}

	// outline:
	// [ ] create payload struct w/ populated fields
	// [ ] execute mutation
	// [ ] store ids in result map with key "owner id"
}

func addDonor(ownerIDs []string) ([]string, error) {
	for _, ownerID := range ownerIDs {
		donorCount := rand.Intn(100) + 50
		donorInputs := []map[string]interface{}{}
		for donorCount > 0 {
			dob, age := randomDOBAndAge()

			input := map[string]interface{}{
				"street":    randomString(streets),
				"city":      randomString(cities),
				"county":    randomString(counties),
				"state":     randomString(states),
				"zip":       randomInt(zips),
				"owner":     ownerID,
				"dob":       dob,
				"age":       age,
				"sex":       randomString(sexes),
				"race":      randomString(races),
				"specimens": []string{},
				"consents":  []string{},
			}

			donorInputs = append(donorInputs, input)
			donorCount--
		}

		// outline:
		// [ ] create payload struct w/ populated fields
		// [ ] execute mutation
		// [ ] store ids in result map with key "owner id"
	}
}

func addConsent(ownerID, donorID, formID, protocolID string) (string, error) {
	now := time.Now()
	input := map[string]interface{}{
		"owner":           ownerID,
		"donor":           donorID,
		"specimen":        "",
		"protocol":        protocolID,
		"form":            formID,
		"consentedDate":   now.String(),
		"retentionPeriod": 360,
		"destructionDate": now.AddDate(0, 0, 360).String(),
	}

	// outline:
	// [ ] create payload struct w/ populated fields
	// [ ] execute mutation
	// [ ] store ids in result map with key "owner id"
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

func randomDOBAndAge() (string, int) {
	currentYear := time.Now().Year()

	yo := rand.Intn(50) + 20

	year := currentYear - yo
	month := rand.Intn(12) + 1
	day := rand.Intn(25) + 1

	dob := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC).String()

	return dob, yo
}

// outline:
// [x] add owner org
// - [x] input:
// - - [x] name string
// - - [x] created/updated on time
// - [x] output: owner org id
// [x] add lab / storage org
// - [x] input:
// - - [x] name string
// - - [x] created/updated on time
// - - [x] owner org id
// - [x] output:
// - - [x] lab / storage org id
// [x] add user
// - [x] input:
// - - [x] owner id
// - - [x] (various value / enum inputs)
// - [x] output:
// - - [x] user id
// [x] add plan
// - [x] input:
// - - [x] name string
// - - [x] owner / lab storage org ids
// - [x] output:
// - - [x] plan id
// [x] add protocol
// - [x] input:
// - - [x] owner id
// - - [x] (various value / enum inputs)
// - [x] output:
// - - [x] protocol id
// [x] add consent form
// - [x] input:
// - - [x] owner id
// - - [x] title / body string
// - [x] output:
// - - [x] consent form id
// [x] add protocol form
// - [x] input:
// - - [x] protocol id string (generated)
// - - [x] protocol ids string array
// - - [x] owner id
// - - [x] title / body string
// - [x] output:
// - - [x] protocol form id
// [x] add donor
// - [x] input:
// - - [x] owner id
// - - [x] (various value / enum inputs)
// - - [x] (not including consents / specimens)
// - [x] output:
// - - [x] donor id
// [ ] add consent
// - [ ] input:
// - - [ ] owner id
// - - [ ] donor id
// - - [ ] consent form id
// - - [ ] protocol id (non-generated)
// - - [ ] (various value / enum inputs)
// - - [ ] (not including specimens)
// - [ ] output:
// - - [ ] consent id
// [ ] add blood specimen
// - [ ] input:
// - - [ ] donor id
// - - [ ] consent id
// - - [ ] owner / lab / storage id
// - - [ ] protocol id
// - - [ ] (various value / enum inputs)
// - [ ] output:
// - - [ ] blood specimen id
// [ ] add test
// - [ ] input:
// - - [ ] owner id
// - - [ ] lab id
// - - [ ] specimens id
// - [ ] output:
// - - [ ] test id
// [ ] add result
// - [ ] input:
// - - [ ] owner id
// - - [ ] notes string
// - - [ ] test id
// - [ ] output:
// - - [ ] result id
