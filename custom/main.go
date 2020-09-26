package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/forstmeier/backend/auth"
	"github.com/forstmeier/backend/config"
)

func main() {
	configPath := flag.String("config", "../etc/config/config.json", "path to config json file")

	flag.Parse()

	cfg, err := config.New(*configPath)
	if err != nil {
		log.Fatal("error reading config:", err.Error())
	}

	ac := auth.New(cfg)

	managementToken, err := ac.GetManagementAPIToken()
	if err != nil {
		log.Fatal("error getting auth0 management api token:", err.Error())
	}

	handler := usersHandler(
		cfg.Folivora.HelperSecret,
		managementToken,
		cfg.Auth0.AudienceURL, // same as the api url
		cfg.Folivora.DgraphURL,
	)
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
