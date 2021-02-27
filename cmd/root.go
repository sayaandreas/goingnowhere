package cmd

import (
	"fmt"
	"net/http"

	"github.com/sayaandreas/goingnowhere/api"
	"github.com/sayaandreas/goingnowhere/storage"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "goingnowhere",
	Short: "Run http server",
	Run: func(cmd *cobra.Command, args []string) {
		store := storage.NewStorageSession()
		httpHandler := api.NewHandler(store)
		http.ListenAndServe(viper.GetString("addr"), httpHandler)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(migrateCmd)
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
