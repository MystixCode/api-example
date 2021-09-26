package cmd

import (
	"first_go_app/config"
	"first_go_app/pkg/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"fmt"
	"os"
)

var cfgFile string

var (
	rootCmd = &cobra.Command{
		Use:   "first_go_app",
		Short: "My first go app",
		Long:  `This app contains boilerplate for cli via cobra and boilerplate for json config via viper.`,
	}
)

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}

func init() {
	initConfig()
}

func initConfig() {

	cobra.OnInitialize(config.SetDefaults, logger.Init)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.json)")

	viper.SetConfigType("json")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigFile("./config.json")
	}
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	//TODO if config file not exists create it with default values

}
