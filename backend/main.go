package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/forstmeier/tbd/backend/server"
)

func main() {
	ctx := context.Background()

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	backendServer := server.NewServer("http://127.0.0.1:8080/graphql", &server.GraphQLClient{})

	go backendServer.Start()
	go func() {
		<-sigs
		done <- true
	}()

	<-done
	backendServer.Stop(ctx)
}
