package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	auth0TokenURL = "https://folivora.us.auth0.com/oauth/token"
	auth0APIURL   = "AUTH0_URL"
)

func main() {
	server := &server{}
	if PROD {
		cfg, err := readConfig()
		if err != nil {
			log.Fatal("error reading config:", err.Error())
		}

		token, err := getAuth0APIToken(auth0TokenURL, cfg.Auth0)
		if err != nil {
			log.Fatal("error getting auth0 management api token:", err.Error())
		}

		handler := usersHandler(cfg.Folivora.HelperSecret, token, auth0APIURL)
		server = newServer(handler)
	} else {
		server = newServer(mockUsersHandler)
	}

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
