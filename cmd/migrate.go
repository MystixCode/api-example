package cmd

import (
	"api-example/internal/server"
	"api-example/pkg/logger"
	"api-example/pkg/models"

	"github.com/spf13/cobra"
)

// migrateCmd creates database structure based on models
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate database",
	Run: func(cmd *cobra.Command, args []string) {
		serve := server.New()
		serve.InitDatabase()
		err := serve.Database.AutoMigrate(
			models.User{},
			models.Test{},
		)
		if err != nil {
			logger.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
