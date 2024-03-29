package cmd

import (
	"api-example/config"
	"api-example/pkg/logger"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"fmt"
	"os"
)

var cfgFile string

var (
	rootCmd = &cobra.Command{
		Use:   "api-example",
		Short: "Example API",
		Long:  `This is an example api written in go.`,
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
