package auth0

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

type responseJSON struct {
	AccessToken string `json:"access_token"`
}

// GetManagementAPIToken fetches a token for the Auth0 Management API.
func (a *Auth0) GetManagementAPIToken() (string, error) {
	backend := a.config.Auth0.Backend
	payloadString := fmt.Sprintf(
		"grant_type=client_credentials&client_id=%s&client_secret=%s&audience=%s",
		backend.ClientID,
		backend.ClientSecret,
		a.config.Auth0.AudienceURL,
	)

	payload := strings.NewReader(payloadString)
	req, err := http.NewRequest("POST", a.config.Auth0.TokenURL, payload)
	if err != nil {
		return "", errAuth0(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := a.client.Do(req)
	if err != nil {
		return "", errAuth0(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errAuth0(err)
	}

	managementToken := gjson.Get(string(bodyBytes), "access_token")
	return managementToken.String(), nil
}
