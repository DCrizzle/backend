// +build mock

package entities

import (
	"net/http"

	entint "github.com/forstmeier/internal/nlp/entities"
)

func Handler(folivoraSecret, internalSecret, dgraphURL string, classifier entint.Classifier) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if folivoraSecret != r.Header.Get("folivora-custom-secret") {
			http.Error(w, handlers.ErrIncorrectSecret, http.StatusBadRequest)
			return
		}

		var dgraphReqJSON handlers.DgraphRequest
		if err := json.NewDecoder(r.Body).Decode(&dgraphReqJSON); err != nil {
			http.Error(w, handlers.ErrIncorrectRequestBody, http.StatusBadRequest)
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
			http.Error(w, handlers.ErrDgraphMutation, http.StatusInternalServerError)
			return
		}

		responseBody := fmt.Sprintf(`{"owner": {"id": ""}, "form": {"id": ""}, "person": []}`)
		fmt.Fprintf(w, responseBody)
	}
}
