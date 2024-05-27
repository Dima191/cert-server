package app

import (
	"context"
	"fmt"
	"github.com/Dima191/cert-server/internal/api"
	pb "github.com/Dima191/cert-server/pkg/cert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
)

type App struct {
	grpc            *grpc.Server
	serviceProvider serviceProvider
	impl            *api.Implementation

	logger *slog.Logger

	configPath string
}

func (a *App) initImplementation(ctx context.Context) error {
	s, err := a.serviceProvider.certService(ctx)
	if err != nil {
		return err
	}
	a.impl = api.NewImplementation(s, a.logger)
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.logger, a.configPath)
	return nil
}

func (a *App) initGRPCServer(_ context.Context) error {
	a.grpc = grpc.NewServer()

	reflection.Register(a.grpc)

	pb.RegisterCertServer(a.grpc, a.impl)

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	deps := []func(ctx context.Context) error{
		a.initServiceProvider,
		a.initImplementation,
		a.initGRPCServer,
	}

	for _, f := range deps {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) runGRPCServer() error {
	cfg, err := a.serviceProvider.Config()
	if err != nil {
		a.logger.Error("failed to get config for implementation app", slog.String("error", err.Error()))
		return err
	}

	a.logger.Info("implementation app started", slog.Int("port", cfg.Port))

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		a.logger.Error("failed to create connection for implementation app", slog.String("error", err.Error()))
		return err
	}

	if err = a.grpc.Serve(listener); err != nil {
		a.logger.Error("failed to server implementation app", slog.String("error", err.Error()))
		return err
	}

	return nil
}

func (a *App) Run() error {
	return a.runGRPCServer()
}

func (a *App) Stop() {
	a.grpc.GracefulStop()

	if rep := a.serviceProvider.rep; rep != nil {
		rep.Close()
		//if err := ; err != nil {
		//	a.logger.Error("repository close error", slog.String("error", err.Error()))
		//}
		return
	}

	a.logger.Warn("repository was not been initialized")
}

func New(ctx context.Context, logger *slog.Logger, configPath string) (*App, error) {
	logger.Info("creating implementation app")

	a := App{}
	a.logger = logger
	a.configPath = configPath

	logger.Info("initializing all dependencies")

	err := a.initDeps(ctx)
	if err != nil {
		logger.Error("failed to initialize all dependencies", slog.String("error", err.Error()))
		return nil, fmt.Errorf("dependencies initialization error: %s", err.Error())
	}

	logger.Info("all dependencies initialized")

	return &a, nil
}
