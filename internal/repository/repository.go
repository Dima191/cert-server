package certrepository

import (
	"context"
	"github.com/Dima191/cert-server/internal/models"
)

type Repository interface {
	Cert(ctx context.Context, domain string) (*models.Cert, error)
	Close()
}
