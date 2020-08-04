package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const auth0URL = "https://folivora.us.auth0.com/oauth/token"

func main() {
	server := &server{}
	if PROD {
		token, err := getAuth0ManagementAPIToken(auth0URL)
		if err != nil {
			log.Fatal("error getting auth0 management api token:", err.Error())
		}

		server = newServer(token, usersHandler)
	} else {
		token := "MOCK_AUTH0_MANAGEMENT_API_TOKEN"
		server = newServer(token, mockUsersHandler)
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
