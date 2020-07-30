package demo

import "log"

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

		protocolIDs, protocolFormIDs, planIDs, err := addProtocolsFormsPlans(ownerID, labIDs, storageIDs)
		if err != nil {
			log.Fatal("add protocols/protocol forms/plans error:", err.Error())
		}

		results[ownerID]["protocols"] = protocolIDs
		results[ownerID]["protocolForms"] = protocolFormIDs
		results[ownerID]["plans"] = planIDs

	}
}
