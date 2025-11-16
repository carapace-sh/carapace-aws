package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_deleteDataMigrationCmd = &cobra.Command{
	Use:   "delete-data-migration",
	Short: "Deletes the specified data migration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_deleteDataMigrationCmd).Standalone()

	dms_deleteDataMigrationCmd.Flags().String("data-migration-identifier", "", "The identifier (name or ARN) of the data migration to delete.")
	dms_deleteDataMigrationCmd.MarkFlagRequired("data-migration-identifier")
	dmsCmd.AddCommand(dms_deleteDataMigrationCmd)
}
