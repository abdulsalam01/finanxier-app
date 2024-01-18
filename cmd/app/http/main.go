package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/api-sekejap/cmd/app"
	"github.com/api-sekejap/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/sirupsen/logrus"
)

const (
	configPath = "./config/manager"
)

func main() {
	ctx := context.Background()

	logrus.Info("Load config manager")
	configs, err := config.NewConfigManager(configPath)
	if err != nil {
		logrus.Errorf("Error when loading config file %v", err)
		return
	}

	// Init initializer.
	logrus.Info("Load initializer helper")
	baseInitializer, err := app.Initializer(ctx, configs)

	// Init Schema migrations.
	logrus.Warnf("Run schema migrations on %s", baseInitializer.Database.Config().ConnConfig.Database)
	err = config.SchemaMigrate(baseInitializer.Database.Config().ConnString(), app.DatabaseVersion)
	if err != nil {
		logrus.Errorf("Error when setup migrations %v", err)
		return
	}

	// Init routes.
	logrus.Info("Setup routes")
	routes := initializeRoutes(ctx)

	// Start server.
	logrus.Infof("Server start at port :%s", configs.App.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", configs.App.Port), routes)
}

func initializeRoutes(
	ctx context.Context,
) *chi.Mux {
	r := chi.NewMux()

	// Middleware section.
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers.
	}))

	// Endpoint section.
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("welcome"))
	})

	return r
}
