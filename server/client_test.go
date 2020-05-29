package server

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func Test_graphQLClient(t *testing.T) {
	responseData := `{"data":"key":"value"}`

	gqlc := newGraphQLClient(&http.Client{
		Timeout: 10 * time.Second,
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, responseData)
	})

	server := httptest.NewServer(mux)

	body := strings.NewReader(`{"query":"queryData { key }"}`)

	url := server.URL + "/graphql"
	response, err := gqlc.post(url, "token", body)
	if err != nil {
		t.Fatal(err)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	bodyString := buf.String()

	if bodyString != responseData {
		t.Errorf("response data received: %s, expected: %s", bodyString, responseData)
	}
}
