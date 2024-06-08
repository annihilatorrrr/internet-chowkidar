package db

import (
	"context"
	"log/slog"

	"github.com/gnulinuxindia/internet-chowkidar/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func ProvideDBConn(conf *config.Config) (*pgxpool.Pool, error) {
	if pool == nil {
		var err error

		pool, err = pgxpool.New(context.Background(), conf.DbUrl)
		if err != nil {
			slog.Error("error connecting to db", err)
			return nil, err
		}
	}

	return pool, nil
}

func ProvideQueries(pool *pgxpool.Pool) *Queries {
	return New(pool)
}
