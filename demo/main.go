package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/forstmeier/backend/demo/loader"
)

type config struct {
	Token string `json:"token"`
}

func main() {
	log.Println("start demo loading")

	content, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal("error reading config file:", err.Error())
	}

	c := config{}
	if err := json.Unmarshal(content, &c); err != nil {
		log.Fatal("error unmarshalling config file:", err.Error())
	}

	loader.LoadDemo(c.Token)
	log.Println("end demo loading")
}
