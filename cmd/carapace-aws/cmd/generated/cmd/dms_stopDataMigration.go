package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_stopDataMigrationCmd = &cobra.Command{
	Use:   "stop-data-migration",
	Short: "Stops the specified data migration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_stopDataMigrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_stopDataMigrationCmd).Standalone()

		dms_stopDataMigrationCmd.Flags().String("data-migration-identifier", "", "The identifier (name or ARN) of the data migration to stop.")
		dms_stopDataMigrationCmd.MarkFlagRequired("data-migration-identifier")
	})
	dmsCmd.AddCommand(dms_stopDataMigrationCmd)
}
