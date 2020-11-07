package cmd

import (
	"first_go_app/internal/server"

	"fmt"
	"github.com/spf13/cobra"
)

var serve = &cobra.Command{
	Use:   "serve",
	Short: "Just a test command TODO...",
	Long:  `long description here TODO...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Serving API Server...")
		serve := server.New()
		serve.Run()
	},
}

func init() {
	rootCmd.AddCommand(serve)
}
