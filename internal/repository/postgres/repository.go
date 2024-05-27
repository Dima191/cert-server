package postgresrepository

import (
	"context"
	"github.com/Dima191/cert-server/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type rep struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, connStr string) (certrepository.Repository, error) {
	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, err
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, err
	}

	r := new(rep)
	r.pool = pool

	return r, nil

}
