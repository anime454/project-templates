package postgresql

import (
	"context"
	"fmt"

	"github.com/anime454/project-templates/go/hexagonal/car-parking-system/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Adapter struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, postgresCfg config.PostgreSQLConfig) (*Adapter, error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		postgresCfg.User,
		postgresCfg.Password,
		postgresCfg.Host,
		postgresCfg.Port,
		postgresCfg.DB,
	)

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("parse dsn: %w", err)
	}

	cfg.MaxConns = int32(postgresCfg.MaxConns)
	cfg.MinConns = int32(postgresCfg.MinConns)
	cfg.MaxConnLifetime = postgresCfg.MaxConnLifetime
	cfg.MaxConnIdleTime = postgresCfg.MaxConnIdleTime

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("connect: %w", err)
	}

	// Verify connection
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping: %w", err)
	}

	return &Adapter{pool: pool}, nil
}

func (a *Adapter) Pool() *pgxpool.Pool {
	return a.pool
}

func (a *Adapter) Close() error {
	if a.pool == nil {
		return nil
	}
	a.pool.Close()
	return nil
}
