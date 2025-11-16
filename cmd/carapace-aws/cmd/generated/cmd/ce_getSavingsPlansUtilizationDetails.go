package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_getSavingsPlansUtilizationDetailsCmd = &cobra.Command{
	Use:   "get-savings-plans-utilization-details",
	Short: "Retrieves attribute data along with aggregate utilization and savings data for a given time period.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_getSavingsPlansUtilizationDetailsCmd).Standalone()

	ce_getSavingsPlansUtilizationDetailsCmd.Flags().String("data-type", "", "The data type.")
	ce_getSavingsPlansUtilizationDetailsCmd.Flags().String("filter", "", "Filters Savings Plans utilization coverage data for active Savings Plans dimensions.")
	ce_getSavingsPlansUtilizationDetailsCmd.Flags().String("max-results", "", "The number of items to be returned in a response.")
	ce_getSavingsPlansUtilizationDetailsCmd.Flags().String("next-token", "", "The token to retrieve the next set of results.")
	ce_getSavingsPlansUtilizationDetailsCmd.Flags().String("sort-by", "", "The value that you want to sort the data by.")
	ce_getSavingsPlansUtilizationDetailsCmd.Flags().String("time-period", "", "The time period that you want the usage and costs for.")
	ce_getSavingsPlansUtilizationDetailsCmd.MarkFlagRequired("time-period")
	ceCmd.AddCommand(ce_getSavingsPlansUtilizationDetailsCmd)
}
