package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_createRelationalDatabaseCmd = &cobra.Command{
	Use:   "create-relational-database",
	Short: "Creates a new database in Amazon Lightsail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_createRelationalDatabaseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_createRelationalDatabaseCmd).Standalone()

		lightsail_createRelationalDatabaseCmd.Flags().String("availability-zone", "", "The Availability Zone in which to create your new database.")
		lightsail_createRelationalDatabaseCmd.Flags().String("master-database-name", "", "The meaning of this parameter differs according to the database engine you use.")
		lightsail_createRelationalDatabaseCmd.Flags().String("master-user-password", "", "The password for the master user.")
		lightsail_createRelationalDatabaseCmd.Flags().String("master-username", "", "The name for the master user.")
		lightsail_createRelationalDatabaseCmd.Flags().String("preferred-backup-window", "", "The daily time range during which automated backups are created for your new database if automated backups are enabled.")
		lightsail_createRelationalDatabaseCmd.Flags().String("preferred-maintenance-window", "", "The weekly time range during which system maintenance can occur on your new database.")
		lightsail_createRelationalDatabaseCmd.Flags().String("publicly-accessible", "", "Specifies the accessibility options for your new database.")
		lightsail_createRelationalDatabaseCmd.Flags().String("relational-database-blueprint-id", "", "The blueprint ID for your new database.")
		lightsail_createRelationalDatabaseCmd.Flags().String("relational-database-bundle-id", "", "The bundle ID for your new database.")
		lightsail_createRelationalDatabaseCmd.Flags().String("relational-database-name", "", "The name to use for your new Lightsail database resource.")
		lightsail_createRelationalDatabaseCmd.Flags().String("tags", "", "The tag keys and optional values to add to the resource during create.")
		lightsail_createRelationalDatabaseCmd.MarkFlagRequired("master-database-name")
		lightsail_createRelationalDatabaseCmd.MarkFlagRequired("master-username")
		lightsail_createRelationalDatabaseCmd.MarkFlagRequired("relational-database-blueprint-id")
		lightsail_createRelationalDatabaseCmd.MarkFlagRequired("relational-database-bundle-id")
		lightsail_createRelationalDatabaseCmd.MarkFlagRequired("relational-database-name")
	})
	lightsailCmd.AddCommand(lightsail_createRelationalDatabaseCmd)
}
