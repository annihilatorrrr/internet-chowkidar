//go:generate wire
//go:build wireinject

package di

import (
	"github.com/gnulinuxindia/internet-chowkidar/internal/config"
	"github.com/gnulinuxindia/internet-chowkidar/pkg/db"
	"github.com/google/wire"
)

var dbSet = wire.NewSet(
	db.ProvideQueries,
	db.ProvideDBConn,
	config.ProvideConfig,
)

func InjectDb() (*db.Queries, error) {
	wire.Build(dbSet)
	return &db.Queries{}, nil
}
