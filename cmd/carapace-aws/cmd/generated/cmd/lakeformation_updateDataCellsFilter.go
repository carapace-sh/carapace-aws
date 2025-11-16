package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_updateDataCellsFilterCmd = &cobra.Command{
	Use:   "update-data-cells-filter",
	Short: "Updates a data cell filter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_updateDataCellsFilterCmd).Standalone()

	lakeformation_updateDataCellsFilterCmd.Flags().String("table-data", "", "A `DataCellsFilter` structure containing information about the data cells filter.")
	lakeformation_updateDataCellsFilterCmd.MarkFlagRequired("table-data")
	lakeformationCmd.AddCommand(lakeformation_updateDataCellsFilterCmd)
}
