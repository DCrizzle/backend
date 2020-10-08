package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	// entint "github.com/forstmeier/internal/nlp/entities"

	"github.com/forstmeier/backend/auth"
	"github.com/forstmeier/backend/config"
	"github.com/forstmeier/backend/custom/handlers/entities"
	"github.com/forstmeier/backend/custom/handlers/users"
	"github.com/forstmeier/backend/custom/middleware"
	"github.com/forstmeier/backend/custom/server"
)

// NOTE: this is temporary while there are no trained models available
type mockClassifier struct{}

func (mc *mockClassifier) ClassifyEntities(blob string, docType string) (map[string][]string, error) {
	return map[string][]string{}, nil
}

func main() {
	configPath := flag.String("config", "../etc/config/config.json", "path to config json file")

	flag.Parse()

	cfg, err := config.New(*configPath)
	if err != nil {
		log.Fatalf("error reading config: %s\n", err.Error())
	}

	ac := auth.New(cfg)

	managementToken, err := ac.GetManagementAPIToken()
	if err != nil {
		log.Fatalf("error getting auth0 management api token: %s\n", err.Error())
	}

	root := middleware.New(cfg.Folivora.CustomSecret)

	userHandler := users.Handler(
		managementToken,
		cfg.Auth0.AudienceURL, // same as the api url
		cfg.Folivora.DgraphURL,
	)

	// classifier, err := pkgent.New("custom/handlers/entities/config.json", "custom/handlers/entities/models/")
	mc := &mockClassifier{}
	if err != nil {
		log.Fatalf("error creating entities classifier: %s\n", err.Error())
	}

	entitiesHandler := entities.Handler(
		cfg.Folivora.DgraphURL,
		mc,
	)

	customServer := server.New(root.Middleware, userHandler, entitiesHandler)

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
