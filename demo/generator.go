package demo

import (
	"math/rand"
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

func addLabStorageOrgs(ownerIDs []string) (map[string]map[string][]string, error) {
	result := make(map[string]map[string]string)
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

	for _, ownerID := range ownerIDs {
		labCount := rand.Intn(len(labNames))
		labInputs := []map[string]interface{}{}
		for labCount > 0 {
			labInput := map[string]interface{}{
				"street":    randomString(streets),
				"city":      randomString(cities),
				"county":    randomString(counties),
				"state":     randomString(states),
				"zip":       randomInt(zips),
				"name":      orgNames[i],
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
			// [ ] store ids in result map with key "labs"
		}

		storageName := storageNames[rand.Intn(len(storageNames))]
		storageInput := map[string]interface{}{
			"street":    randomString(streets),
			"city":      randomString(cities),
			"county":    randomString(counties),
			"state":     randomString(states),
			"zip":       randomInt(zips),
			"name":      orgNames[i],
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
		// [ ] store ids in result map with key "storages"
	}

	return result, nil
}

func randomString(options []string) string {
	return options[rand.Intn(len(options))]
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
// [ ] add user
// - [ ] input:
// - - [ ] owner id
// - - [ ] (various value / enum inputs)
// - [ ] output:
// - - [ ] user id
// [ ] add plan
// - [ ] input:
// - - [ ] name string
// - - [ ] owner / lab storage org ids
// - [ ] output:
// - - [ ] plan id
// [ ] add protocol
// - [ ] input:
// - - [ ] owner id
// - - [ ] (various value / enum inputs)
// - [ ] output:
// - - [ ] protocol id
// [ ] add consent form
// - [ ] input:
// - - [ ] owner id
// - - [ ] title / body string
// - [ ] output:
// - - [ ] consent form id
// [ ] add protocol form
// - [ ] input:
// - - [ ] protocol id string (generated)
// - - [ ] protocol ids string array
// - - [ ] owner id
// - - [ ] title / body string
// - [ ] output:
// - - [ ] protocol form id
// [ ] add donor
// - [ ] input:
// - - [ ] owner id
// - - [ ] (various value / enum inputs)
// - - [ ] (not including consents / specimens)
// - [ ] output:
// - - [ ] donor id
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

// func randomInt(options []int) int {
// 	return options[rand.Intn(len(options))]
// }
//
// func randomInts(count int, options []int) []int {
// 	ints := []int{}
// 	for i := 0; i < count; i++ {
// 		ints = append(ints, options[rand.Intn(len(options))])
// 	}
// 	return ints
// }
//
// func dob() string {
// 	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
// 	return t.Format(time.RFC3339)
// }
//
// func id() string {
// 	return uuid.New().String()
// }
