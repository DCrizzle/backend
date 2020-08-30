package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// NOTE: move these consts into config.json
const (
	auth0TokenURL = "https://folivora.us.auth0.com/oauth/token"
	// NOTE: possibly remove trailing "/" for uniformity but this would require changing url processing logic/tests
	auth0APIURL = "https://folivora.us.auth0.com/api/v2"
	dgraphURL   = "http://localhost:8080/graphql"
)

func main() {
	cfg, err := readConfig()
	if err != nil {
		log.Fatal("error reading config:", err.Error())
	}

	token, err := getAuth0APIToken(auth0TokenURL, cfg.Auth0)
	if err != nil {
		log.Fatal("error getting auth0 management api token:", err.Error())
	}

	handler := usersHandler(cfg.Folivora.HelperSecret, token, auth0APIURL, dgraphURL)
	server := newServer(handler)

	ctx := context.Background()

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go server.start()
	go func() {
		<-sigs
		done <- true
	}()

	<-done
	server.stop(ctx)
}
