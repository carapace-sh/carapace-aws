package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_deleteRelationalDatabaseSnapshotCmd = &cobra.Command{
	Use:   "delete-relational-database-snapshot",
	Short: "Deletes a database snapshot in Amazon Lightsail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_deleteRelationalDatabaseSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_deleteRelationalDatabaseSnapshotCmd).Standalone()

		lightsail_deleteRelationalDatabaseSnapshotCmd.Flags().String("relational-database-snapshot-name", "", "The name of the database snapshot that you are deleting.")
		lightsail_deleteRelationalDatabaseSnapshotCmd.MarkFlagRequired("relational-database-snapshot-name")
	})
	lightsailCmd.AddCommand(lightsail_deleteRelationalDatabaseSnapshotCmd)
}
