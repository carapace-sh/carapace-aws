package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getRelationalDatabaseSnapshotCmd = &cobra.Command{
	Use:   "get-relational-database-snapshot",
	Short: "Returns information about a specific database snapshot in Amazon Lightsail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getRelationalDatabaseSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_getRelationalDatabaseSnapshotCmd).Standalone()

		lightsail_getRelationalDatabaseSnapshotCmd.Flags().String("relational-database-snapshot-name", "", "The name of the database snapshot for which to get information.")
		lightsail_getRelationalDatabaseSnapshotCmd.MarkFlagRequired("relational-database-snapshot-name")
	})
	lightsailCmd.AddCommand(lightsail_getRelationalDatabaseSnapshotCmd)
}
