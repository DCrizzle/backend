// +build !mock

package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/forstmeier/backend/custom/handlers"
	"github.com/forstmeier/backend/graphql"
)

// Handler is an HTTP listener for User type @custom directive events.
//
// createUser: adds a user to Auth0 and to Dgraph with the Auth0 ID field
// editUser: updates an Auth0 user role or password in Auth0 and Dgraph
// removeUser: deletes an Auth0 user from Auth0 and Dgraph
func Handler(managementToken, auth0URL, dgraphURL string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var dgraphReqJSON handlers.DgraphRequest
		if err := json.NewDecoder(r.Body).Decode(&dgraphReqJSON); err != nil {
			http.Error(w, handlers.ErrIncorrectRequestBody, http.StatusBadRequest)
			return
		}

		var auth0Req *http.Request
		var auth0Err error

		var dgraphMutation string
		var dgraphVariables interface{}

		if r.Method == http.MethodPost {
			auth0CreateURL := auth0URL + "users"
			auth0Req, auth0Err = createUserAuth0Req(dgraphReqJSON, auth0CreateURL)

			dgraphMutation = graphql.AddUsersMutation
		} else if r.Method == http.MethodPatch {
			auth0UpdateURL := auth0URL + "users/" + url.PathEscape(*dgraphReqJSON.Auth0ID)
			auth0Req, auth0Err = updateUserAuth0Req(dgraphReqJSON, auth0UpdateURL)

			updateUserVariables := map[string]interface{}{
				"filter": map[string]interface{}{
					"auth0ID": map[string]interface{}{
						"eq": *dgraphReqJSON.Auth0ID,
					},
				},
			}

			if dgraphReqJSON.Role != nil {
				updateUserVariables["filter"] = map[string]interface{}{
					"auth0ID": map[string]interface{}{
						"eq": *dgraphReqJSON.Auth0ID,
					},
				}
				updateUserVariables["set"] = map[string]interface{}{
					"role": *dgraphReqJSON.Role,
				}
			}

			dgraphMutation = graphql.UpdateUserMutation
			dgraphVariables = updateUserVariables
		} else if r.Method == http.MethodDelete {
			auth0DeleteURL := auth0URL + "users/" + url.PathEscape(*dgraphReqJSON.Auth0ID)
			auth0Req, auth0Err = deleteUserAuth0Req(auth0DeleteURL)

			dgraphMutation = graphql.DeleteUserMutation
			dgraphVariables = map[string]interface{}{
				"auth0ID": map[string]interface{}{
					"eq": *dgraphReqJSON.Auth0ID,
				},
			}
		} else {
			http.Error(w, handlers.ErrIncorrectHTTPMethod, http.StatusBadRequest)
			return
		}

		if auth0Err != nil {
			http.Error(w, handlers.ErrCreatingAuth0Request, http.StatusInternalServerError)
			return
		}

		auth0Req.Header.Set("Authorization", "Bearer "+managementToken)
		auth0Req.Header.Set("Content-Type", "application/json")

		httpClient := &http.Client{}

		auth0Resp, err := httpClient.Do(auth0Req)
		if err != nil || !checkSuccess(auth0Resp.StatusCode) {
			http.Error(w, handlers.ErrExecutingAuth0Request, http.StatusInternalServerError)
			return
		}
		defer auth0Resp.Body.Close()

		var auth0RespJSON handlers.Auth0Response
		if r.Method == http.MethodPost {
			if err := json.NewDecoder(auth0Resp.Body).Decode(&auth0RespJSON); err != nil {
				http.Error(w, handlers.ErrUnmarshallingResponseBody, http.StatusBadRequest)
				return
			}

			dgraphVariables = createUserDgraphReq(dgraphReqJSON, auth0RespJSON.Auth0ID)
		}

		dgraphClient := graphql.New(
			httpClient,
			dgraphURL,
			r.Header.Get("X-Auth0-Token"),
		)

		_, err = dgraphClient.SendRequest(dgraphMutation, dgraphVariables)
		if err != nil {
			http.Error(w, handlers.ErrDgraphMutation, http.StatusBadRequest)
			return
		}

		// response payloads ensure non-error message from dgraph
		if r.Method == http.MethodPost {
			dgraphResponseBytes, err := json.Marshal(dgraphVariables.([]map[string]interface{})[0])
			if err != nil {
				http.Error(w, handlers.ErrMarshallingDgraphJSON, http.StatusBadRequest)
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

func createUserAuth0Req(dgraphReqJSON handlers.DgraphRequest, url string) (*http.Request, error) {
	createUserJSON := handlers.CreateUserRequest{
		Email:    *dgraphReqJSON.Email,
		Password: *dgraphReqJSON.Password,
		AppMetadata: handlers.AppMetadata{
			Role:  dgraphReqJSON.Role,
			OrgID: dgraphReqJSON.Owner,
		},
		FirstName:  *dgraphReqJSON.FirstName,
		LastName:   *dgraphReqJSON.LastName,
		Connection: "Username-Password-Authentication",
	}

	createUserByte, err := json.Marshal(createUserJSON)
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, url, bytes.NewReader(createUserByte))
}

func updateUserAuth0Req(dgraphReqJSON handlers.DgraphRequest, url string) (*http.Request, error) {
	updateUserJSON := handlers.UpdateUserRequest{}
	if dgraphReqJSON.Password != nil {
		updateUserJSON.Password = dgraphReqJSON.Password
	}
	if dgraphReqJSON.Role != nil {
		updateUserJSON.AppMetadata = handlers.AppMetadata{
			Role: dgraphReqJSON.Role,
		}
	}

	updateUserByte, err := json.Marshal(updateUserJSON)
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPatch, url, bytes.NewReader(updateUserByte))
}

func deleteUserAuth0Req(url string) (*http.Request, error) {
	return http.NewRequest(http.MethodDelete, url, nil)
}

func checkSuccess(status int) bool {
	return status == http.StatusOK || status == http.StatusCreated || status == http.StatusNoContent
}

func createUserDgraphReq(dgraphReqJSON handlers.DgraphRequest, auth0ID string) []map[string]interface{} {
	return []map[string]interface{}{
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
}
