package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/cheggaaa/pb/v3"
	"github.com/google/uuid"

	"github.com/forstmeier/backend/auth0"
	"github.com/forstmeier/backend/config"
	"github.com/forstmeier/backend/graphql"
)

type dgraphClient struct {
	*graphql.Dgraph
}

func loadDemo(cfg *config.Config) error {
	results := make(map[string]map[string][]string)

	cfg, err := config.New("../../etc/config/config.json")
	if err != nil {
		log.Fatal("error reading config file:", err.Error())
	}

	ac := auth0.New(cfg)

	userToken, err := ac.GetUserToken("TEST_FORSTMEIER")
	if err != nil {
		log.Fatalf("error getting user token: %s", err.Error())
	}

	dc := &dgraphClient{
		graphql.New(
			&http.Client{},
			cfg.Folivora.DgraphURL,
			userToken,
		),
	}

	ownerIDs, err := dc.addOwnerOrgs()
	if err != nil {
		log.Fatalf("add owner orgs error: %s", err.Error())
	}

	helperCount := 10 // number of helper methods being called without owner orgs
	bar := pb.StartNew(helperCount * len(ownerIDs))

	for ownerIndex, ownerID := range ownerIDs {
		if err := ac.UpdateUserToken("TEST_FORSTMEIER", ownerID); err != nil {
			log.Fatalf("error updating user token: %s", err.Error())
		}

		userToken, err := ac.GetUserToken("TEST_FORSTMEIER")
		if err != nil {
			log.Fatalf("error getting updated user token: %s", err.Error())
		}

		dc = &dgraphClient{
			graphql.New(
				&http.Client{},
				cfg.Folivora.DgraphURL,
				userToken,
			),
		}

		results[ownerID] = make(map[string][]string)
		labIDs, storageIDs, err := dc.addLabAndStorageOrgs(ownerID)
		if err != nil {
			log.Fatalf("add lab/storage orgs error: %s", err.Error())
		}
		bar.Increment()

		results[ownerID]["labs"] = labIDs
		results[ownerID]["storages"] = storageIDs

		userIDs, err := dc.addUsers(ownerID, ownerIndex, labIDs, storageIDs)
		if err != nil {
			log.Fatalf("add users error: %s", err.Error())
		}
		bar.Increment()

		results[ownerID]["users"] = userIDs

		protocolIDs, planIDs, err := dc.addProtocolsAndPlans(ownerID, labIDs, storageIDs)
		if err != nil {
			log.Fatalf("add protocols/plans error: %s", err.Error())
		}
		bar.Increment()

		results[ownerID]["protocols"] = protocolIDs
		results[ownerID]["plans"] = planIDs

		externalIDs := make([]string, len(protocolIDs))
		for i := range externalIDs {
			externalIDs[i] = uuid.New().String()
		}

		protocolFormIDs, err := dc.addProtocolForms(ownerID, protocolIDs, externalIDs)
		if err != nil {
			log.Fatalf("add protocol forms error: %s", err.Error())
		}
		bar.Increment()

		results[ownerID]["protocolForms"] = protocolFormIDs

		consentFormIDs, err := dc.addConsentForms(ownerID, len(protocolIDs))
		if err != nil {
			log.Fatalf("add consent forms error: %s", err.Error())
		}
		bar.Increment()

		results[ownerID]["consentForms"] = consentFormIDs

		donorIDs, err := dc.addDonor(ownerID)
		if err != nil {
			log.Fatalf("add donors error: %s", err.Error())
		}
		bar.Increment()

		results[ownerID]["donors"] = donorIDs

		if len(protocolIDs) != len(consentFormIDs) {
			log.Fatalf("inequal protocol and consent form count, protocols: %d, consent forms: %d\n", len(protocolIDs), len(consentFormIDs))
		}

		consentIDs := []string{}
		bloodSpecimenIDs := []string{}
		for _, donorID := range donorIDs {
			i := rand.Intn(len(protocolIDs))
			consentID, err := dc.addConsent(
				ownerID,
				donorID,
				consentFormIDs[i],
				protocolIDs[i],
			)
			if err != nil {
				log.Fatalf("add consent error: %s", err.Error())
			}

			consentIDs = append(consentIDs, consentID)

			donorSpecimenIDs, err := dc.addBloodSpecimens(
				ownerID,
				donorID,
				consentID,
				protocolIDs[i],
			)
			if err != nil {
				log.Fatalf("add blood specimens error: %s", err.Error())
			}

			bloodSpecimenIDs = append(bloodSpecimenIDs, donorSpecimenIDs...)
		}
		bar.Increment()
		bar.Increment()

		results[ownerID]["consents"] = consentIDs
		results[ownerID]["bloodSpecimens"] = bloodSpecimenIDs

		splitSpecimenIDs := chunkBy(bloodSpecimenIDs, 25)

		testIDs := []string{}
		resultIDs := []string{}
		for _, chunkSpecimenIDs := range splitSpecimenIDs {
			testID, err := dc.addTest(
				ownerID,
				randomString(labIDs),
				chunkSpecimenIDs,
			)
			if err != nil {
				log.Fatalf("add test error: %s", err.Error())
			}

			testIDs = append(testIDs, testID)

			resultID, err := dc.addResult(ownerID, testID)
			if err != nil {
				log.Fatalf("add result error: %s", err.Error())
			}

			resultIDs = append(resultIDs, resultID)
		}
		bar.Increment()
		bar.Increment()

		results[ownerID]["tests"] = testIDs
		results[ownerID]["results"] = resultIDs
	}

	bar.Finish()

	jsonData, err := json.Marshal(results)
	if err != nil {
		log.Fatalf("error marshalling results data: %s", err.Error())
	}

	if err := ioutil.WriteFile("results.json", jsonData, 0644); err != nil {
		log.Fatalf("error writing results to file: %s", err.Error())
	}

	return nil
}
