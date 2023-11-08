package server

import (
	"example/komposervice/pkg/utils"

	"github.com/golang-migrate/migrate/v4"
	"github.com/urfave/cli/v2"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var Command = []*cli.Command{
	{
		Name:    "migrate",
		Aliases: []string{"m"},
		Usage:   "migrate database",
		Action: func(_ *cli.Context) error {
			const migrateUrl = "file://pkg/db/migration"
			databaseUrl, err := utils.ConnectionURLBuilder("pg-migrate")
			if err != nil {
				return err
			}
			_migrate, err := migrate.New(migrateUrl, databaseUrl)
			if err != nil {
				return err
			}
			return _migrate.Up()
		},
	},
	{
		Name:    "worker",
		Aliases: []string{"w"},
		Usage:   "run worker handle tasks in queue",
		Action: func(_ *cli.Context) error {
			return AsyncWorker(10)
		},
	},
	{
		Name:    "server",
		Aliases: []string{"s"},
		Usage:   "run api server",
		Action: func(_ *cli.Context) error {
			return Server()
		},
	},
}
