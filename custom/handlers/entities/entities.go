package entities

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/forstmeier/backend/custom/handlers"
	"github.com/forstmeier/backend/graphql"
)

// Handler is an HTTP listener for classify entity @custom directive events
func Handler(folivoraSecret, internalSecret, entitiesURL, dgraphURL string) http.HandlerFunc {
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

		payload, err := json.Marshal(map[string]string{
			"blob":    dgraphReqJSON.Blob,
			"docType": dgraphReqJSON.DocType,
		})
		if err != nil {
			http.Error(w, handlers.ErrMarshallingEntitiesJSON, http.StatusInternalServerError)
			return
		}

		entitiesReq, err := http.NewRequest(http.MethodPost, entitiesURL, bytes.NewReader(payload))
		if err != nil {
			http.Error(w, handlers.ErrCreatingEntitiesRequest, http.StatusInternalServerError)
			return
		}

		entitiesReq.Header.Set("folivora-internal-secret", internalSecret)
		entitiesReq.Header.Set("Content-Type", "application/json")

		httpClient := &http.Client{}

		entitiesResp, err := httpClient.Do(entitiesReq)
		if entitiesResp.StatusCode != http.StatusOK {
			http.Error(w, handlers.ErrExecutingEntitiesRequest, http.StatusInternalServerError)
			return
		}
		defer entitiesResp.Body.Close()

		entitiesRespJSON := handlers.EntitiesResponse{}
		if err := json.NewDecoder(entitiesResp.Body).Decode(&entitiesRespJSON); err != nil {
			http.Error(w, handlers.ErrUnmarshallingResponseBody, http.StatusInternalServerError)
			return
		}

		dgraphClient := graphql.New(
			httpClient,
			dgraphURL,
			r.Header.Get("X-Auth0-Token"),
		)

		dgraphVariables := map[string]interface{}{
			"owner": dgraphReqJSON.Owner,
			"form":  dgraphReqJSON.Form,
		}

		entitiesData := entitiesRespJSON.Data
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
