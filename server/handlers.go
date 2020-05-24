package server

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/csrf"
)

func graphQLHandler(dgraphURL string, gql graphQL) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var output io.ReadCloser

		if r.Method == http.MethodPost {
			response, err := gql.mutation(dgraphURL, r.Body)
			if err != nil {
				http.Error(w, fmt.Sprintf(`{"message":"%s"}`, errMutationRequest), http.StatusInternalServerError)
				return
			}

			output = response.Body

		} else if r.Method == http.MethodGet {
			queryURL := dgraphURL + r.URL.Path + "?" + r.URL.RawQuery
			response, err := gql.query(queryURL)
			if err != nil {
				http.Error(w, fmt.Sprintf(`{"message":"%s"}`, errQueryRequest), http.StatusInternalServerError)
				return
			}

			output = response.Body
		}

		w.Header().Set("X-CSRF-Token", csrf.Token(r))

		buffer := new(bytes.Buffer)
		buffer.ReadFrom(output)
		fmt.Fprint(w, buffer.String())
	})
}
