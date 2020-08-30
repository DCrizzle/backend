// +build !mock

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	errIncorrectSecret       = "incorrect secret received"
	errIncorrectRequestBody  = "incorrect request body received"
	errIncorrectHTTPMethod   = "unsupported http method received"
	errMarshallingCreateJSON = "error marshalling create json"
	errMarshallingUpdateJSON = "error marshalling update json"
	errCreatingAuth0Request  = "error creating auth0 request"
	errExecutingAuth0Request = "error executing auth0 request"
	errIncorrectResponseBody = "incorrect response body received"
	errDgraphMutation        = "error executing dgraph mutation"
)

const createUserMutation = "mutation CreateUser($input: [AddUserInput!]!) { addUser(input: $input) { user { email } } }"

const editUserMutation = "mutation EditUser($input: UpdateUserInput!) { updateUser(input: $input) { user { email } } }"

const removeUserMutation = "mutation RemoveUser($filter: UserFilter!) { deleteUser(filter: $filter) { user { email } } }"

func usersHandler(secret, token, auth0URL, dgraphURL string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if secret != r.Header.Get("folivora-helper-secret") {
			http.Error(w, errIncorrectSecret, http.StatusBadRequest)
			return
		}

		var dgraphReqJSON dgraphRequest
		if err := json.NewDecoder(r.Body).Decode(&dgraphReqJSON); err != nil {
			http.Error(w, errIncorrectRequestBody, http.StatusBadRequest)
			return
		}

		var auth0Req *http.Request
		var auth0Err error

		var dgraphMutation string
		var dgraphVariables interface{}

		// NOTE: split into helper functions
		if r.Method == http.MethodPost {
			createUserJSON := createUserRequest{
				Email:    dgraphReqJSON.Email,
				Password: dgraphReqJSON.Password,
				AppMetadata: appMetadata{
					Role:  dgraphReqJSON.Role,
					OrgID: dgraphReqJSON.Owner,
				},
				FirstName:  dgraphReqJSON.FirstName,
				LastName:   dgraphReqJSON.LastName,
				Connection: "Username-Password-Authentication",
			}

			createUserByte, err := json.Marshal(createUserJSON)
			if err != nil {
				http.Error(w, errMarshallingCreateJSON, http.StatusInternalServerError)
				return
			}

			auth0CreateURL := auth0URL + "/users"
			auth0Req, auth0Err = http.NewRequest(
				http.MethodPost,
				auth0CreateURL,
				bytes.NewReader(createUserByte),
			)
			dgraphMutation = createUserMutation
		} else if r.Method == http.MethodPatch {
			updateUserJSON := updateUserRequest{}
			if dgraphReqJSON.Password != "" {
				updateUserJSON.Password = dgraphReqJSON.Password
			}
			if dgraphReqJSON.Role != "" {
				updateUserJSON.AppMetadata = appMetadata{
					Role: dgraphReqJSON.Role,
				}
			}

			updateUserByte, err := json.Marshal(updateUserJSON)
			if err != nil {
				http.Error(w, errMarshallingUpdateJSON, http.StatusInternalServerError)
				return
			}

			auth0UpdateURL := auth0URL + "/users/" + url.PathEscape(dgraphReqJSON.Auth0ID)
			auth0Req, auth0Err = http.NewRequest(
				http.MethodPatch,
				auth0UpdateURL,
				bytes.NewReader(updateUserByte),
			)
			dgraphMutation = editUserMutation

			userUpdates := make(map[string]interface{})
			if dgraphReqJSON.Role != "" {
				userUpdates["role"] = dgraphReqJSON.Role
			}
			userUpdates["owner"] = map[string]string{
				"id": dgraphReqJSON.Owner,
			}
			dgraphVariables = userUpdates
		} else if r.Method == http.MethodDelete {
			auth0DeleteURL := auth0URL + "/users/" + url.PathEscape(dgraphReqJSON.Auth0ID)
			auth0Req, auth0Err = http.NewRequest(
				http.MethodDelete,
				auth0DeleteURL,
				nil,
			)

			dgraphMutation = removeUserMutation
			dgraphVariables = map[string]interface{}{
				"id": dgraphReqJSON.Owner,
			}
		} else {
			http.Error(w, errIncorrectHTTPMethod, http.StatusBadRequest)
			return
		}

		if auth0Err != nil {
			http.Error(w, errCreatingAuth0Request, http.StatusInternalServerError)
			return
		}

		auth0Req.Header.Set("Authorization", "Bearer "+token)
		auth0Req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}

		auth0Resp, err := client.Do(auth0Req)
		defer auth0Resp.Body.Close()
		if err != nil || !checkSuccess(auth0Resp.StatusCode) {
			http.Error(w, errExecutingAuth0Request, http.StatusInternalServerError)
			return
		}

		var auth0RespJSON auth0Response
		if r.Method == http.MethodPost {
			if err := json.NewDecoder(auth0Resp.Body).Decode(&auth0RespJSON); err != nil {
				http.Error(w, errIncorrectResponseBody, http.StatusBadRequest)
				return
			}

			dgraphVariables = map[string]interface{}{
				"owner": map[string]string{
					"id": dgraphReqJSON.Owner,
				},
				"email":     dgraphReqJSON.Email,
				"password":  dgraphReqJSON.Password,
				"firstName": dgraphReqJSON.FirstName,
				"lastName":  dgraphReqJSON.LastName,
				"role":      dgraphReqJSON.Role,
				"org": map[string]string{
					"id": dgraphReqJSON.Org,
				},
				"auth0ID": auth0RespJSON.Auth0ID,
			}
		}

		if err := sendDgraphRequest(
			dgraphURL,
			r.Header.Get("X-Auth0-Token"),
			dgraphMutation,
			dgraphVariables,
		); err != nil {
			http.Error(w, errDgraphMutation, http.StatusBadRequest)
			return
		}

		if r.Method == http.MethodPost {
			fmt.Fprintf(w, fmt.Sprintf(`{"message": "success", "auth0ID": "%s"}`, auth0RespJSON.Auth0ID))
		} else {
			fmt.Fprintf(w, fmt.Sprintf(`{"message": "success", "auth0ID": "%s"}`, dgraphReqJSON.Auth0ID))
		}
	})
}

func checkSuccess(status int) bool {
	return status == http.StatusOK || status == http.StatusCreated || status == http.StatusNoContent
}

type dgraphRequest struct {
	Owner     string `json:"owner,omitempty"`
	Auth0ID   string `json:"authZeroID,omitempty"` // NOTE: "authZeroID" field name is necessary due to a Dgraph limitation
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	Role      string `json:"role,omitempty"`
	Org       string `json:"org,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}

type createUserRequest struct {
	Email       string      `json:"email"`
	Password    string      `json:"password"`
	AppMetadata appMetadata `json:"app_metadata"`
	FirstName   string      `json:"given_name"`
	LastName    string      `json:"family_name"`
	Connection  string      `json:"connection"`
}

type updateUserRequest struct {
	AppMetadata appMetadata `json:"app_metadata,omitempty"`
	Password    string      `json:"password,omitempty"`
	Connection  string      `json:"connection,omitempty"`
}

type appMetadata struct {
	Role  string `json:"role,omitempty"`
	OrgID string `json:"orgID,omitempty"` // NOTE: possibly change to "ownerID" for consistency
}

type auth0Response struct {
	Auth0ID string `json:"user_id"`
}
