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
	"github.com/forstmeier/backend/custom/handlers/users"
	"github.com/forstmeier/backend/custom/server"
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

	handler := users.Handler(
		cfg.Folivora.CustomSecret,
		managementToken,
		cfg.Auth0.AudienceURL, // same as the api url
		cfg.Folivora.DgraphURL,
	)
	customServer := server.New(handler)

	ctx := context.Background()

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go customServer.Start()
	go func() {
		<-sigs
		done <- true
	}()

	<-done
	customServer.Stop(ctx)
}
