package app

import (
	"context"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/adapter/grpc"
	httpserver "github.com/sorawaslocked/ap2final_api_gateway/internal/adapter/http"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/config"
	grpcconn "github.com/sorawaslocked/ap2final_api_gateway/internal/pkg/grpc"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/usecase"
	"github.com/sorawaslocked/ap2final_base/pkg/logger"
	grpcActorSvc "github.com/sorawaslocked/ap2final_protos_gen/service/actor"
	grpcMovieSvc "github.com/sorawaslocked/ap2final_protos_gen/service/movie"
	grpcSessionSvc "github.com/sorawaslocked/ap2final_protos_gen/service/session"
	grpcUserSvc "github.com/sorawaslocked/ap2final_protos_gen/service/user"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const serviceName = "api-gateway"

type App struct {
	httpServer *httpserver.API
	log        *slog.Logger
}

func New(cfg *config.Config, log *slog.Logger) (*App, error) {
	log.Info("starting service", slog.String("service", serviceName))

	movieServiceGRPCConn, err := grpcconn.Connect(
		cfg.GRPC.Client.MovieServiceURL,
		cfg.GRPC.Client,
	)
	if err != nil {
		log.Error(
			"failed to connect to grpc service",
			slog.String("service", "movie service"),
			slog.String("url", cfg.GRPC.Client.MovieServiceURL),
			logger.Err(err),
		)
	}

	userServiceGRPCConn, err := grpcconn.Connect(
		cfg.GRPC.Client.UserServiceURL,
		cfg.GRPC.Client,
	)
	if err != nil {
		log.Error(
			"failed to connect to grpc service",
			slog.String("service", "user service"),
			slog.String("url", cfg.GRPC.Client.UserServiceURL),
			logger.Err(err),
		)
	}

	// Actor Service Connection
	actorServiceGRPCConn, err := grpcconn.Connect(
		cfg.GRPC.Client.ActorServiceURL,
		cfg.GRPC.Client,
	)
	if err != nil {
		log.Error(
			"failed to connect to grpc service",
			slog.String("service", "actor service"),
			slog.String("url", cfg.GRPC.Client.ActorServiceURL),
			logger.Err(err),
		)
	}

	// Session Service Connection
	sessionServiceGRPCConn, err := grpcconn.Connect(
		cfg.GRPC.Client.SessionServiceURL,
		cfg.GRPC.Client,
	)
	if err != nil {
		log.Error(
			"failed to connect to grpc service",
			slog.String("service", "session service"),
			slog.String("url", cfg.GRPC.Client.SessionServiceURL),
			logger.Err(err),
		)
	}

	movieServiceGRPCClient := grpcMovieSvc.NewMovieServiceClient(movieServiceGRPCConn)
	movieServiceGRPCHandler := grpc.NewMovie(movieServiceGRPCClient)

	userServiceGRPCClient := grpcUserSvc.NewUserServiceClient(userServiceGRPCConn)
	userServiceGRPCHandler := grpc.NewUser(userServiceGRPCClient)

	actorServiceGRPCClient := grpcActorSvc.NewActorServiceClient(actorServiceGRPCConn)
	actorServiceGRPCHandler := grpc.NewActor(actorServiceGRPCClient)

	sessionServiceGRPCClient := grpcSessionSvc.NewSessionServiceClient(sessionServiceGRPCConn)
	sessionServiceGRPCHandler := grpc.NewSession(sessionServiceGRPCClient)

	movieUseCase := usecase.NewMovie(movieServiceGRPCHandler)
	userUseCase := usecase.NewUser(userServiceGRPCHandler)
	actorUseCase := usecase.NewActor(actorServiceGRPCHandler)
	sessionUseCase := usecase.NewSession(sessionServiceGRPCHandler)

	httpServer := httpserver.New(cfg.HTTPServer, log, movieUseCase, userUseCase, actorUseCase, sessionUseCase)

	app := &App{
		httpServer: httpServer,
		log:        log,
	}

	return app, nil
}

func (a *App) Run() {
	a.httpServer.MustRun()

	a.log.Info("service started", slog.String("service", serviceName))

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	s := <-shutdown

	a.log.Info("received signal", slog.String("signal", s.String()))
	a.Stop()
	a.log.Info("graceful shutdown complete")
}

func (a *App) Stop() {
	a.log.Info("shutting down http server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	a.httpServer.Stop(ctx)
}
