// +build !mock

package entities

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/forstmeier/backend/custom/handlers"
	"github.com/forstmeier/backend/graphql"

	entint "github.com/forstmeier/internal/nlp/entities"
)

// Handler is an HTTP listener for classify entity @custom directive events
func Handler(folivoraSecret, internalSecret, dgraphURL string, classifier entint.Classifier) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if folivoraSecret != r.Header.Get("folivora-custom-secret") {
			http.Error(w, handlers.ErrIncorrectSecret, http.StatusBadRequest)
			return
		}

		var dgraphReqJSON handlers.DgraphEntitiesRequest
		if err := json.NewDecoder(r.Body).Decode(&dgraphReqJSON); err != nil {
			http.Error(w, handlers.ErrIncorrectRequestBody, http.StatusBadRequest)
			return
		}

		entitiesData, err := classifier.ClassifyEntities(dgraphReqJSON.Blob, dgraphReqJSON.DocType)
		if err != nil {
			http.Error(w, handlers.ErrClassifyingEntities, http.StatusInternalServerError)
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

		if _, ok := entitiesData["person"]; ok {
			dgraphVariables["person"] = entitiesData["person"]
		}

		_, err = dgraphClient.SendRequest(graphql.AddEntitiesMutation, dgraphVariables)
		if err != nil {
			http.Error(w, handlers.ErrDgraphMutation, http.StatusInternalServerError)
			return
		}

		responseBody := fmt.Sprintf(`{"owner": {"id": ""}, "form": {"id": ""}, "person": []}`)
		fmt.Fprintf(w, responseBody)
	})
}
