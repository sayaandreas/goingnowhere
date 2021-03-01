package cmd

import (
	"fmt"

	pgadapter "github.com/casbin/casbin-pg-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/go-pg/pg/v9"
	_ "github.com/golang-migrate/migrate/v4/source/file" //
	_ "github.com/lib/pq"

	"github.com/spf13/cobra"
)

var casbinCmd = &cobra.Command{
	Use:   "casbin",
	Short: "Run db migrations",
	Run: func(cmd *cobra.Command, args []string) {
		// dsn := viper.GetString("dsn")
		// a, _ := pgadapter.NewAdapter(dsn)
		// e, _ := casbin.NewEnforcer("./rbac_model.conf", a)
		// e, err := casbin.NewEnforcer("./rbac_model.conf", "./policy.csv")
		// e.LoadPolicy()
		// ok, err := enforce("bob", "data2", "write")
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(ok)

		enforce("bob", "data2", "write")
	},
}

func enforce(sub string, obj string, act string) {
	// Load model configuration file and policy store adapter
	// a, err := pgadapter.NewAdapter("postgres://postgres:password@localhost:5432/goingnowhere
	// ?sslmode=disable")
	opts, _ := pg.ParseURL("postgres://postgres:password@localhost:5432/goingnowhere?sslmode=disable")
	db := pg.Connect(opts)
	a, err := pgadapter.NewAdapterByDB(db)
	enforcer, _ := casbin.NewEnforcer("./rbac_model.conf", a)
	// Load policies from DB dynamically
	err = enforcer.LoadPolicy()
	if err != nil {
		fmt.Println(err)
	}
	// Verify
	ok, err := enforcer.Enforce(sub, obj, act)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ok)
}
