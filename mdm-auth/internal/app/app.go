package app

import (
	"context"
	"errors"
	"log"
	"mdm-auth/config"
	"mdm-auth/internal/delivery/rest"
	"mdm-auth/internal/providers"
	"mdm-auth/internal/repositories"
	"mdm-auth/internal/services"
	db "mdm-auth/internal/utils/database"
	"mdm-auth/internal/utils/logging"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartApp() {
	cfg, err := config.LoadLocalConfig()
	if err != nil {
		log.Fatal("Could not get configs", err.Error())
	}

	authDb, err := db.InitPostgresDb(cfg.DbURL)
	if err != nil {
		log.Fatal("Could not init postgresql connection", err.Error())
	}

	repositories := repositories.NewRepositories(authDb)

	providers, err := providers.NewProviders(cfg.Keycloak)
	if err != nil {
		log.Fatal("could initialize providers", err.Error())
	}
	services := services.NewServices(&services.Deps{
		Providers:    providers,
		Repositories: repositories,
	})

	handlers := rest.NewHandler(services)

	srv := rest.NewServer(cfg, handlers.Init(cfg))

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logging.Info("Server started")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second
	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logging.Error("failed to stop server:", err.Error())
	}
}
