package main

import (
	"context"
	"log"
	"net"

	"github.com/finanxier-app/cmd/app"
	"github.com/finanxier-app/config"
	"github.com/finanxier-app/config/tools"
	productHandler "github.com/finanxier-app/internal/handler/product/grpc"
	"github.com/finanxier-app/internal/repository/product"
	productUc "github.com/finanxier-app/internal/usecase/product"
	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/finanxier-app/proto/gen"
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

	// Init usecase.
	logrus.Info("Initialize usecase")
	productUsecase := productUc.New(productRepo)

	// Init handler.
	logrus.Info("Initialize handler")
	productHandler := productHandler.New(productUsecase, &baseInitializer)

	// Init gRPC.
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor())
	// Register.
	pb.RegisterProductServiceServer(grpcServer, productHandler)
	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	logrus.Info("Server listening at ", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		logrus.Fatalf("Failed to serve: %v", err)
	}
}
