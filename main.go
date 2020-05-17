package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/forstmeier/backend/server"
)

func main() {
	ctx := context.Background()

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	backendServer, err := server.NewServer("config.json", &server.GraphQLClient{})
	if err != nil {
		panic(err)
	}

	go backendServer.Start()
	go func() {
		<-sigs
		done <- true
	}()

	<-done
	backendServer.Stop(ctx)
}
