package repository

import (
	"database/sql"
	"fmt"

	"github.com/osg_template/internal/pkg/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func NewPostgres() *bun.DB {
	cfg := config.Load()

	postgresUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	)

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(postgresUrl)))

	return bun.NewDB(sqldb, pgdialect.New())
}
