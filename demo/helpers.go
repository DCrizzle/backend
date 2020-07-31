package demo

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func addOwnerOrgs() ([]string, error) {
	inputs := []map[string]interface{}{}
	for i := 0; i < len(orgs); i++ {
		input := map[string]interface{}{
			"street":    randomString(streets),
			"city":      randomString(cities),
			"county":    randomString(counties),
			"state":     randomString(states),
			"zip":       randomInt(zips),
			"name":      orgs[i],
			"users":     []string{},
			"createdOn": "",
			"updatedOn": "",
		}
		inputs = append(inputs, input)
	}

	return sendRequest(addOwnerOrgsMutation, inputs)
}

func addLabStorageOrgs(ownerID string) ([]string, []string, error) {
	labInputs := []map[string]interface{}{}
	labCount := rand.Intn(len(labs))
	for labCount > 0 {
		labInput := map[string]interface{}{
			"street":    randomString(streets),
			"city":      randomString(cities),
			"county":    randomString(counties),
			"state":     randomString(states),
			"zip":       randomInt(zips),
			"name":      labs[labCount-1],
			"users":     []string{},
			"createdOn": "",
			"updatedOn": "",
			"owner":     ownerID,
			"specimens": []string{},
			"plans":     []string{},
		}

		labInputs = append(labInputs, labInput)
		labCount--
	}

	labs, err := sendRequest(addLabOrgsMutation, labInputs)
	if err != nil {
		return nil, nil, err
	}

	storageIndex := rand.Intn(len(storages))
	storageName := storages[storageIndex]
	storageInput := map[string]interface{}{
		"street":    randomString(streets),
		"city":      randomString(cities),
		"county":    randomString(counties),
		"state":     randomString(states),
		"zip":       randomInt(zips),
		"name":      storages[storageIndex],
		"users":     []string{},
		"createdOn": "",
		"updatedOn": "",
		"owner":     ownerID,
		"specimens": []string{},
		"plans":     []string{},
	}

	storages, err := sendRequest(addStorageOrgsMutation, storageInput)
	if err != nil {
		return nil, nil, err
	}

	return labs, storages, nil
}

func addUsers(ownerID string, labIDs, storageIDs []string) ([]string, error) {
	inputs := []map[string]interface{}{}
	for i, user := range users {
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

	return sendRequest(addUsersMutation, inputs)
}

func addProtocolsAndPlans(ownerID string, labIDs, storageIDs []string) ([]string, []string, error) {
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
			"race":        randomString(RACE),
			"sex":         randomString(SEX),
			"specimens":   "",
		}

		protocolInputs = append(protocolInputs, input)
	}

	protocolIDs, err := sendRequest(addProtocolsMutation, protocolInputs)
	if err != nil {
		return nil, nil, err
	}

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

	planIDs, err := sendRequest(addPlansMutation, planInputs)
	if err != nil {
		return nil, nil, err
	}

	return protocolIDs, planIDs, nil
}

func addProtocolForms(ownerID string, protocolIDs, protocolExternalIDs []string) ([]string, error) {
	protocolFormInputs := []map[string]interface{}{}
	for k, protocolID := range protocolIDs {
		input := map[string]interface{}{
			"owner":      ownerID,
			"title":      randomString(titles),
			"body":       randomString(bodies),
			"protocol":   protocolID,
			"protocolID": protocolExternalIDs[k],
		}

		protocolFormInputs = append(protocolFormInputs, input)
	}

	return sendRequest(addProtocolFormsMutation, protocolFormInputs)
}

func addConsentForms(ownerID string, count int) ([]string, error) {
	consentFormInputs := []map[string]interface{}{}
	for count > 0 {
		input := map[string]interface{}{
			"owner": ownerID,
			"title": randomString(titles),
			"body":  randomString(bodies),
		}

		consentFormInputs = append(consentFormInputs, input)
		count--
	}

	return sendRequest(addConsentFormsMutation, consentFormInputs)
}

func addDonor(ownerID string) ([]string, error) {
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
			"sex":       randomString(SEX),
			"race":      randomString(RACE),
			"specimens": []string{},
			"consents":  []string{},
		}

		donorInputs = append(donorInputs, input)
		donorCount--
	}

	return sendRequest(addDonorMutation, donorInputs)
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

func addBloodSpecimens(ownerID, donorID, consentID, protocolID string) ([]string, error) {
	specimenCount := rand.Intn(10) + 1

	specimenInputs := []map[string]interface{}{}

	year := time.Now().Year()
	month := rand.Intn(12) + 1
	day := rand.Intn(25) + 1
	collectionDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC).String()

	status := randomString(STATUS)
	destructionDate := ""
	if status == "DESTROYED" {
		destructionDate = time.Now().String()
	}

	for specimenCount > 0 {
		input := map[string]interface{}{
			"externalID":      uuid.New().String(),
			"type":            SPECIMEN_TYPE[0],
			"collectionDate":  collectionDate,
			"donor":           donorID,
			"container":       CONTAINER[0],
			"status":          status,
			"destructionDate": destructionDate,
			"description":     randomString(descriptions),
			"consent":         consentID,
			"owner":           ownerID,
			"lab":             "", // NOTE: add later (?)
			"storage":         "", // NOTE: add later (?)
			"protocol":        protocolID,
			"tests":           []string{},
			"bloodType":       randomString(BLOOD_TYPE),
			"volume":          1.0,
		}

		specimenInputs = append(specimenInputs, input)

		specimenCount--
	}

	// outline:
	// [ ] create payload struct w/ populated fields
	// [ ] execute mutation
	// [ ] store ids in result map with key "owner id"
}

func addTest(ownerID, labID string, specimenIDs []string) ([]string, error) {
	input := map[string]interface{}{
		"description": randomString(descriptions),
		"owner":       ownerID,
		"lab":         labID,
		"specimens":   specimenIDs,
		"results":     []string{},
	}

	// outline:
	// [ ] create payload struct w/ populated fields
	// [ ] execute mutation
	// [ ] store ids in result map with key "owner id"
}

func addResult(ownerID, testID string) ([]string, error) {
	input := map[string]interface{}{
		"owner": ownerID,
		"notes": randomString(notes),
		"test":  testID,
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
