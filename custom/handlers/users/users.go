// +build !mock

package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	internal "github.com/forstmeier/internal/auth0/users"

	"github.com/forstmeier/backend/custom/handlers"
	"github.com/forstmeier/backend/graphql"
)

// Handler is an HTTP listener for User type @custom directive events.
//
// createUser: adds a user to Auth0 and to Dgraph with the Auth0 ID field
// editUser: updates an Auth0 user role or password in Auth0 and Dgraph
// removeUser: deletes an Auth0 user from Auth0 and Dgraph
func Handler(dgraphURL string, client internal.Client) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var dgraphReqJSON handlers.DgraphRequest
		if err := json.NewDecoder(r.Body).Decode(&dgraphReqJSON); err != nil {
			http.Error(w, handlers.ErrIncorrectRequestBody, http.StatusBadRequest)
			return
		}

		var resp *internal.Response

		var dgraphMutation string
		var dgraphVariables interface{}

		if r.Method == http.MethodPost {
			createRequest := internal.Request{
				Owner:     dgraphReqJSON.Owner,
				Email:     dgraphReqJSON.Email,
				Password:  dgraphReqJSON.Password,
				Role:      dgraphReqJSON.Role,
				Org:       dgraphReqJSON.Org,
				FirstName: dgraphReqJSON.FirstName,
				LastName:  dgraphReqJSON.LastName,
			}

			createResp, err := client.CreateUser(createRequest)
			if err != nil {
				http.Error(w, handlers.ErrIncorrectRequestBody, http.StatusBadRequest)
				return
			}
			resp = createResp

			dgraphMutation = graphql.AddUsersMutation
			dgraphVariables = createUserDgraphReq(dgraphReqJSON, resp.Auth0ID)
		} else if r.Method == http.MethodPatch {
			updateRequest := internal.Request{
				Auth0ID:  dgraphReqJSON.Auth0ID,
				Owner:    dgraphReqJSON.Owner,
				Password: dgraphReqJSON.Password,
				Role:     dgraphReqJSON.Role,
			}

			updateResp, err := client.UpdateUser(updateRequest)
			if err != nil {
				// outline:
				// [ ] handle error
			}
			resp = updateResp

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
			deleteRequest := internal.Request{
				Auth0ID: dgraphReqJSON.Auth0ID,
			}

			deleteResp, err := client.DeleteUser(deleteRequest)
			if err != nil {
				// outline:
				// [ ] handle error
			}
			resp = deleteResp

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

		httpClient := &http.Client{}

		dgraphClient := graphql.New(
			httpClient,
			dgraphURL,
			r.Header.Get("X-Auth0-Token"),
		)

		_, err := dgraphClient.SendRequest(dgraphMutation, dgraphVariables)
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
