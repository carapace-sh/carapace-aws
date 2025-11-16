package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_listTableStorageOptimizersCmd = &cobra.Command{
	Use:   "list-table-storage-optimizers",
	Short: "Returns the configuration of all storage optimizers associated with a specified table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_listTableStorageOptimizersCmd).Standalone()

	lakeformation_listTableStorageOptimizersCmd.Flags().String("catalog-id", "", "The Catalog ID of the table.")
	lakeformation_listTableStorageOptimizersCmd.Flags().String("database-name", "", "Name of the database where the table is present.")
	lakeformation_listTableStorageOptimizersCmd.Flags().String("max-results", "", "The number of storage optimizers to return on each call.")
	lakeformation_listTableStorageOptimizersCmd.Flags().String("next-token", "", "A continuation token, if this is a continuation call.")
	lakeformation_listTableStorageOptimizersCmd.Flags().String("storage-optimizer-type", "", "The specific type of storage optimizers to list.")
	lakeformation_listTableStorageOptimizersCmd.Flags().String("table-name", "", "Name of the table.")
	lakeformation_listTableStorageOptimizersCmd.MarkFlagRequired("database-name")
	lakeformation_listTableStorageOptimizersCmd.MarkFlagRequired("table-name")
	lakeformationCmd.AddCommand(lakeformation_listTableStorageOptimizersCmd)
}
