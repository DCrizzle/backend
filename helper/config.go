package main

import (
	"encoding/json"
	"io/ioutil"
)

type config struct {
	Auth0    auth0    `json:"AUTH0"`
	Folivora folivora `json:"FOLIVORA"`
}

type auth0 struct {
	ClientID     string `json:"CLIENT_ID"`
	ClientSecret string `json:"CLIENT_SECRET"`
	Audience     string `json:"AUDIENCE"`
}

type folivora struct {
	HelperSecret string `json:"HELPER_SECRET"`
}

func readConfig() (*config, error) {
	configBytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	c := &config{}
	if err := json.Unmarshal(configBytes, c); err != nil {
		return nil, err
	}

	return c, nil
}
