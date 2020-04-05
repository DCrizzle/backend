package main

import (
	"os"

	"github.com/forstmeier/tbd/api"
)

func main() {
	addr := os.Getenv("OPENTORY_ADDR")

	server, err := api.NewServer(addr)
	if err != nil {
		panic(err)
	}

	_ = server
}
