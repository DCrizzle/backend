package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	server := &server{}
	if PROD {
		client := &http.Client{}
		token := getAuth0ManagementAPIToken(client)
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
