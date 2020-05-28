package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_middleware(t *testing.T) {

}

func Test_getPEMCert(t *testing.T) {

	tests := []struct {
		description string
	}{
		{},
	}

	for _, test := range tests {

		t.Run(test.description, func(t *testing.T) {

			fmt.Println("test:", test)

			mux := http.NewServeMux()
			mux.HandleFunc("/.well-known/jwks.json", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, `{"test":"stuff"}`)
			})

			server := httptest.NewServer(mux)

			url := server.URL
			if err != nil {
				t.Fatal("error parsing test url:", err)
			}

			t.Error("url:", url)

		})
	}
}
