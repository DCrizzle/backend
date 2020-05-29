package server

import (
	"context"
	"io/ioutil"
	"os"
	"testing"
)

func TestServer(t *testing.T) {
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

	server, err := NewServer(configFile.Name())

	if server == nil && err != nil {
		t.Error("nil received creating new server")
	}

	go server.Start()

	server.Stop(context.Background())
}
