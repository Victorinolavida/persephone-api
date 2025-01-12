package infrastructure

import (
	"database/sql"
	"fmt"
	"github.com/Victorinolavida/persephone-api/config"
	"github.com/Victorinolavida/persephone-api/pkg/logger"
	"github.com/uptrace/bun"

	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func NewDB(config config.DBConfig) (*bun.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.Username, config.Password, config.Host, config.Port, config.DbName)
	fmt.Println(dsn)
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	logger.GetLogger().Infof("connected to database")
	return db, nil
}
