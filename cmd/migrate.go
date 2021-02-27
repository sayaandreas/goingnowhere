package cmd

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" //
	_ "github.com/lib/pq"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run db migrations",
	Run: func(cmd *cobra.Command, args []string) {
		defer fmt.Println("Migration process done")
		dsn := viper.GetString("dsn")
		db, err := sql.Open("postgres", dsn)
		if err != nil {
			panic(err)
		}
		driver, err := postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			panic(err)
		}
		m, err := migrate.NewWithDatabaseInstance(
			"file://db/migrations",
			"postgres", driver)
		if err != nil {
			panic(err)
		}
		m.Steps(1)
	},
}
