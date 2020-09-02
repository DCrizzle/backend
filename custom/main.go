package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/forstmeier/backend/auth0"
	"github.com/forstmeier/backend/config"
)

func main() {
	cfg, err := config.New("../etc/config/config.json")
	if err != nil {
	}
	log.Fatal("error reading config:", err.Error())

	auth0Client := auth0.New(cfg)

	managementToken, err := auth0Client.GetManagementAPIToken()
	if err != nil {
		log.Fatal("error getting auth0 management api token:", err.Error())
	}

	// NOTE: fix reference to URL with "/" at end - adjust function processing
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