// +build !mock

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log" // TEMP
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
		log.Println("secret:", secret)

		var dgraphReqJSON dgraphRequest
		if err := json.NewDecoder(r.Body).Decode(&dgraphReqJSON); err != nil {
			http.Error(w, errIncorrectRequestBody, http.StatusBadRequest)
			return
		}
		log.Printf("dgraph request json: %+v\n", dgraphReqJSON)

		var auth0Req *http.Request
		var auth0ReqErr error
		if r.Method == http.MethodPost {
			log.Println("post method")
			createUserJSON := createUserRequest{
				Email:    dgraphReqJSON.Email,
				Password: dgraphReqJSON.Password,
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
			log.Println("create user request:", string(createUserByte))

			auth0Req, auth0ReqErr = http.NewRequest(
				http.MethodPost,
				url+"users",
				bytes.NewReader(createUserByte),
			)
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

			auth0Req, auth0ReqErr = http.NewRequest(
				http.MethodPatch,
				url+"users/"+dgraphReqJSON.Auth0ID,
				bytes.NewReader(updateUserByte),
			)
		} else if r.Method == http.MethodDelete {
			auth0Req, auth0ReqErr = http.NewRequest(
				http.MethodDelete,
				url+"users/"+dgraphReqJSON.Auth0ID,
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
		auth0Req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}

		auth0Resp, err := client.Do(auth0Req)
		defer auth0Resp.Body.Close()
		if err != nil || !checkSuccess(auth0Resp.StatusCode) {
			http.Error(w, errExecutingAuth0Request, http.StatusInternalServerError)
			return
		}

		if r.Method == http.MethodDelete {
			fmt.Fprintf(w, fmt.Sprintf(`{"message": "success", "auth0ID": "%s"}`, dgraphReqJSON.Auth0ID))
		} else {
			var auth0RespJSON auth0Response
			if err := json.NewDecoder(auth0Resp.Body).Decode(&auth0RespJSON); err != nil {
				http.Error(w, errIncorrectResponseBody, http.StatusBadRequest)
				return
			}
			fmt.Fprintf(w, fmt.Sprintf(`{"message": "success", "auth0ID": "%s"}`, auth0RespJSON.Auth0ID))
		}
	})
}

func checkSuccess(status int) bool {
	return status == http.StatusOK || status == http.StatusCreated || status == http.StatusNoContent
}

type dgraphRequest struct {
	Auth0ID   string `json:"auth0ID"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	OrgID     string `json:"orgID"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
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
	OrgID string `json:"orgID,omitempty"`
}

type auth0Response struct {
	Auth0ID string `json:"user_id"`
}
