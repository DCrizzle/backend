// +build !mock

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/forstmeier/backend/graphql"
)

func usersHandler(folivoraSecret, managementToken, auth0URL, dgraphURL string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if folivoraSecret != r.Header.Get("folivora-helper-secret") {
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

		// outline:
		// [ ] note: split into helper functions
		// [ ] "create auth0 request"
		// [ ] "create dgraph inputs"
		if r.Method == http.MethodPost {
			createUserJSON := createUserRequest{
				Email:    *dgraphReqJSON.Email,
				Password: *dgraphReqJSON.Password,
				AppMetadata: appMetadata{
					Role:  dgraphReqJSON.Role,
					OrgID: dgraphReqJSON.Owner,
				},
				FirstName:  *dgraphReqJSON.FirstName,
				LastName:   *dgraphReqJSON.LastName,
				Connection: "Username-Password-Authentication",
			}

			createUserByte, err := json.Marshal(createUserJSON)
			if err != nil {
				http.Error(w, errMarshallingCreateJSON, http.StatusInternalServerError)
				return
			}

			auth0CreateURL := auth0URL + "users"
			auth0Req, auth0Err = http.NewRequest(
				http.MethodPost,
				auth0CreateURL,
				bytes.NewReader(createUserByte),
			)
			dgraphMutation = graphql.AddUsersMutation
		} else if r.Method == http.MethodPatch {
			updateUserJSON := updateUserRequest{}
			updateUserVariables := map[string]interface{}{
				"filter": map[string]interface{}{
					"auth0ID": map[string]interface{}{
						"eq": *dgraphReqJSON.Auth0ID,
					},
				},
			}

			if dgraphReqJSON.Password != nil {
				updateUserJSON.Password = dgraphReqJSON.Password
			}
			if dgraphReqJSON.Role != nil {
				updateUserJSON.AppMetadata = appMetadata{
					Role: dgraphReqJSON.Role,
				}
				updateUserVariables["filter"] = map[string]interface{}{
					"auth0ID": map[string]interface{}{
						"eq": *dgraphReqJSON.Auth0ID,
					},
				}
				updateUserVariables["set"] = map[string]interface{}{
					"role": *dgraphReqJSON.Role,
				}
			}

			updateUserByte, err := json.Marshal(updateUserJSON)
			if err != nil {
				http.Error(w, errMarshallingUpdateJSON, http.StatusInternalServerError)
				return
			}

			auth0UpdateURL := auth0URL + "users/" + url.PathEscape(*dgraphReqJSON.Auth0ID)
			auth0Req, auth0Err = http.NewRequest(
				http.MethodPatch,
				auth0UpdateURL,
				bytes.NewReader(updateUserByte),
			)
			dgraphMutation = graphql.UpdateUserMutation
			dgraphVariables = updateUserVariables
		} else if r.Method == http.MethodDelete {
			auth0DeleteURL := auth0URL + "users/" + url.PathEscape(*dgraphReqJSON.Auth0ID)
			auth0Req, auth0Err = http.NewRequest(
				http.MethodDelete,
				auth0DeleteURL,
				nil,
			)

			dgraphMutation = graphql.DeleteUserMutation
			dgraphVariables = map[string]interface{}{
				"auth0ID": map[string]interface{}{
					"eq": *dgraphReqJSON.Auth0ID,
				},
			}
		} else {
			http.Error(w, errIncorrectHTTPMethod, http.StatusBadRequest)
			return
		}

		if auth0Err != nil {
			http.Error(w, errCreatingAuth0Request, http.StatusInternalServerError)
			return
		}

		auth0Req.Header.Set("Authorization", "Bearer "+managementToken)
		auth0Req.Header.Set("Content-Type", "application/json")

		httpClient := &http.Client{}

		auth0Resp, err := httpClient.Do(auth0Req)
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
					"auth0ID": auth0RespJSON.Auth0ID,
				},
			}
		}

		dgraphClient := graphql.New(
			httpClient,
			dgraphURL,
			r.Header.Get("X-Auth0-Token"),
		)

		_, err = dgraphClient.SendRequest(dgraphMutation, dgraphVariables)
		if err != nil {
			http.Error(w, errDgraphMutation, http.StatusBadRequest)
			return
		}

		if r.Method == http.MethodPost {
			fmt.Fprintf(w, fmt.Sprintf(`{"message": "success", "auth0ID": "%s"}`, auth0RespJSON.Auth0ID))
		} else {
			fmt.Fprintf(w, fmt.Sprintf(`{"message": "success", "auth0ID": "%s"}`, *dgraphReqJSON.Auth0ID))
		}
	})
}

func checkSuccess(status int) bool {
	return status == http.StatusOK || status == http.StatusCreated || status == http.StatusNoContent
}
