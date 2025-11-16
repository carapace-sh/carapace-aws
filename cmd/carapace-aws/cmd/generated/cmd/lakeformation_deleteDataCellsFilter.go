package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_deleteDataCellsFilterCmd = &cobra.Command{
	Use:   "delete-data-cells-filter",
	Short: "Deletes a data cell filter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_deleteDataCellsFilterCmd).Standalone()

	lakeformation_deleteDataCellsFilterCmd.Flags().String("database-name", "", "A database in the Glue Data Catalog.")
	lakeformation_deleteDataCellsFilterCmd.Flags().String("name", "", "The name given by the user to the data filter cell.")
	lakeformation_deleteDataCellsFilterCmd.Flags().String("table-catalog-id", "", "The ID of the catalog to which the table belongs.")
	lakeformation_deleteDataCellsFilterCmd.Flags().String("table-name", "", "A table in the database.")
	lakeformationCmd.AddCommand(lakeformation_deleteDataCellsFilterCmd)
}
