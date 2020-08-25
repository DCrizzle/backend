package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type responseJSON struct {
	AccessToken string `json:"access_token"`
}

func getAuth0APIToken(url string, auth0Config auth0) (string, error) {
	payloadString := fmt.Sprintf(
		"grant_type=client_credentials&client_id=%s&client_secret=%s&audience=%s",
		auth0Config.ClientID,
		auth0Config.ClientSecret,
		auth0Config.Audience,
	)

	payload := strings.NewReader(payloadString)
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	resJSON := responseJSON{}
	if err := json.Unmarshal(resBytes, &resJSON); err != nil {
		return "", err
	}

	return resJSON.AccessToken, nil
}
