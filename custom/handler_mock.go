// +build mock

package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/forstmeier/backend/graphql"
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

			dgraphMutation = graphql.AddUsersMutation
			dgraphVariables = []map[string]interface{}{
				map[string]interface{}{
					"owner": map[string]string{
						"id": *dgraphReqJSON.Owner,
					},
					"email":     *dgraphReqJSON.Email,
					"firstName": *dgraphReqJSON.FirstName,
					"lastName":  *dgraphReqJSON.LastName,
					"role":      *dgraphReqJSON.Role,
					"org": map[string]string{
						"id": *dgraphReqJSON.Org,
					},
					"auth0ID": auth0ID,
				},
			}
		} else if r.Method == http.MethodPatch {
			auth0ID = *dgraphReqJSON.Auth0ID

			dgraphMutation = graphql.UpdateUserMutation
			userUpdates := map[string]interface{}{
				"filter": map[string]interface{}{
					"auth0ID": map[string]interface{}{
						"eq": *dgraphReqJSON.Auth0ID,
					},
				},
			}

			if dgraphReqJSON.Role != nil {
				userUpdates["filter"] = map[string]interface{}{
					"auth0ID": map[string]interface{}{
						"eq": *dgraphReqJSON.Auth0ID,
					},
				}
				userUpdates["set"] = map[string]interface{}{
					"role": *dgraphReqJSON.Role,
				}
			}
			dgraphVariables = userUpdates
		} else if r.Method == http.MethodDelete {
			auth0ID = *dgraphReqJSON.Auth0ID

			dgraphMutation = graphql.DeleteUserMutation
			dgraphVariables = map[string]interface{}{
				"auth0ID": map[string]interface{}{
					"eq": *dgraphReqJSON.Auth0ID,
				},
			}
		}

		dgraphClient := graphql.New(
			&http.Client{},
			dgraphURL,
			r.Header.Get("X-Auth0-Token"),
		)

		_, err := dgraphClient.SendRequest(dgraphMutation, dgraphVariables)
		if err != nil {
			http.Error(w, errDgraphMutation, http.StatusBadRequest)
			return
		}

		if r.Method == http.MethodPost {
			dgraphResponseBytes, err := json.Marshal(dgraphVariables.([]map[string]interface{})[0])
			if err != nil {
				http.Error(w, errMarshallingDgraphJSON, http.StatusBadRequest)
				return
			}
			fmt.Fprintf(w, string(dgraphResponseBytes))
		} else {
			// outline:
			// [ ] add in values from auth0 response / received dgraph request
			responseBody := fmt.Sprintf(`{"owner": {"id": ""}, "email": "", "firstName": "", "lastName": "", "role": "", "org": {"id": ""}, "auth0ID": "%s"}`, *dgraphReqJSON.Auth0ID)
			fmt.Fprintf(w, responseBody)
		}
	})
}
