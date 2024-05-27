package certservice

import (
	"context"
	"github.com/Dima191/cert-server/internal/models"
)

type Service interface {
	Cert(ctx context.Context, domain string) (*models.Cert, error)
}
