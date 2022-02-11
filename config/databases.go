package config

import (
	"context"
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	Ctx  context.Context
	MSDB *sql.DB
	PGDB *gorm.DB
)

func InitializePGDB(dsn string) (err error) {
	PGDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "", // table name prefix, table for `User` would be `t_users`
		},
	})

	return
}

func InitializeMSDB(dsn string) (err error) {

	// Create connection pool
	MSDB, err = sql.Open("sqlserver", dsn)

	if err != nil {
		return err
	}
	Ctx = context.Background()
	err = MSDB.PingContext(Ctx)

	return err
}
