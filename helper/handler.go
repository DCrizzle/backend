// +build !mock

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	errIncorrectSecret       = "incorrect secret received"
	errIncorrectBody         = "incorrect body received"
	errMarshallingCreateJSON = "error marshalling create json"
	errMarshallingUpdateJSON = "error marshalling update json"
	errCreatingAuth0Request  = "error creating auth0 request"
)

func usersHandler(secret, token, url string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if secret != r.Header.Get("folivora-helper-secret") {
			http.Error(w, errIncorrectSecret, http.StatusBadRequest)
			return
		}

		var requestJSON dgraphRequest
		if err := json.NewDecoder(r.Body).Decode(&requestJSON); err != nil {
			http.Error(w, errIncorrectBody, http.StatusBadRequest)
			return
		}

		var auth0Request *http.Request
		var auth0RequestErr error
		if r.Method == http.MethodPost {
			createUserJSON := createUser{
				Email: requestJSON.Email,
				AppMetadata: appMetadata{
					Role:  requestJSON.Role,
					OrgID: requestJSON.OrgID,
				},
				FirstName:  requestJSON.FirstName,
				LastName:   requestJSON.LastName,
				Connection: "Username-Password-Authentication",
			}

			createUserByte, err := json.Marshal(createUserJSON)
			if err != nil {
				http.Error(w, errMarshallingCreateJSON, http.StatusInternalServerError)
				return
			}

			auth0Request, auth0RequestErr = http.NewRequest(
				"POST",
				url,
				bytes.NewReader(createUserByte),
			)

		} else if r.Method == http.MethodPatch {
			updateUserJSON := updateUser{
				AppMetadata: appMetadata{
					Role: requestJSON.Role,
				},
			}

			updateUserByte, err := json.Marshal(updateUserJSON)
			if err != nil {
				http.Error(w, errMarshallingUpdateJSON, http.StatusInternalServerError)
				return
			}

			auth0Request, auth0RequestErr = http.NewRequest(
				"PATCH",
				url+"/"+requestJSON.UserID,
				bytes.NewReader(updateUserByte),
			)

		} else if r.Method == http.MethodDelete {
			auth0Request, auth0RequestErr = http.NewRequest(
				"DELETE",
				url+"/"+requestJSON.UserID,
				nil,
			)
		}

		if auth0RequestErr != nil {
			http.Error(w, errCreatingAuth0Request, http.StatusInternalServerError)
			return
		}

		auth0Request.Header.Set("Authorization", "Bearer "+token)
		auth0Response, err := http.DefaultClient.Do(auth0Request)
		if err != nil || !checkSuccess(auth0Response.StatusCode) {
			http.Error(w, errCreatingAuth0Request, http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, `{"message": "success"}`)
	})
}

func checkSuccess(status int) bool {
	return status == http.StatusOK || status == http.StatusCreated || status == http.StatusNoContent
}

type dgraphRequest struct {
	UserID    string `json:"user_id"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	OrgID     string `json:"orgID"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type createUser struct {
	Email       string      `json:"email"`
	AppMetadata appMetadata `json:"app_metadata"`
	FirstName   string      `json:"given_name"`
	LastName    string      `json:"family_name"`
	Connection  string      `json:"connection"`
}

type updateUser struct {
	AppMetadata appMetadata `json:"app_metadata"`
}

type appMetadata struct {
	Role  string `json:"role"`
	OrgID string `json:"orgID"`
}
