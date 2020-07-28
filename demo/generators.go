package demo

// import (
// 	"math/rand"
// 	"time"
//
// 	"github.com/google/uuid"
// )
//
// const dgraphURL = ""
//
// func generateDonors(count int, owner string) ([]string, error) {
// 	ds := []donor{}
// 	for i := 0; i < count; i++ {
// 		d := donor{
// 			Street:    randomString(streets),
// 			City:      randomString(cities),
// 			County:    randomString(counties),
// 			State:     randomString(states),
// 			ZIP:       randomInt(zips),
// 			Owner:     owner,
// 			Age:       randomInt(ages),
// 			DOB:       dob(),
// 			Sex:       randomString(SEX),
// 			Race:      randomString(RACE),
// 			Specimens: []string{},
// 			Consents:  []string{},
// 		}
//
// 		ds = append(ds, d)
// 	}
//
// 	variables := map[string][]donor{
// 		"input": ds,
// 	}
//
// 	input := payload{
// 		Query:     "",
// 		Variables: variables,
// 	}
//
// 	return sendMutation(input)
// }
//
// func generateOwnerOrg(name string) (string, error) {
// 	now := time.Now().Format(time.RFC3339)
// 	oo := ownerOrg{
// 		org: org{
// 			Street:    randomString(streets),
// 			City:      randomString(cities),
// 			County:    randomString(counties),
// 			State:     randomString(states),
// 			ZIP:       randomInt(zips),
// 			Name:      name,
// 			Users:     []string{},
// 			CreatedOn: now,
// 			UpdatedOn: now,
// 		},
// 		Labs:     []string{},
// 		Storages: []string{},
// 	}
//
// 	variables := map[string]interface{}{
// 		"input": oo,
// 	}
//
// 	input := payload{
// 		Query:     "",
// 		Variables: variables,
// 	}
//
// 	id, err := sendMutation(input)
// 	if err != nil {
// 		return "", err
// 	}
//
// 	return id[0], nil
// }
//
// func generateLabStorageOrgs(name, owner, orgType string) (string, error) {
// 	now := time.Now().Format(time.RFC3339)
// 	var o interface{}
// 	rootOrg := org{
// 		Street:    randomString(streets),
// 		City:      randomString(cities),
// 		County:    randomString(counties),
// 		State:     randomString(states),
// 		ZIP:       randomInt(zips),
// 		Name:      name,
// 		Users:     []string{},
// 		CreatedOn: now,
// 		UpdatedOn: now,
// 	}
//
// 	if orgType == "lab" {
// 		o = labOrg{
// 			org:       rootOrg,
// 			Owner:     owner,
// 			Specimens: []string{},
// 			Plans:     []string{},
// 		}
// 	} else if orgType == "storage" {
// 		o = storageOrg{
// 			org:       rootOrg,
// 			Owner:     owner,
// 			Specimens: []string{},
// 			Plans:     []string{},
// 		}
// 	}
//
// 	variables := map[string]interface{}{
// 		"input": o,
// 	}
//
// 	input := payload{
// 		Query:     "",
// 		Variables: variables,
// 	}
//
// 	id, err := sendMutation(input)
// 	if err != nil {
// 		return "", err
// 	}
//
// 	return id[0], nil
// }
//
// func generateProtocolConsentForm(owner, formType string) (string, error) {
// 	var f interface{}
// 	if formType == "protocol" {
// 		f = protocolForm{
// 			Owner:      owner,
// 			Title:      randomString(titles),
// 			Body:       randomString(bodies),
// 			ProtocolID: id(),
// 			Protocols:  []string{},
// 		}
// 	} else if formType == "consent" {
// 		f = consentForm{
// 			Owner:    owner,
// 			Title:    randomString(titles),
// 			Body:     randomString(bodies),
// 			Consents: []string{},
// 		}
// 	}
//
// 	variables := map[string]interface{}{
// 		"input": f,
// 	}
//
// 	input := payload{
// 		Query:     "",
// 		Variables: variables,
// 	}
//
// 	id, err := sendMutation(input)
// 	if err != nil {
// 		return "", err
// 	}
//
// 	return id[0], nil
// }
//
// func generateProtocol(owner, protocolID string) (string, error) {
// 	p := protocol{
// 		Street:      randomString(streets),
// 		City:        randomString(cities),
// 		County:      randomString(counties),
// 		State:       randomString(states),
// 		ZIP:         randomInt(zips),
// 		Owner:       owner,
// 		Name:        randomString(protocolNames),
// 		Description: randomString(descriptions),
// 		Form:        protocolID,
// 		Plan:        "non_id",
// 		Ages:        randomInts(5, ages),
// 		AgeStart:    25,
// 		AgeEnd:      45,
// 		DOBs:        []string{},
// 		DOBStart:    "",
// 		DOBEnd:      "",
// 		Race:        randomString(RACE),
// 		Sex:         randomString(SEX),
// 		Specimens:   []string{},
// 	}
//
// 	variables := map[string]interface{}{
// 		"input": p,
// 	}
//
// 	input := payload{
// 		Query:     "",
// 		Variables: variables,
// 	}
//
// 	id, err := sendMutation(input)
// 	if err != nil {
// 		return "", err
// 	}
//
// 	return id[0], nil
// }
//
// func generateConsent(owner, donor, protocol string) (string, error) {
// 	c := consent{
// 		Owner:           owner,
// 		Donor:           donor,
// 		Specimen:        "",
// 		Protocol:        protocol,
// 		ConsentDate:     "",
// 		RetentionPeriod: 0,
// 		DestructionDate: "",
// 	}
//
// 	variables := map[string]interface{}{
// 		"input": c,
// 	}
//
// 	input := payload{
// 		Query:     "",
// 		Variables: variables,
// 	}
//
// 	id, err := sendMutation(input)
// 	if err != nil {
// 		return "", err
// 	}
//
// 	return id[0], nil
// }
//
// func generateTest(owner string, labs, specimens []string) (string, error) {
// 	t := specimenTest{
// 		Owner:       owner,
// 		Description: randomString(descriptions),
// 		Labs:        labs,
// 		Specimens:   specimens,
// 		Results:     []string{},
// 	}
//
// 	variables := map[string]interface{}{
// 		"input": t,
// 	}
//
// 	input := payload{
// 		Query:     "",
// 		Variables: variables,
// 	}
//
// 	id, err := sendMutation(input)
// 	if err != nil {
// 		return "", err
// 	}
//
// 	return id[0], nil
// }
//
// func generateSpecimens(count int, owner, donor, consent, lab, storage, protocol string) ([]string, error) {
// 	ss := []bloodSpecimen{}
// 	for i := 0; i < count; i++ {
// 		s := bloodSpecimen{
// 			Owner:           owner,
// 			ExternalID:      id(),
// 			Type:            SPECIMEN_TYPE[0],
// 			CollectionDate:  "",
// 			Container:       CONTAINER[0],
// 			Status:          randomString(STATUS),
// 			DestructionDate: "",
// 			Description:     randomString(descriptions),
// 			BloodType:       randomString(BLOOD_TYPE),
// 			Volume:          1.0,
// 			Tests:           []string{},
// 			Donor:           donor,
// 			Consent:         consent,
// 			Lab:             lab,
// 			Storage:         storage,
// 			Protocol:        protocol,
// 		}
//
// 		ss = append(ss, s)
// 	}
//
// 	variables := map[string]interface{}{
// 		"input": ss,
// 	}
//
// 	input := payload{
// 		Query:     "",
// 		Variables: variables,
// 	}
//
// 	return sendMutation(input)
// }
//
