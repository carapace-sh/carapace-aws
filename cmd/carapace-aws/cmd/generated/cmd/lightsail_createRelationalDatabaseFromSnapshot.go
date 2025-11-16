package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_createRelationalDatabaseFromSnapshotCmd = &cobra.Command{
	Use:   "create-relational-database-from-snapshot",
	Short: "Creates a new database from an existing database snapshot in Amazon Lightsail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_createRelationalDatabaseFromSnapshotCmd).Standalone()

	lightsail_createRelationalDatabaseFromSnapshotCmd.Flags().String("availability-zone", "", "The Availability Zone in which to create your new database.")
	lightsail_createRelationalDatabaseFromSnapshotCmd.Flags().String("publicly-accessible", "", "Specifies the accessibility options for your new database.")
	lightsail_createRelationalDatabaseFromSnapshotCmd.Flags().String("relational-database-bundle-id", "", "The bundle ID for your new database.")
	lightsail_createRelationalDatabaseFromSnapshotCmd.Flags().String("relational-database-name", "", "The name to use for your new Lightsail database resource.")
	lightsail_createRelationalDatabaseFromSnapshotCmd.Flags().String("relational-database-snapshot-name", "", "The name of the database snapshot from which to create your new database.")
	lightsail_createRelationalDatabaseFromSnapshotCmd.Flags().String("restore-time", "", "The date and time to restore your database from.")
	lightsail_createRelationalDatabaseFromSnapshotCmd.Flags().String("source-relational-database-name", "", "The name of the source database.")
	lightsail_createRelationalDatabaseFromSnapshotCmd.Flags().String("tags", "", "The tag keys and optional values to add to the resource during create.")
	lightsail_createRelationalDatabaseFromSnapshotCmd.Flags().String("use-latest-restorable-time", "", "Specifies whether your database is restored from the latest backup time.")
	lightsail_createRelationalDatabaseFromSnapshotCmd.MarkFlagRequired("relational-database-name")
	lightsailCmd.AddCommand(lightsail_createRelationalDatabaseFromSnapshotCmd)
}
