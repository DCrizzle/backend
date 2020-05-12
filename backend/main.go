package main

import (
	"context"
	"flag"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/forstmeier/tbd/backend/server"
)

func main() {
	ctx := context.Background()

	schemaPath := flag.String("schema", "", "target database schema file path")
	databaseURL := flag.String("database", "", "target dgraph database url")

	flag.Parse()

	cmd := exec.Command("/bin/sh", "database/start.sh", *schemaPath)
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	backendServer := server.NewServer(*databaseURL, &server.GraphQLClient{})

	go backendServer.Start()
	go func() {
		<-sigs
		done <- true
	}()

	<-done
	backendServer.Stop(ctx)
}
