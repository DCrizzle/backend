package main

import (
	"fmt" // TEMP
	"os"

	"github.com/forstmeier/tbd/api"
)

func main() {
	addr := os.Getenv("OPENTORY_ADDR")
	fmt.Println("addr:", addr) // TEMP

	server, err := api.NewServer(addr)
	if err != nil {
		panic(err)
	}

	err = server.Start()
	fmt.Println("err:", err) // TEMP
}
