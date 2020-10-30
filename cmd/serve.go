package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serve = &cobra.Command{
	Use:   "serve",
	Short: "Just a test command TODO...",
	Long:  `long description here TODO...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO: serve API Server...")
		fmt.Println(viper.GetString("server.host") + ":" + viper.GetString("server.port"))
	},
}

func init() {
	rootCmd.AddCommand(serve)
}
