package demo

import "log"

func loadDemo() {

	ownerIDs, err := addOwnerOrgs()
	if err != nil {
		log.Fatal("add owner orgs error:", err.Error())
	}

	log.Println("ownerIDs:", ownerIDs)

}
