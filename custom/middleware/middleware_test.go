package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"

	"github.com/forstmeier/backend/custom/handlers"
)

func testHandler(w http.ResponseWriter, r *http.Request) {}

func TestMiddlware(t *testing.T) {
	tests := []struct {
		description        string
		requestSecret      string
		responseStatusCode int
		responseBody       string
	}{
		{
			description:        "incorrect request secret provided",
			requestSecret:      "incorrect_secret",
			responseStatusCode: http.StatusBadRequest,
			responseBody:       handlers.ErrIncorrectSecret,
		},
		{
			description:        "correct request secret provided",
			requestSecret:      "correct_secret",
			responseStatusCode: http.StatusOK,
			responseBody:       "",
		},
	}

	testRouter := mux.NewRouter()
	testRouter.HandleFunc("/", testHandler)

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			mw := New("correct_secret")

			testRouter.Use(mw.Middleware)

			req, err := http.NewRequest(
				http.MethodGet,
				"/",
				nil,
			)
			if err != nil {
				t.Fatalf("error creating request: %s\n", err.Error())
			}
			req.Header.Set("folivora-custom-secret", test.requestSecret)

			rec := httptest.NewRecorder()

			testRouter.ServeHTTP(rec, req)

			if status := rec.Code; status != test.responseStatusCode {
				t.Errorf("status received: %d, expected: %d", status, test.responseStatusCode)
			}

			if body := strings.TrimSuffix(rec.Body.String(), "\n"); body != test.responseBody {
				t.Errorf("body received: %s, expected: %s", body, test.responseBody)
			}
		})
	}
}
