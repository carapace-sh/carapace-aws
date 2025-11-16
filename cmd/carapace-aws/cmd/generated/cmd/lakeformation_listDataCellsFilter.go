package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_listDataCellsFilterCmd = &cobra.Command{
	Use:   "list-data-cells-filter",
	Short: "Lists all the data cell filters on a table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_listDataCellsFilterCmd).Standalone()

	lakeformation_listDataCellsFilterCmd.Flags().String("max-results", "", "The maximum size of the response.")
	lakeformation_listDataCellsFilterCmd.Flags().String("next-token", "", "A continuation token, if this is a continuation call.")
	lakeformation_listDataCellsFilterCmd.Flags().String("table", "", "A table in the Glue Data Catalog.")
	lakeformationCmd.AddCommand(lakeformation_listDataCellsFilterCmd)
}
