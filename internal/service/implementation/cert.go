package certserviceimpl

import (
	"context"
	"errors"
	"github.com/Dima191/cert-server/internal/models"
	certrepositoryerror "github.com/Dima191/cert-server/internal/repository/error"
	"google.golang.org/grpc/status"
	"log/slog"
	"time"
)

const (
	timeout = 100 * time.Millisecond
)

func (s *service) Cert(ctx context.Context, domain string) (*models.Cert, error) {
	if err := s.validate(domain); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	cert, err := s.rep.Cert(ctx, domain)
	if err != nil {
		if errors.Is(err, certrepositoryerror.ErrNoRows) {
			s.logger.Error("no requested data", slog.String("domain", domain))
			return nil, status.Errorf(5, "no data")
		}

		s.logger.Error("failed to get implementation", slog.String("error", err.Error()))
		return nil, status.Errorf(13, "get implementation error")
	}
	return cert, nil
}
