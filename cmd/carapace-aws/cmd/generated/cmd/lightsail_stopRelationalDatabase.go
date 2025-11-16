package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_stopRelationalDatabaseCmd = &cobra.Command{
	Use:   "stop-relational-database",
	Short: "Stops a specific database that is currently running in Amazon Lightsail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_stopRelationalDatabaseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_stopRelationalDatabaseCmd).Standalone()

		lightsail_stopRelationalDatabaseCmd.Flags().String("relational-database-name", "", "The name of your database to stop.")
		lightsail_stopRelationalDatabaseCmd.Flags().String("relational-database-snapshot-name", "", "The name of your new database snapshot to be created before stopping your database.")
		lightsail_stopRelationalDatabaseCmd.MarkFlagRequired("relational-database-name")
	})
	lightsailCmd.AddCommand(lightsail_stopRelationalDatabaseCmd)
}
