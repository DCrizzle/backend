package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/forstmeier/backend/demo/loader"
)

func main() {
	log.Println("start demo loading")

	content, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal("error reading config file: ", err.Error())
	}

	cfg := loader.Config{}
	if err := json.Unmarshal(content, &cfg); err != nil {
		log.Fatal("error unmarshalling config file: ", err.Error())
	}

	if err := loader.LoadDemo(cfg); err != nil {
		log.Fatal("error loading demo: ", err.Error())
	}
	log.Println("end demo loading")
}
