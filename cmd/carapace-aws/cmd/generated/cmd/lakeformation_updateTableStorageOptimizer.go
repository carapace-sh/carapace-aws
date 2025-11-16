package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_updateTableStorageOptimizerCmd = &cobra.Command{
	Use:   "update-table-storage-optimizer",
	Short: "Updates the configuration of the storage optimizers for a table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_updateTableStorageOptimizerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lakeformation_updateTableStorageOptimizerCmd).Standalone()

		lakeformation_updateTableStorageOptimizerCmd.Flags().String("catalog-id", "", "The Catalog ID of the table.")
		lakeformation_updateTableStorageOptimizerCmd.Flags().String("database-name", "", "Name of the database where the table is present.")
		lakeformation_updateTableStorageOptimizerCmd.Flags().String("storage-optimizer-config", "", "Name of the configuration for the storage optimizer.")
		lakeformation_updateTableStorageOptimizerCmd.Flags().String("table-name", "", "Name of the table for which to enable the storage optimizer.")
		lakeformation_updateTableStorageOptimizerCmd.MarkFlagRequired("database-name")
		lakeformation_updateTableStorageOptimizerCmd.MarkFlagRequired("storage-optimizer-config")
		lakeformation_updateTableStorageOptimizerCmd.MarkFlagRequired("table-name")
	})
	lakeformationCmd.AddCommand(lakeformation_updateTableStorageOptimizerCmd)
}
