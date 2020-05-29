package server

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gorilla/csrf"
)

const errMutationRequest = "server: error invoking backend database mutation"

func graphQLHandler(dgraphURL string, gql graphQL) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Auth0-Token")

		response, err := gql.post(dgraphURL, token, r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"message":"%s"}`, errMutationRequest), http.StatusInternalServerError)
			return
		}

		w.Header().Set("X-CSRF-Token", csrf.Token(r))

		buffer := new(bytes.Buffer)
		buffer.ReadFrom(response.Body)
		fmt.Fprint(w, buffer.String())
	})
}
