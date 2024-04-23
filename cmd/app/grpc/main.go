package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/finanxier-app/cmd/app"
	"github.com/finanxier-app/config"
	"github.com/finanxier-app/config/tools"
	productHandler "github.com/finanxier-app/internal/handler/product/grpc"
	userHandler "github.com/finanxier-app/internal/handler/user"
	"github.com/finanxier-app/internal/repository/product"
	"github.com/finanxier-app/internal/repository/user"
	productUc "github.com/finanxier-app/internal/usecase/product"
	userUc "github.com/finanxier-app/internal/usecase/user"
	_middleware "github.com/finanxier-app/middleware/grpc"
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
	userRepo := user.New(baseInitializer.Database, configs.SecretKey)

	// Init usecase.
	logrus.Info("Initialize usecase")
	productUsecase := productUc.New(productRepo)
	userUsecase := userUc.New(userRepo)

	// Init handler.
	logrus.Info("Initialize handler")
	productHandler := productHandler.New(productUsecase, &baseInitializer)
	userHandler := userHandler.New(userUsecase, &baseInitializer)

	// Register all methods.
	grpcServer := registerMethods(
		ctx,
		productHandler,
		userHandler,
	)

	// Server configurations.
	// Init gRPC.
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.App.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	logrus.Info("Server listening at ", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		logrus.Fatalf("Failed to serve: %v", err)
	}
}

func registerMethods(
	ctx context.Context,
	pHandle *productHandler.Handler,
	uHandle *userHandler.Handler,
) *grpc.Server {
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		_middleware.JWTAuthMiddleware,
		_middleware.LoggerMiddleware,
	))

	// Register.
	pb.RegisterProductServiceServer(grpcServer, pHandle)
	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	return grpcServer
}
