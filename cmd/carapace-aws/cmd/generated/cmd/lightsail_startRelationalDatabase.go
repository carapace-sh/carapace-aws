package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_startRelationalDatabaseCmd = &cobra.Command{
	Use:   "start-relational-database",
	Short: "Starts a specific database from a stopped state in Amazon Lightsail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_startRelationalDatabaseCmd).Standalone()

	lightsail_startRelationalDatabaseCmd.Flags().String("relational-database-name", "", "The name of your database to start.")
	lightsail_startRelationalDatabaseCmd.MarkFlagRequired("relational-database-name")
	lightsailCmd.AddCommand(lightsail_startRelationalDatabaseCmd)
}
