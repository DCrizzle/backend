package main

import (
	"log"

	"github.com/forstmeier/backend/auth0"
	"github.com/forstmeier/backend/config"
)

func main() {
	log.Println("start token fetching")

	cfg, err := config.New("../../etc/config/config.json")
	if err != nil {
		log.Fatal("error reading config file:", err.Error())
	}

	client := auth0.New(cfg)

	token, err := client.GetUserToken("TEST_FORSTMEIER")
	if err != nil {
		log.Fatal("error fetching token:", err.Error())
	}

	log.Println("token:", token)
}
