package postgresrepository

import (
	"context"
	"errors"
	"github.com/Dima191/cert-server/internal/models"
	certrepositoryerror "github.com/Dima191/cert-server/internal/repository/error"
	"github.com/jackc/pgx/v5"
)

func (r *rep) Cert(ctx context.Context, domain string) (*models.Cert, error) {
	var cert models.Cert

	row := r.pool.QueryRow(ctx, "select private_key, certificate_chain from cert where domain = $1", domain)

	err := row.Scan(&cert.PrivateKey, &cert.CertificateChain)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, certrepositoryerror.ErrNoRows
		}
		return nil, err
	}

	return &cert, nil
}
