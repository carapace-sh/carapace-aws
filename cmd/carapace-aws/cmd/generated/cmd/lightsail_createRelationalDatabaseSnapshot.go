package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_createRelationalDatabaseSnapshotCmd = &cobra.Command{
	Use:   "create-relational-database-snapshot",
	Short: "Creates a snapshot of your database in Amazon Lightsail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_createRelationalDatabaseSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_createRelationalDatabaseSnapshotCmd).Standalone()

		lightsail_createRelationalDatabaseSnapshotCmd.Flags().String("relational-database-name", "", "The name of the database on which to base your new snapshot.")
		lightsail_createRelationalDatabaseSnapshotCmd.Flags().String("relational-database-snapshot-name", "", "The name for your new database snapshot.")
		lightsail_createRelationalDatabaseSnapshotCmd.Flags().String("tags", "", "The tag keys and optional values to add to the resource during create.")
		lightsail_createRelationalDatabaseSnapshotCmd.MarkFlagRequired("relational-database-name")
		lightsail_createRelationalDatabaseSnapshotCmd.MarkFlagRequired("relational-database-snapshot-name")
	})
	lightsailCmd.AddCommand(lightsail_createRelationalDatabaseSnapshotCmd)
}
