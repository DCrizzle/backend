package loader

import (
	"log"
	// "math/rand"
	"net/http"
	// "github.com/google/uuid"
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

func LoadDemo(cfg Config) {
	results := make(map[string]map[string][]string)

	ac := &auth0Client{
		httpClient: &http.Client{},
	}

	userToken, err := ac.getUserToken(cfg)
	if err != nil {
		log.Fatal("error getting user token:", err.Error())
	}

	dc := &dgraphClient{
		httpClient: &http.Client{},
		dgraphURL:  cfg.DgraphURL,
		userToken:  userToken,
	}

	ownerIDs, err := dc.addOwnerOrgs()
	if err != nil {
		log.Fatal("add owner orgs error:", err.Error())
	}
	log.Println("ownerIDs:", ownerIDs)

	for ownerIndex, ownerID := range ownerIDs {
		if err := ac.updateUserToken(cfg.UserID, ownerID, cfg.Audience, cfg.ManagementToken); err != nil {
			log.Fatal("error updating user token:", err.Error())
		}

		userToken, err := ac.getUserToken(cfg)
		if err != nil {
			log.Fatal("error getting updated user token:", err.Error())
		}
		dc.userToken = userToken

		results[ownerID] = make(map[string][]string)
		labIDs, storageIDs, err := dc.addLabAndStorageOrgs(ownerID)
		if err != nil {
			log.Fatal("add lab/storage orgs error:", err.Error())
		}
		log.Println("labIDs:", labIDs)
		log.Println("storageIDs:", storageIDs)

		results[ownerID]["labs"] = labIDs
		results[ownerID]["storages"] = storageIDs

		_ = ownerIndex // TEMP
		// userIDs, err := dc.addUsers(ownerID, ownerIndex, labIDs, storageIDs)
		// if err != nil {
		// 	log.Fatal("add users error:", err.Error())
		// }
		// log.Println("userIDs:", userIDs)
		//
		// results[ownerID]["users"] = userIDs

		protocolIDs, planIDs, err := dc.addProtocolsAndPlans(ownerID, labIDs, storageIDs)
		if err != nil {
			log.Fatal("add protocols/plans error:", err.Error())
		}
		log.Println("protocolIDs:", protocolIDs)
		log.Println("planIDs:", planIDs)

		results[ownerID]["protocols"] = protocolIDs
		results[ownerID]["plans"] = planIDs

		// externalIDs := make([]string, len(protocolIDs))
		// for i := range externalIDs {
		// 	externalIDs[i] = uuid.New().String()
		// }
		//
		// protocolFormIDs, err := addProtocolForms(ownerID, protocolIDs, externalIDs)
		// if err != nil {
		// 	log.Fatal("add protocol forms error:", err.Error())
		// }
		//
		// results[ownerID]["protocolForms"] = protocolFormIDs
		//
		// consentFormIDs, err := addConsentForms(ownerID, len(protocolIDs))
		// if err != nil {
		// 	log.Fatal("add consent forms error:", err.Error())
		// }
		//
		// results[ownerID]["consentForms"] = consentFormIDs
		//
		// donorIDs, err := addDonor(ownerID)
		// if err != nil {
		// 	log.Fatal("add donors error:", err.Error())
		// }
		//
		// results[ownerID]["donors"] = donorIDs
		//
		// if len(protocolIDs) != len(consentFormIDs) {
		// 	log.Fatalf("inequal protocol and consent form count, protocols: %d, consent forms: %d\n", len(protocolIDs), len(consentFormIDs))
		// }
		//
		// consentIDs := []string{}
		// bloodSpecimenIDs := []string{}
		// for _, donorID := range donorIDs {
		// 	i := rand.Intn(len(protocolIDs))
		// 	consentID, err := addConsent(
		// 		ownerID,
		// 		donorID,
		// 		consentFormIDs[i],
		// 		protocolIDs[i],
		// 	)
		// 	if err != nil {
		// 		log.Fatal("add consent error:", err.Error())
		// 	}
		//
		// 	consentIDs = append(consentIDs, consentID)
		//
		// 	donorSpecimenIDs, err := addBloodSpecimens(
		// 		ownerID,
		// 		donorID,
		// 		consentID,
		// 		protocolIDs[i],
		// 	)
		// 	if err != nil {
		// 		log.Fatal("add blood specimens error:", err.Error())
		// 	}
		//
		// 	bloodSpecimenIDs = append(bloodSpecimenIDs, donorSpecimenIDs...)
		// }
		//
		// results[ownerID]["consents"] = consentIDs
		// results[ownerID]["bloodSpecimens"] = bloodSpecimenIDs
		//
		// testIDs := []string{}
		// resultIDs := []string{}
		// testChunk := 25
		// chunkSize := (len(bloodSpecimenIDs) + testChunk - 1) / testChunk
		// for i := 0; i < len(bloodSpecimenIDs); i += chunkSize {
		// 	end := i + chunkSize
		// 	if end > len(bloodSpecimenIDs) {
		// 		end = len(bloodSpecimenIDs)
		// 	}
		//
		// 	testSpecimens := bloodSpecimenIDs[i:end]
		// 	testID, err := addTest(
		// 		ownerID,
		// 		randomString(labIDs),
		// 		testSpecimens,
		// 	)
		// 	if err != nil {
		// 		log.Fatal("add test error:", err.Error())
		// 	}
		//
		// 	testIDs = append(testIDs, testID)
		//
		// 	resultID, err := addResult(ownerID, testID)
		// 	if err != nil {
		// 		log.Fatal("add result error:", err.Error())
		// 	}
		//
		// 	resultIDs = append(resultIDs, resultID)
		// }
		//
		// results[ownerID]["tests"] = testIDs
		// results[ownerID]["results"] = resultIDs
	}
}
