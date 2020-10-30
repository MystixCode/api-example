package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var testcmd = &cobra.Command{
	Use:   "testcmd",
	Short: "Just a test command",
	Long:  `long description here`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test")
		addr := viper.GetString("server.host") + ":" + viper.GetString("server.port")
		fmt.Println(addr)
	},
}

func init() {
	rootCmd.AddCommand(testcmd)
}
