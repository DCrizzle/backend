package main

import (
	"log"

	"github.com/forstmeier/backend/config"
)

func main() {
	log.Println("start loading demo")

	cfg, err := config.New("../../etc/config/config.json")
	if err != nil {
		log.Fatal("error reading config file:", err.Error())
	}

	if err := loadDemo(cfg); err != nil {
		log.Fatal("error loading demo: ", err.Error())
	}
	log.Println("end loading demo")
}
