package postgres

import (
	"database/sql"

	"main/services/auth-service/internal/config"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type DB interface {
	bun.IDB
}

func New(cfg config.DBConfig) DB {
	pgDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(cfg.GetConnectionString())))
	db := bun.NewDB(pgDB, pgdialect.New())

	return db
}
