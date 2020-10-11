// +build mock

package entities

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/forstmeier/backend/custom/handlers"
	"github.com/forstmeier/backend/graphql"
	entint "github.com/forstmeier/internal/nlp/entities"
)

func Handler(internalSecret, dgraphURL string, classifier entint.Classifier) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var dgraphReqJSON handlers.DgraphRequest
		if err := json.NewDecoder(r.Body).Decode(&dgraphReqJSON); err != nil {
			http.Error(w, errorIncorrectRequestBody, http.StatusBadRequest)
			return
		}

		dgraphClient := graphql.New(
			&http.Client{},
			dgraphURL,
			r.Header.Get("X-Auth0-Token"),
		)

		dgraphVariables := map[string]interface{}{
			"owner": dgraphReqJSON.Owner,
			"form":  dgraphReqJSON.Form,
		}

		_, err = dgraphClient.SendRequest(graphql.AddEntitiesMutation, dgraphVariables)
		if err != nil {
			http.Error(w, errorDgraphMutation, http.StatusInternalServerError)
			return
		}

		responseBody := fmt.Sprintf(`{"owner": {"id": ""}, "form": {"id": ""}, "person": []}`)
		fmt.Fprintf(w, responseBody)
	})
}
