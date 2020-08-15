package loader

import (
	"bytes"
	"encoding/json"
	"fmt" // TEMP
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/tidwall/gjson"
)

type dgraphClient struct {
	httpClient *http.Client
	dgraphURL  string
	userToken  string
}

func (dc *dgraphClient) addOwnerOrgs() ([]string, error) {
	inputs := []map[string]interface{}{}
	for i := 0; i < len(orgs); i++ {
		now := time.Now().Format(time.RFC3339)
		input := map[string]interface{}{
			"street":    randomString(streets),
			"city":      randomString(cities),
			"county":    randomString(counties),
			"state":     randomString(states),
			"zip":       randomInt(zips),
			"country":   randomString(countries),
			"name":      orgs[i],
			"users":     []string{},
			"createdOn": now,
			"updatedOn": now,
			"labs":      []string{},
			"storages":  []string{},
		}
		inputs = append(inputs, input)
	}

	data, err := dc.sendRequest(addOwnerOrgsMutation, inputs)
	if err != nil {
		return nil, err
	}

	idValues := gjson.Get(data, "data.addOwnerOrg.ownerOrg.#.id").Array()
	ids := []string{}
	for _, id := range idValues {
		ids = append(ids, id.String())
	}

	return ids, nil
}

func (dc *dgraphClient) addLabAndStorageOrgs(ownerID string) ([]string, []string, error) {
	fmt.Println("addLabAndStorageOrgs")
	labInputs := []map[string]interface{}{}
	labCount := rand.Intn(len(labs))

	owner := map[string]string{
		"id": ownerID,
	}
	now := time.Now().Format(time.RFC3339)

	for labCount > 0 {
		labInput := map[string]interface{}{
			"street":    randomString(streets),
			"city":      randomString(cities),
			"county":    randomString(counties),
			"state":     randomString(states),
			"zip":       randomInt(zips),
			"country":   randomString(countries),
			"name":      labs[labCount-1],
			"users":     []string{},
			"createdOn": now,
			"updatedOn": now,
			"owner":     owner,
			"specimens": []string{},
			"plans":     []string{},
		}

		labInputs = append(labInputs, labInput)
		labCount--
	}
	fmt.Println("lab inputs:", labInputs)

	labs, err := dc.sendRequest(addLabOrgsMutation, labInputs)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println("labs:", labs)

	labIDValues := gjson.Get(labs, "data.addLabOrg.labOrg.#.id").Array()
	labIDs := []string{}
	for _, id := range labIDValues {
		labIDs = append(labIDs, id.String())
	}

	storageIndex := rand.Intn(len(storages))
	storageInput := map[string]interface{}{
		"street":    randomString(streets),
		"city":      randomString(cities),
		"county":    randomString(counties),
		"state":     randomString(states),
		"country":   randomString(countries),
		"zip":       randomInt(zips),
		"name":      storages[storageIndex],
		"users":     []string{},
		"createdOn": now,
		"updatedOn": now,
		"owner":     owner,
		"specimens": []string{},
		"plans":     []string{},
	}

	storages, err := dc.sendRequest(addStorageOrgsMutation, storageInput)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println("storages:", storages)

	storageIDValues := gjson.Get(storages, "data.addStorageOrg.#.storageOrg.id").Array()
	storageIDs := []string{}
	for _, id := range storageIDValues {
		storageIDs = append(storageIDs, id.String())
	}

	return labIDs, storageIDs, nil
}

func (dc *dgraphClient) addUsers(ownerID string, labIDs, storageIDs []string) ([]string, error) {
	inputs := []map[string]interface{}{}
	for _, user := range users {
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

	data, err := dc.sendRequest(addUsersMutation, inputs)
	if err != nil {
		return nil, err
	}

	idValues := gjson.Get(data, "data.addUser.user.#.id").Array()
	ids := []string{}
	for _, id := range idValues {
		ids = append(ids, id.String())
	}

	return ids, nil
}

func (dc *dgraphClient) addProtocolsAndPlans(ownerID string, labIDs, storageIDs []string) ([]string, []string, error) {
	dobStart := time.Date(1977, time.May, 25, 22, 0, 0, 0, time.UTC)
	dobEnd := time.Date(2005, time.May, 19, 22, 0, 0, 0, time.UTC)

	protocolInputs := []map[string]interface{}{}
	for _, protocolName := range protocolNames {
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
			"race":        randomString(race),
			"sex":         randomString(sex),
			"specimens":   "",
		}

		protocolInputs = append(protocolInputs, input)
	}

	protocols, err := dc.sendRequest(addProtocolsMutation, protocolInputs)
	if err != nil {
		return nil, nil, err
	}

	protocolsIDValues := gjson.Get(protocols, "data.addProtocol.protocol.#.id").Array()
	protocolIDs := []string{}
	for _, id := range protocolsIDValues {
		protocolIDs = append(protocolIDs, id.String())
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

	plans, err := dc.sendRequest(addPlansMutation, planInputs)
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

func (dc *dgraphClient) addProtocolForms(ownerID string, protocolIDs, protocolExternalIDs []string) ([]string, error) {
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

	data, err := dc.sendRequest(addProtocolFormsMutation, protocolFormInputs)
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

func (dc *dgraphClient) addConsentForms(ownerID string, count int) ([]string, error) {
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

	data, err := dc.sendRequest(addConsentFormsMutation, consentFormInputs)
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

func (dc *dgraphClient) addDonor(ownerID string) ([]string, error) {
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
			"sex":       randomString(sex),
			"race":      randomString(race),
			"specimens": []string{},
			"consents":  []string{},
		}

		donorInputs = append(donorInputs, input)
		donorCount--
	}

	data, err := dc.sendRequest(addDonorsMutation, donorInputs)
	if err != nil {
		return nil, err
	}

	idValues := gjson.Get(data, "data.addDonor.donor.#.id").Array()
	ids := []string{}
	for _, id := range idValues {
		ids = append(ids, id.String())
	}

	return ids, nil
}

func (dc *dgraphClient) addConsent(ownerID, donorID, formID, protocolID string) (string, error) {
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

	data, err := dc.sendRequest(addConsentsMutation, input)
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

// NOTE: this could be improved to return the input specimenInputs variable
// which would be submitted in bulk by the calling scope
func (dc *dgraphClient) addBloodSpecimens(ownerID, donorID, consentID, protocolID string) ([]string, error) {
	specimenCount := rand.Intn(10) + 1

	specimenInputs := []map[string]interface{}{}

	year := time.Now().Year()
	month := rand.Intn(12) + 1
	day := rand.Intn(25) + 1
	collectionDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC).String()

	status := randomString(status)
	destructionDate := ""
	if status == "DESTROYED" {
		destructionDate = time.Now().String()
	}

	for specimenCount > 0 {
		input := map[string]interface{}{
			"externalID":      uuid.New().String(),
			"type":            specimenType[0],
			"collectionDate":  collectionDate,
			"donor":           donorID,
			"container":       container[0],
			"status":          status,
			"destructionDate": destructionDate,
			"description":     randomString(descriptions),
			"consent":         consentID,
			"owner":           ownerID,
			"lab":             "", // NOTE: add later (?)
			"storage":         "", // NOTE: add later (?)
			"protocol":        protocolID,
			"tests":           []string{},
			"bloodType":       randomString(bloodType),
			"volume":          1.0,
		}

		specimenInputs = append(specimenInputs, input)

		specimenCount--
	}

	data, err := dc.sendRequest(addBloodSpecimensMutation, specimenInputs)
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

func (dc *dgraphClient) addTest(ownerID, labID string, specimenIDs []string) (string, error) {
	input := map[string]interface{}{
		"description": randomString(descriptions),
		"owner":       ownerID,
		"lab":         labID,
		"specimens":   specimenIDs,
		"results":     []string{},
	}

	data, err := dc.sendRequest(addTestsMutation, input)
	if err != nil {
		return "", err
	}

	idValues := gjson.Get(data, "data.addTest.test.#.id").Array()
	ids := []string{}
	for _, id := range idValues {
		ids = append(ids, id.String())
	}

	return ids[0], nil
}

func (dc *dgraphClient) addResult(ownerID, testID string) (string, error) {
	input := map[string]interface{}{
		"owner": ownerID,
		"notes": randomString(notes),
		"test":  testID,
	}

	data, err := dc.sendRequest(addResultsMutation, input)
	if err != nil {
		return "", err
	}

	idValues := gjson.Get(data, "data.addResult.result.#.id").Array()
	ids := []string{}
	for _, id := range idValues {
		ids = append(ids, id.String())
	}

	return ids[0], nil
}

func (dc *dgraphClient) sendRequest(mutation string, input interface{}) (string, error) {
	variables := map[string]interface{}{
		"input": input,
	}

	p := payload{
		Query:     mutation,
		Variables: variables,
	}

	b, err := json.Marshal(p)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", dc.dgraphURL, bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}
	req.Header.Set("X-Auth0-Token", dc.userToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := dc.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
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
