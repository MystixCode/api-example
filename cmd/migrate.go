package cmd

import (
	"first_go_app/internal/server"
	"fmt"

	//"first_go_app/pkg/logger"
	"first_go_app/pkg/models"

	"github.com/spf13/cobra"
)

// migrateCmd represents the version command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate database",
	Run: func(cmd *cobra.Command, args []string) {
		serve := server.New()
		serve.InitDatabase()
		err := serve.Database.AutoMigrate(
			models.User{},
			//			models.Group{},
			//			models.Scope{},
			//			models.ResourceServer{},
			//			models.Client{},
		)
		if err != nil {
			fmt.Println(err)
			//			logger.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
