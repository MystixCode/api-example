package cmd

import (
	"first_go_app/config"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig, config.SetDefaults)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.json)")
}

func initConfig() {
	viper.SetConfigType("json")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	//TODO if config file not exists create it with default values
}
