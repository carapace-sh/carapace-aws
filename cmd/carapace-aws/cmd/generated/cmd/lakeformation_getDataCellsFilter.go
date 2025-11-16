package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_getDataCellsFilterCmd = &cobra.Command{
	Use:   "get-data-cells-filter",
	Short: "Returns a data cells filter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_getDataCellsFilterCmd).Standalone()

	lakeformation_getDataCellsFilterCmd.Flags().String("database-name", "", "A database in the Glue Data Catalog.")
	lakeformation_getDataCellsFilterCmd.Flags().String("name", "", "The name given by the user to the data filter cell.")
	lakeformation_getDataCellsFilterCmd.Flags().String("table-catalog-id", "", "The ID of the catalog to which the table belongs.")
	lakeformation_getDataCellsFilterCmd.Flags().String("table-name", "", "A table in the database.")
	lakeformation_getDataCellsFilterCmd.MarkFlagRequired("database-name")
	lakeformation_getDataCellsFilterCmd.MarkFlagRequired("name")
	lakeformation_getDataCellsFilterCmd.MarkFlagRequired("table-catalog-id")
	lakeformation_getDataCellsFilterCmd.MarkFlagRequired("table-name")
	lakeformationCmd.AddCommand(lakeformation_getDataCellsFilterCmd)
}
