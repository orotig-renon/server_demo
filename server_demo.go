package main

import (
	"fmt"

	"github.com/orotig/server_demo/config"
	"github.com/orotig/server_demo/models"
	"github.com/orotig/server_demo/routers"
)

//POSTGRES DATABASE
var pg_server = "database"
var pg_port = 5432
var pg_user = "postgres"
var pg_password = "postgres"
var pg_database = "orotig"
var pg_encrypt = "disable"

func main() {

	pg_dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		pg_server, pg_user, pg_password, pg_database, pg_port, pg_encrypt)

	if err := config.InitializePGDB(pg_dsn); err != nil {
		panic("failed to connect Postgres database")
	}

	if err := models.Migrate(config.PGDB); err != nil {
		panic("failed to migrate Postgres database")
	}

	router := routers.SetupRouter()
	router.Run(":8080")
}
