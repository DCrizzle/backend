package server

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

type mockGraphQL struct {
	mutationResponse *http.Response
	mutationError    error
	queryResponse    *http.Response
	queryError       error
}

func (mgql *mockGraphQL) mutation(url string, body io.Reader) (*http.Response, error) {
	return mgql.mutationResponse, mgql.mutationError
}

func (mgql *mockGraphQL) query(url string) (*http.Response, error) {
	return mgql.queryResponse, mgql.queryError
}

func TestServer(t *testing.T) {
	mgql := &mockGraphQL{}

	configFile, err := ioutil.TempFile("", "ServerConfig*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(configFile.Name())

	_, err = configFile.Write([]byte(`
		{
			"auth0_api_audience":"audience",
			"auth0_api_client_secret": "secret",
			"auth0_domain": "domain",
			"csrf_key": "key",
			"dgraph_url": "url"
		}
	`))
	if err != nil {
		t.Fatal(err)
	}

	server, err := NewServer(configFile.Name(), mgql)

	if server == nil && err != nil {
		t.Error("nil received creating new server")
	}

	go server.Start()

	server.Stop(context.Background())
}

func Test_graphQLClient(t *testing.T) {
	responseData := `{"data":"key":"value"}`

	gqlc := &GraphQLClient{}

	mux := http.NewServeMux()
	mux.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, responseData)
	})

	server := httptest.NewServer(mux)

	t.Run("test mutation implementation", func(t *testing.T) {
		body := strings.NewReader(`query getData{ key }`)

		response, err := gqlc.mutation(server.URL+"/graphql", body)
		if err != nil {
			t.Fatal(err)
		}

		buf := new(bytes.Buffer)
		buf.ReadFrom(response.Body)
		bodyString := buf.String()

		if bodyString != responseData {
			t.Errorf("response data received: %s, expected: %s", bodyString, responseData)
		}
	})

	t.Run("test query implementation", func(t *testing.T) {
		response, err := gqlc.query(server.URL + "/graphql?query%20getData%7B%20key%20%7D")
		if err != nil {
			t.Fatal(err)
		}

		buf := new(bytes.Buffer)
		buf.ReadFrom(response.Body)
		bodyString := buf.String()

		if bodyString != responseData {
			t.Errorf("response data received: %s, expected: %s", bodyString, responseData)
		}
	})
}
