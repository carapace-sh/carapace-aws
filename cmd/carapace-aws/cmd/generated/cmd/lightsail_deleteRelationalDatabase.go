package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_deleteRelationalDatabaseCmd = &cobra.Command{
	Use:   "delete-relational-database",
	Short: "Deletes a database in Amazon Lightsail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_deleteRelationalDatabaseCmd).Standalone()

	lightsail_deleteRelationalDatabaseCmd.Flags().String("final-relational-database-snapshot-name", "", "The name of the database snapshot created if `skip final snapshot` is `false`, which is the default value for that parameter.")
	lightsail_deleteRelationalDatabaseCmd.Flags().String("relational-database-name", "", "The name of the database that you are deleting.")
	lightsail_deleteRelationalDatabaseCmd.Flags().String("skip-final-snapshot", "", "Determines whether a final database snapshot is created before your database is deleted.")
	lightsail_deleteRelationalDatabaseCmd.MarkFlagRequired("relational-database-name")
	lightsailCmd.AddCommand(lightsail_deleteRelationalDatabaseCmd)
}
