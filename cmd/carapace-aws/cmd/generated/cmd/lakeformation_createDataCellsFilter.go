package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_createDataCellsFilterCmd = &cobra.Command{
	Use:   "create-data-cells-filter",
	Short: "Creates a data cell filter to allow one to grant access to certain columns on certain rows.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_createDataCellsFilterCmd).Standalone()

	lakeformation_createDataCellsFilterCmd.Flags().String("table-data", "", "A `DataCellsFilter` structure containing information about the data cells filter.")
	lakeformation_createDataCellsFilterCmd.MarkFlagRequired("table-data")
	lakeformationCmd.AddCommand(lakeformation_createDataCellsFilterCmd)
}
