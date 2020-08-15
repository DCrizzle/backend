package loader

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/tidwall/gjson"
)

type payload struct {
	Query     string      `json:"query"`
	Variables interface{} `json:"variables"`
}

type auth0Client struct {
	httpClient *http.Client
}

// https://auth0.com/docs/api/authentication?shell#resource-owner-password
func (ac *auth0Client) getUserToken(cfg Config) (string, error) {
	data := fmt.Sprintf(`{
		"grant_type": "password",
		"username": "%s",
		"password": "%s",
		"audience": "%s",
		"scope": "%s",
		"client_id": "%s",
		"client_secret": "%s"
	}`,
		cfg.Username,
		cfg.Password,
		cfg.Audience,
		cfg.Scope,
		cfg.ClientID,
		cfg.ClientSecret,
	)

	req, err := http.NewRequest(http.MethodPost, cfg.Audience+"/oauth/token", strings.NewReader(data))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.ManagementToken))

	resp, err := ac.httpClient.Do(req)
	if err != nil {
		return "", err
	}

	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	userToken := gjson.Get(string(bodyData), "access_token")

	return userToken.String(), nil
}

// https://auth0.com/docs/api/management/v2#!/Users/patch_users_by_id
func (ac *auth0Client) updateUserToken(userID, orgID, audience, managementToken string) error {
	encodedUserID := url.QueryEscape(userID)
	data := fmt.Sprintf(`{
		"app_metadata": {
			"orgID": "%s"
		}
	}`, orgID)

	req, err := http.NewRequest(http.MethodPatch, audience+"/users/"+encodedUserID, strings.NewReader(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", managementToken))

	resp, err := ac.httpClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("non-200 status received: %d", resp.StatusCode)
	}

	return nil
}
