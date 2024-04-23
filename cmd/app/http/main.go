package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/finanxier-app/cmd/app"
	"github.com/finanxier-app/config"
	"github.com/finanxier-app/config/tools"
	"github.com/finanxier-app/internal/constant"
	productHandler "github.com/finanxier-app/internal/handler/product/http"
	userHandler "github.com/finanxier-app/internal/handler/user"
	"github.com/finanxier-app/internal/repository/product"
	"github.com/finanxier-app/internal/repository/user"
	productUc "github.com/finanxier-app/internal/usecase/product"
	userUc "github.com/finanxier-app/internal/usecase/user"

	_middleware "github.com/finanxier-app/middleware/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"

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
	if err != nil {
		return
	}

	// Only do for dev mode.
	if configs.IsDevelopmentMode() {
		// Init Schema migrations.
		logrus.Infof("Run schema migrations on %s", baseInitializer.Database.Config().ConnConfig.Database)
		err := tools.SchemaMigrate(baseInitializer.Database.Config().ConnString(), app.DatabaseVersion)
		if err != nil {
			logrus.Errorf("Fails when setup migrations %v", err)
		}

		// Init schema seeders.
		logrus.Info("Run schema seeders")
		err = tools.SchemaSeed(ctx, baseInitializer.DatabaseHelper)
		if err != nil {
			logrus.Errorf("Fails when setup seeder %v", err)
		}
	}

	// Init repo.
	logrus.Info("Initialize repo")
	productRepo := product.New(baseInitializer.Database)
	userRepo := user.New(baseInitializer.Database, configs.SecretKey)

	// Init usecase.
	logrus.Info("Initialize usecase")
	productUsecase := productUc.New(productRepo)
	userUsecase := userUc.New(userRepo)

	// Init handler.
	logrus.Info("Initialize handler")
	productHandler := productHandler.New(productUsecase, &baseInitializer)
	userHandler := userHandler.New(userUsecase, &baseInitializer)

	// Init routes.
	logrus.Info("Setup routes")
	routes := initializeRoutes(
		ctx,
		productHandler,
		userHandler,
		configs.SecretKey,
	)

	// Start server.
	logrus.Infof("Server start at port :%s", configs.App.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", configs.App.Port), routes)
}

func initializeRoutes(
	ctx context.Context,
	productHandler *productHandler.Handler,
	userHandler *userHandler.Handler,
	jwtKeySetup string,
) *chi.Mux {
	r := chi.NewMux()
	// Base url of all endpoints.
	baseApiVersion := "/api/v1"

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

	// Enable httprate request limiter of 100 requests per minute.
	//
	// In the code example below, rate-limiting is bound to the request IP address
	// via the LimitByIP middleware handler.
	// To have a single rate-limiter for all requests, use httprate.LimitAll(..).
	r.Use(httprate.LimitByIP(100, 1*time.Minute))

	// Endpoint section.
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("ok"))
	})

	// Helper endpoint.
	r.Get(baseApiVersion+"/token-generator", _middleware.GenericMiddleware(userHandler.GenerateJWT))
	// Core endpoints.
	r.Route(baseApiVersion, func(r chi.Router) {
		// Protect.
		r.Use(_middleware.JWTAuthMiddleware(jwtKeySetup))

		// Product management routes.
		r.Route("/product", func(r chi.Router) {
			r.Get("/{id}", _middleware.GenericMiddleware(productHandler.GetProduct))
			r.Get(constant.DefaultTrailing, _middleware.GenericMiddleware(productHandler.GetProducts))

			r.Post(constant.DefaultTrailing, _middleware.GenericMiddleware(productHandler.CreateProduct))
		})

		// User management routes.
		r.Route("/user", func(r chi.Router) {
			r.Get("/current", _middleware.GenericMiddleware(userHandler.GetCurrentUser))
		})
	})

	return r
}
