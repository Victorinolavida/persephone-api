package infrastructure

import (
	"database/sql"
	"fmt"
	"github.com/Victorinolavida/go-crm-api/internal/config"
	"github.com/Victorinolavida/go-crm-api/internal/logger"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type DB struct {
	db *bun.DB
}

func NewDB(config config.DBConfig) (*DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.Username, config.Password, config.Host, config.Port, config.DbName)
	fmt.Println(dsn)
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	logger.GetLogger().Infof("connected to database")
	return &DB{db}, nil
}
