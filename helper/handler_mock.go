// +build mock

package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

const (
	errIncorrectRequestBody = "incorrect request body received"
	errDgraphMutation       = "error executing dgraph mutation"
)

func usersHandler(secret, token, auth0URL, dgraphURL string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var dgraphReqJSON dgraphRequest
		if err := json.NewDecoder(r.Body).Decode(&dgraphReqJSON); err != nil {
			http.Error(w, errIncorrectRequestBody, http.StatusBadRequest)
			return
		}

		auth0ID := ""
		var dgraphMutation string
		var dgraphVariables interface{}

		if r.Method == http.MethodPost {
			id := uuid.New().String()
			hexID := hex.EncodeToString([]byte(id))
			auth0ID = "auth0|" + hexID

			dgraphMutation = createUserMutation
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
				"auth0ID": auth0ID,
			}
		} else if r.Method == http.MethodPatch {
			auth0ID = dgraphReqJSON.Auth0ID

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
			auth0ID = dgraphReqJSON.Auth0ID

			dgraphMutation = removeUserMutation
			dgraphVariables = map[string]interface{}{
				"id": dgraphReqJSON.Owner,
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

		fmt.Fprintf(w, fmt.Sprintf(`{"message": "success", "auth0ID": "%s"}`, auth0ID))
	})
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
