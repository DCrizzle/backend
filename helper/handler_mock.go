// +build mock

package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

const errIncorrectRequestBody = "incorrect request body received"

func usersHandler(secret, token, auth0URL, dgraphURL string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var dgraphReqJSON dgraphRequest
		if err := json.NewDecoder(r.Body).Decode(&dgraphReqJSON); err != nil {
			http.Error(w, errIncorrectRequestBody, http.StatusBadRequest)
			return
		}

		auth0ID := ""
		if r.Method == http.MethodPost {
			id := uuid.New().String()
			hexID := hex.EncodeToString([]byte(id))
			auth0ID = "auth0|" + hexID
		} else if r.Method == http.MethodPatch {
			auth0ID = dgraphReqJSON.Auth0ID
		} else if r.Method == http.MethodDelete {
			auth0ID = dgraphReqJSON.Auth0ID
		}

		fmt.Fprintf(w, fmt.Sprintf(`{"message": "success", "auth0ID": "%s"}`, auth0ID))
	})
}

type dgraphRequest struct {
	Auth0ID   string `json:"auth0ID"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	OrgID     string `json:"orgID"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
