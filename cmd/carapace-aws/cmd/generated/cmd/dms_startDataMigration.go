package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_startDataMigrationCmd = &cobra.Command{
	Use:   "start-data-migration",
	Short: "Starts the specified data migration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_startDataMigrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_startDataMigrationCmd).Standalone()

		dms_startDataMigrationCmd.Flags().String("data-migration-identifier", "", "The identifier (name or ARN) of the data migration to start.")
		dms_startDataMigrationCmd.Flags().String("start-type", "", "Specifies the start type for the data migration.")
		dms_startDataMigrationCmd.MarkFlagRequired("data-migration-identifier")
		dms_startDataMigrationCmd.MarkFlagRequired("start-type")
	})
	dmsCmd.AddCommand(dms_startDataMigrationCmd)
}
