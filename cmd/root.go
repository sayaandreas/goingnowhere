package cmd

import (
	"fmt"
	"net/http"

	pgadapter "github.com/casbin/casbin-pg-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/go-pg/pg/v9"
	"github.com/sayaandreas/goingnowhere/api"
	"github.com/sayaandreas/goingnowhere/storage"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "goingnowhere",
	Short: "Run http server",
	Run: func(cmd *cobra.Command, args []string) {
		store := storage.NewStorageSession()
		opts, err := pg.ParseURL(viper.GetString("dsn"))
		if err != nil {
			panic(err)
		}
		db := pg.Connect(opts)
		adapter, err := pgadapter.NewAdapterByDB(db)
		e, err := casbin.NewEnforcer("./rbac_model.conf", adapter)
		if err != nil {
			panic(err)
		}
		err = e.LoadPolicy()
		if err != nil {
			panic(err)
		}
		httpHandler := api.NewHandler(store, e)
		http.ListenAndServe(viper.GetString("addr"), httpHandler)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(migrateCmd)
	rootCmd.AddCommand(casbinCmd)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	viper.AutomaticEnv()
}
