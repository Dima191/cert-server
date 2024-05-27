package app

import (
	"context"
	"github.com/Dima191/cert-server/internal/config"
	"github.com/Dima191/cert-server/internal/repository"
	postgresrepository "github.com/Dima191/cert-server/internal/repository/postgres"
	certservice "github.com/Dima191/cert-server/internal/service"
	certserviceimpl "github.com/Dima191/cert-server/internal/service/implementation"
	"log/slog"
)

type serviceProvider struct {
	rep        certrepository.Repository
	cfg        *config.Config
	service    certservice.Service
	configPath string

	logger *slog.Logger
}

func (sp *serviceProvider) Config() (*config.Config, error) {
	if sp.cfg == nil {
		var err error
		sp.cfg, err = config.New(sp.configPath)
		if err != nil {
			return nil, err
		}
	}
	return sp.cfg, nil
}

func (sp *serviceProvider) postgresRepository(ctx context.Context) (certrepository.Repository, error) {
	if sp.rep == nil {
		cfg, err := sp.Config()
		if err != nil {
			return nil, err
		}

		sp.rep, err = postgresrepository.New(ctx, cfg.DBConnStr)
		if err != nil {
			return nil, err
		}
	}

	return sp.rep, nil
}

func (sp *serviceProvider) certService(ctx context.Context) (certservice.Service, error) {
	if sp.service == nil {
		rep, err := sp.postgresRepository(ctx)
		if err != nil {
			return nil, err
		}
		sp.service = certserviceimpl.New(rep, sp.logger)
	}

	return sp.service, nil
}

func newServiceProvider(logger *slog.Logger, configPath string) serviceProvider {
	return serviceProvider{
		logger:     logger,
		configPath: configPath,
	}
}
