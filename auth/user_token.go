package auth

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/tidwall/gjson"
)

// GetUserToken fetches a user access token for the given user stored in
// the config.json file.
//
// Auth0 docs reference: https://auth0.com/docs/api/authentication?shell#resource-owner-password
func (a *Auth) GetUserToken(user string) (string, error) {
	grantType := "http://auth0.com/oauth/grant-type/password-realm"
	realm := "Username-Password-Authentication"
	cfgUser := a.config.Auth0.Users[user]
	data := fmt.Sprintf(
		"grant_type=%s&username=%s&password=%s&client_id=%s&realm=%s",
		grantType,
		cfgUser.Username,
		cfgUser.Password,
		a.config.Auth0.Frontend.ClientID,
		realm,
	)

	tokenURL := "https://" + a.config.Auth0.DomainURL + "/oauth/token"
	req, err := http.NewRequest(http.MethodPost, tokenURL, strings.NewReader(data))
	if err != nil {
		return "", newErrorNewRequest(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := a.client.Do(req)
	if err != nil {
		return "", newErrorClientDo(err)
	}

	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", newErrorReadResponseBody(err)
	}

	userToken := gjson.Get(string(bodyData), "id_token")
	return userToken.String(), nil
}

// UpdateUserToken sets the "orgID" field on the user app_metadata.
//
// Auth0 docs reference: https://auth0.com/docs/api/management/v2#!/Users/patch_users_by_id
func (a *Auth) UpdateUserToken(user, orgID, managementToken string) error {
	encodedUserID := url.QueryEscape(a.config.Auth0.Users[user].ID)
	data := fmt.Sprintf(`{
		"app_metadata": {
			"orgID": "%s"
		}
	}`, orgID)

	userURL := a.config.Auth0.AudienceURL + "users/" + encodedUserID
	req, err := http.NewRequest(http.MethodPatch, userURL, strings.NewReader(data))
	if err != nil {
		return newErrorNewRequest(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", managementToken))

	resp, err := a.client.Do(req)
	if err != nil {
		return newErrorClientDo(err)
	}
	if resp.StatusCode == http.StatusUnauthorized {
		return newErrorStatusUnauthorized("management api token may be expired")
	}
	if resp.StatusCode != http.StatusOK {
		return newErrorNon200Status(resp.StatusCode)
	}

	return nil
}
