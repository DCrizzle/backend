package loader

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	// "log" // TEMP
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
	labInputs := []map[string]interface{}{}
	labCount := rand.Intn(len(labs)) + 1

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

	labs, err := dc.sendRequest(addLabOrgsMutation, labInputs)
	if err != nil {
		return nil, nil, err
	}

	labIDValues := gjson.Get(labs, "data.addLabOrg.labOrg.#.id").Array()
	labIDs := []string{}
	for _, id := range labIDValues {
		labIDs = append(labIDs, id.String())
	}

	storageIndex := rand.Intn(len(storages))
	storageInput := []map[string]interface{}{
		map[string]interface{}{
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
		},
	}

	storages, err := dc.sendRequest(addStorageOrgsMutation, storageInput)
	if err != nil {
		return nil, nil, err
	}

	storageIDValues := gjson.Get(storages, "data.addStorageOrg.storageOrg.#.id").Array()
	storageIDs := []string{}
	for _, id := range storageIDValues {
		storageIDs = append(storageIDs, id.String())
	}

	return labIDs, storageIDs, nil
}

func (dc *dgraphClient) addUsers(ownerID string, ownerIndex int, labIDs, storageIDs []string) ([]string, error) {
	inputs := []map[string]interface{}{}

	start := (len(users) / len(orgs)) * ownerIndex
	end := start + (len(users) / len(orgs))
	if end > len(users) {
		end = len(users) - 1
	}

	for _, user := range users[start:end] {
		orgID := ""
		if user.role == "USER_STORAGE" {
			orgID = randomString(storageIDs)
		} else if user.role == "USER_LAB" {
			orgID = randomString(labIDs)
		} else {
			orgID = ownerID
		}

		owner := map[string]string{
			"id": ownerID,
		}
		org := map[string]string{
			"id": orgID,
		}
		input := map[string]interface{}{
			"owner":     owner,
			"email":     user.email,
			"firstName": user.first,
			"lastName":  user.last,
			"role":      user.role,
			"org":       org,
			"user_id":   user.userID,
		}

		inputs = append(inputs, input)
	}

	data, err := dc.sendRequest(addUsersMutation, inputs)
	if err != nil {
		return nil, err
	}

	emailValues := gjson.Get(data, "data.addUser.user.#.email").Array()
	emails := []string{}
	for _, email := range emailValues {
		emails = append(emails, email.String())
	}

	return emails, nil
}

func (dc *dgraphClient) addProtocolsAndPlans(ownerID string, labIDs, storageIDs []string) ([]string, []string, error) {
	dobStart := time.Date(1977, time.May, 25, 22, 0, 0, 0, time.UTC)
	dobEnd := time.Date(2005, time.May, 19, 22, 0, 0, 0, time.UTC)
	owner := map[string]string{
		"id": ownerID,
	}

	emptyForm := map[string]interface{}{
		"owner":      owner,
		"title":      "",
		"body":       "",
		"protocolID": "",
	}
	emptyPlan := map[string]interface{}{
		"owner":    owner,
		"name":     "",
		"labs":     []string{},
		"storages": []string{},
	}

	protocolInputs := []map[string]interface{}{}
	for _, protocolName := range protocolNames {
		input := map[string]interface{}{
			"street":      randomString(streets),
			"city":        randomString(cities),
			"county":      randomString(counties),
			"state":       randomString(states),
			"zip":         randomInt(zips),
			"country":     randomString(countries),
			"owner":       owner,
			"name":        protocolName,
			"description": randomString(descriptions),
			"form":        emptyForm,
			"plan":        emptyPlan,
			"dobStart":    dobStart.String(),
			"dobEnd":      dobEnd.String(),
			"race": []string{
				randomString(race),
			},
			"sex": []string{
				randomString(sex),
			},
			"specimens": make([]map[string]string, 0),
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
	if len(protocolIDs) == 0 {
		return nil, nil, errors.New("no protocol ids returned from add mutation")
	}

	planInputs := []map[string]interface{}{}
	for j, planName := range planNames {
		lab := map[string]string{
			"id": randomString(labIDs),
		}
		storage := map[string]string{
			"id": randomString(storageIDs),
		}
		protocol := map[string]string{
			"id": protocolIDs[j],
		}

		input := map[string]interface{}{
			"owner": owner,
			"name":  planName,
			"labs": []map[string]string{
				lab,
			},
			"storages": []map[string]string{
				storage,
			},
			"protocol": protocol,
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
	owner := map[string]string{
		"id": ownerID,
	}
	for donorCount > 0 {
		dob, age := randomDOBAndAge()

		input := map[string]interface{}{
			"street":    randomString(streets),
			"city":      randomString(cities),
			"county":    randomString(counties),
			"state":     randomString(states),
			"zip":       randomInt(zips),
			"country":   randomString(countries),
			"owner":     owner,
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

	status := randomString(status)

	for specimenCount > 0 {
		input := map[string]interface{}{
			"externalID":     uuid.New().String(),
			"type":           specimenType[0],
			"collectionDate": collectionDate,
			"donor":          donor,
			"container":      container[0],
			"status":         status,
			"description":    randomString(descriptions),
			"consent":        consent,
			"owner":          owner,
			"protocol":       protocol,
			"tests":          []string{},
			"bloodType":      randomString(bloodType),
			"volume":         1.0,
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
