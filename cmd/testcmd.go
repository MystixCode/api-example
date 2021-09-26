package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var testcmd = &cobra.Command{
	Use:   "testcmd",
	Short: "Just a test command",
	Long:  `long description here`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test")
	},
}

func init() {
	rootCmd.AddCommand(testcmd)
}
