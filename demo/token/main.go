package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

type Config struct {
	Username string `json:"USERNAME"`
	Password string `json:"PASSWORD"`
	Domain   string `json:"DOMAIN"`
	ClientID string `json:"CLIENT_ID"`
}

func main() {
	log.Println("start token fetching")

	content, err := ioutil.ReadFile("../config.json")
	if err != nil {
		log.Fatal("error reading config file:", err.Error())
	}

	cfg := Config{}
	if err := json.Unmarshal(content, &cfg); err != nil {
		log.Fatal("error unmarshalling config file:", err.Error())
	}

	token, err := getUserToken(cfg)
	if err != nil {
		log.Fatal("error fetching token:", err.Error())
	}

	log.Println("token:", token)
}

func getUserToken(cfg Config) (string, error) {
	grantType := "http://auth0.com/oauth/grant-type/password-realm"
	realm := "Username-Password-Authentication"
	data := fmt.Sprintf(
		"grant_type=%s&username=%s&password=%s&client_id=%s&realm=%s",
		grantType,
		cfg.Username,
		cfg.Password,
		cfg.ClientID,
		realm,
	)

	tokenURL := "https://" + cfg.Domain + "/oauth/token"
	req, err := http.NewRequest(http.MethodPost, tokenURL, strings.NewReader(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	userToken := gjson.Get(string(bodyData), "id_token")
	return userToken.String(), nil
}
