package demo

import (
	"log"

	"github.com/google/uuid"
)

func loadDemo() {

	results := make(map[string]map[string][]string)

	ownerIDs, err := addOwnerOrgs()
	if err != nil {
		log.Fatal("add owner orgs error:", err.Error())
	}

	log.Println("ownerIDs:", ownerIDs)

	for _, ownerID := range ownerIDs {

		results[ownerID] = make(map[string][]string)

		labIDs, storageIDs, err := addLabStorageOrgs(ownerID)
		if err != nil {
			log.Fatal("add lab/storage orgs error:", err.Error())
		}

		results[ownerID]["labs"] = labIDs
		results[ownerID]["storages"] = storageIDs

		userIDs, err := addUsers(ownerID, labIDs, storageIDs)
		if err != nil {
			log.Fatal("add users error:", err.Error())
		}

		results[ownerID]["users"] = userIDs

		protocolIDs, planIDs, err := addProtocolsAndPlans(ownerID, labIDs, storageIDs)
		if err != nil {
			log.Fatal("add protocols/plans error:", err.Error())
		}

		results[ownerID]["protocols"] = protocolIDs
		results[ownerID]["plans"] = planIDs

		externalIDs := make([]string, len(protocolIDs))
		for i := range externalIDs {
			externalIDs[i] = uuid.New().String()
		}

		protocolFormIDs, err := addProtocolForms(ownerID, protocolIDs, externalIDs)
		if err != nil {
			log.Fatal("add protocol forms error:", err.Error())
		}

		consentFormIDs, err := addConsentForms(ownerID, len(protocolIDs))
		if err != nil {
			log.Fatal("add consent forms error:", err.Error())
		}

		donorIDs, err := addDonor(ownerID)
		if err != nil {
			log.Fatal("add donors error:", err.Error())
		}

		consentIDs := []string{}
		for _, donorID := range donorIDs {
			consentID, err := addConsent(
				ownerID,
				donorID,
				randomString(consentFormIDs),
				randomString(protocolIDs),
			)
			if err != nil {
				log.Fatal("add consent error:", err.Error())
			}

			consentIDs = append(consentIDs, consentID)
		}

	}
}
