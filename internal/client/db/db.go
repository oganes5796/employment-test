package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/oganes5796/employment-test/pkg/logger"
	"go.uber.org/zap"
)

func NewPostgresPool(dsn string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		logger.Error("failed to parse dsn", zap.Error(err))
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		logger.Error("failed to create pool", zap.Error(err))
		return nil, err
	}

	return pool, nil
}
