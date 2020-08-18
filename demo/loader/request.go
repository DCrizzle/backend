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

	resp, err := ac.httpClient.Do(req)
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

// https://auth0.com/docs/api/management/v2#!/Users/patch_users_by_id
func (ac *auth0Client) updateUserToken(userID, orgID, audience, managementToken string) error {
	encodedUserID := url.QueryEscape(userID)
	data := fmt.Sprintf(`{
		"app_metadata": {
			"orgID": "%s"
		}
	}`, orgID)

	userURL := audience + "users/" + encodedUserID
	req, err := http.NewRequest(http.MethodPatch, userURL, strings.NewReader(data))
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
