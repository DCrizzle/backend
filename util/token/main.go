package main

import (
	"log"

	"github.com/forstmeier/backend/auth"
	"github.com/forstmeier/backend/config"
)

func main() {
	log.Println("start token fetching")

	cfg, err := config.New("../../etc/config/config.json")
	if err != nil {
		log.Fatal("error reading config file:", err.Error())
	}

	ac := auth.New(cfg)

	token, err := ac.GetUserToken("TEST_FORSTMEIER")
	if err != nil {
		log.Fatal("error fetching token:", err.Error())
	}

	log.Println("token:", token)
}
