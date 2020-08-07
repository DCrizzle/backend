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
	errIncorrectRequestBody  = "incorrect request body received"
	errIncorrectHTTPMethod   = "unsupported http method received"
	errMarshallingCreateJSON = "error marshalling create json"
	errMarshallingUpdateJSON = "error marshalling update json"
	errCreatingAuth0Request  = "error creating auth0 request"
	errExecutingAuth0Request = "error executing auth0 request"
	errIncorrectResponseBody = "incorrect response body received"
)

func usersHandler(secret, token, url string) http.HandlerFunc {
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
		var auth0ReqErr error
		if r.Method == http.MethodPost {
			createUserJSON := createUser{
				Email: dgraphReqJSON.Email,
				AppMetadata: appMetadata{
					Role:  dgraphReqJSON.Role,
					OrgID: dgraphReqJSON.OrgID,
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

			auth0Req, auth0ReqErr = http.NewRequest(
				http.MethodPost,
				url,
				bytes.NewReader(createUserByte),
			)

		} else if r.Method == http.MethodPatch {
			updateUserJSON := updateUser{
				AppMetadata: appMetadata{
					Role: dgraphReqJSON.Role,
				},
			}

			updateUserByte, err := json.Marshal(updateUserJSON)
			if err != nil {
				http.Error(w, errMarshallingUpdateJSON, http.StatusInternalServerError)
				return
			}

			auth0Req, auth0ReqErr = http.NewRequest(
				http.MethodPatch,
				url+"/"+dgraphReqJSON.UserID,
				bytes.NewReader(updateUserByte),
			)

		} else if r.Method == http.MethodDelete {
			auth0Req, auth0ReqErr = http.NewRequest(
				http.MethodDelete,
				url+"/"+dgraphReqJSON.UserID,
				nil,
			)
		} else {
			http.Error(w, errIncorrectHTTPMethod, http.StatusBadRequest)
			return
		}

		if auth0ReqErr != nil {
			http.Error(w, errCreatingAuth0Request, http.StatusInternalServerError)
			return
		}

		auth0Req.Header.Set("Authorization", "Bearer "+token)

		client := &http.Client{}

		auth0Resp, err := client.Do(auth0Req)
		defer auth0Resp.Body.Close()
		if err != nil || !checkSuccess(auth0Resp.StatusCode) {
			http.Error(w, errExecutingAuth0Request, http.StatusInternalServerError)
			return
		}

		var auth0RespJSON auth0Response
		if err := json.NewDecoder(auth0Resp.Body).Decode(&auth0RespJSON); err != nil {
			http.Error(w, errIncorrectResponseBody, http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, fmt.Sprintf(`{"message": "success", "user_id": "%s"}`, auth0RespJSON.UserID))
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

type auth0Response struct {
	UserID string `json:"user_id"`
}
