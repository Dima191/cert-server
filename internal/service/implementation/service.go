package certserviceimpl

import (
	"github.com/Dima191/cert-server/internal/repository"
	certservice "github.com/Dima191/cert-server/internal/service"
	"github.com/go-playground/validator/v10"
	"log/slog"
)

type service struct {
	rep certrepository.Repository
	v   *validator.Validate

	logger *slog.Logger
}

func New(rep certrepository.Repository, logger *slog.Logger) certservice.Service {
	return &service{
		rep:    rep,
		v:      validator.New(),
		logger: logger,
	}
}
