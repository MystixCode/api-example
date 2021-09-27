package cmd

import (
	"api-example/internal/server"
	"api-example/pkg/logger"

	"github.com/spf13/cobra"
)

var serve = &cobra.Command{
	Use:   "serve",
	Short: "Start the API Server",
	Long:  `Starts Database, Api, Router and HttpServer`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("serve")
		serve := server.New()
		serve.Run()
	},
}

func init() {
	rootCmd.AddCommand(serve)
}
