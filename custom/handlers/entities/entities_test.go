package entities

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/forstmeier/backend/custom/handlers"
)

type mockClassifier struct {
	mockClassifierResp map[string][]string
	mockClassifierErr  error
}

func (mc *mockClassifier) ClassifyEntities(blob string, docType string) (map[string][]string, error) {
	return mc.mockClassifierResp, mc.mockClassifierErr
}

func TestHandler(t *testing.T) {
	tests := []struct {
		description        string
		entitiesResp       map[string][]string
		entitiesErr        error
		requestBody        string
		requestSecret      string
		customSecret       string
		responseStatusCode int
		responseBody       string
	}{
		{
			description:        "incorrect secret provided in request to custom",
			entitiesResp:       nil,
			entitiesErr:        nil,
			requestBody:        "",
			requestSecret:      "incorrect_secret",
			customSecret:       "correct_secret",
			responseStatusCode: http.StatusBadRequest,
			responseBody:       handlers.ErrIncorrectSecret,
		},
		{
			description:        "invalid json body received in request to custom",
			entitiesResp:       nil,
			entitiesErr:        nil,
			requestBody:        "",
			requestSecret:      "correct_secret",
			customSecret:       "correct_secret",
			responseStatusCode: http.StatusBadRequest,
			responseBody:       handlers.ErrIncorrectRequestBody,
		},
		{
			description:        "error received in response from classify entities",
			entitiesResp:       nil,
			entitiesErr:        errors.New("mock classify entities error"),
			requestBody:        `{"owner": "", "form": "", "docType": "", "blob": ""}`,
			requestSecret:      "correct_secret",
			customSecret:       "correct_secret",
			responseStatusCode: http.StatusInternalServerError,
			responseBody:       handlers.ErrClassifyingEntities,
		},
		{
			description: "successful classify entities invocation",
			entitiesResp: map[string][]string{
				"person": []string{"test_person"},
			},
			entitiesErr:        nil,
			requestBody:        `{"owner": "", "form": "", "docType": "", "blob": ""}`,
			requestSecret:      "correct_secret",
			customSecret:       "correct_secret",
			responseStatusCode: http.StatusOK,
			responseBody:       `{"owner": {"id": ""}, "form": {"id": ""}, "person": []}`,
		},
	}

	for _, test := range tests {
		mc := &mockClassifier{
			mockClassifierResp: test.entitiesResp,
			mockClassifierErr:  test.entitiesErr,
		}

		dgraphMux := http.NewServeMux()
		dgraphMux.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		dgraphServer := httptest.NewServer(dgraphMux)
		dgraphURL := dgraphServer.URL + "/graphql"

		t.Run(test.description, func(t *testing.T) {
			req, err := http.NewRequest(
				http.MethodPost,
				"/nlp/entities",
				bytes.NewReader([]byte(test.requestBody)),
			)
			if err != nil {
				t.Fatal("error creating request:", err.Error())
			}
			req.Header.Set("folivora-custom-secret", test.requestSecret)

			rec := httptest.NewRecorder()

			handler := http.HandlerFunc(Handler(test.customSecret, dgraphURL, mc))

			handler.ServeHTTP(rec, req)

			if status := rec.Code; status != test.responseStatusCode {
				t.Errorf("status received: %d, expected: %d", status, test.responseStatusCode)
			}

			if body := strings.TrimSuffix(rec.Body.String(), "\n"); body != test.responseBody {
				t.Errorf("body received: %s, expected: %s", body, test.responseBody)
			}
		})
	}
}
