package api

import (
	certservice "github.com/Dima191/cert-server/internal/service"
	pb "github.com/Dima191/cert-server/pkg/cert"
	"log/slog"
)

type Implementation struct {
	pb.UnimplementedCertServer
	service certservice.Service

	logger *slog.Logger
}

func NewImplementation(service certservice.Service, logger *slog.Logger) *Implementation {
	return &Implementation{
		service: service,
		logger:  logger,
	}
}
