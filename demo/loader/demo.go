package loader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/cheggaaa/pb/v3"
	"github.com/google/uuid"
)

type Config struct {
	UserID          string `json:"USER_ID"`
	ManagementToken string `json:"MANAGEMENT_TOKEN"`
	Username        string `json:"USERNAME"`
	Password        string `json:"PASSWORD"`
	Domain          string `json:"DOMAIN"`
	Audience        string `json:"AUDIENCE"`
	ClientID        string `json:"CLIENT_ID"`
	ClientSecret    string `json:"CLIENT_SECRET"`
	DgraphURL       string `json:"DGRAPH_URL"`
}

func LoadDemo(cfg Config) error {
	results := make(map[string]map[string][]string)

	ac := &auth0Client{
		httpClient: &http.Client{},
	}

	userToken, err := ac.getUserToken(cfg)
	if err != nil {
		return fmt.Errorf("error getting user token: %w", err)
	}

	dc := &dgraphClient{
		httpClient: &http.Client{},
		dgraphURL:  cfg.DgraphURL,
		userToken:  userToken,
	}

	ownerIDs, err := dc.addOwnerOrgs()
	if err != nil {
		return fmt.Errorf("add owner orgs error: %w", err)
	}

	helperCount := 10 // number of helper methods being called without owner orgs
	bar := pb.StartNew(helperCount * len(ownerIDs))

	for ownerIndex, ownerID := range ownerIDs {
		if err := ac.updateUserToken(cfg.UserID, ownerID, cfg.Audience, cfg.ManagementToken); err != nil {
			return fmt.Errorf("error updating user token: %w", err)
		}

		userToken, err := ac.getUserToken(cfg)
		if err != nil {
			return fmt.Errorf("error getting updated user token: %w", err)
		}
		dc.userToken = userToken

		results[ownerID] = make(map[string][]string)
		labIDs, storageIDs, err := dc.addLabAndStorageOrgs(ownerID)
		if err != nil {
			return fmt.Errorf("add lab/storage orgs error: %w", err)
		}
		bar.Increment()

		results[ownerID]["labs"] = labIDs
		results[ownerID]["storages"] = storageIDs

		userIDs, err := dc.addUsers(ownerID, ownerIndex, labIDs, storageIDs)
		if err != nil {
			return fmt.Errorf("add users error: %w", err)
		}
		bar.Increment()

		results[ownerID]["users"] = userIDs

		protocolIDs, planIDs, err := dc.addProtocolsAndPlans(ownerID, labIDs, storageIDs)
		if err != nil {
			return fmt.Errorf("add protocols/plans error: %w", err)
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
			return fmt.Errorf("add protocol forms error: %w", err)
		}
		bar.Increment()

		results[ownerID]["protocolForms"] = protocolFormIDs

		consentFormIDs, err := dc.addConsentForms(ownerID, len(protocolIDs))
		if err != nil {
			return fmt.Errorf("add consent forms error: %w", err)
		}
		bar.Increment()

		results[ownerID]["consentForms"] = consentFormIDs

		donorIDs, err := dc.addDonor(ownerID)
		if err != nil {
			return fmt.Errorf("add donors error: %w", err)
		}
		bar.Increment()

		results[ownerID]["donors"] = donorIDs

		if len(protocolIDs) != len(consentFormIDs) {
			return fmt.Errorf("inequal protocol and consent form count, protocols: %d, consent forms: %d\n", len(protocolIDs), len(consentFormIDs))
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
				return fmt.Errorf("add consent error: %w", err)
			}

			consentIDs = append(consentIDs, consentID)

			donorSpecimenIDs, err := dc.addBloodSpecimens(
				ownerID,
				donorID,
				consentID,
				protocolIDs[i],
			)
			if err != nil {
				return fmt.Errorf("add blood specimens error: %w", err)
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
				return fmt.Errorf("add test error: %w", err)
			}

			testIDs = append(testIDs, testID)

			resultID, err := dc.addResult(ownerID, testID)
			if err != nil {
				return fmt.Errorf("add result error: %w", err)
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
		return fmt.Errorf("error marshalling results data: %w", err)
	}

	if err := ioutil.WriteFile("results.json", jsonData, 0644); err != nil {
		return fmt.Errorf("error writing results to file: %w", err)
	}

	return nil
}
