package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_getSavingsPlansUtilizationCmd = &cobra.Command{
	Use:   "get-savings-plans-utilization",
	Short: "Retrieves the Savings Plans utilization for your account across date ranges with daily or monthly granularity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_getSavingsPlansUtilizationCmd).Standalone()

	ce_getSavingsPlansUtilizationCmd.Flags().String("filter", "", "Filters Savings Plans utilization coverage data for active Savings Plans dimensions.")
	ce_getSavingsPlansUtilizationCmd.Flags().String("granularity", "", "The granularity of the Amazon Web Services utillization data for your Savings Plans.")
	ce_getSavingsPlansUtilizationCmd.Flags().String("sort-by", "", "The value that you want to sort the data by.")
	ce_getSavingsPlansUtilizationCmd.Flags().String("time-period", "", "The time period that you want the usage and costs for.")
	ce_getSavingsPlansUtilizationCmd.MarkFlagRequired("time-period")
	ceCmd.AddCommand(ce_getSavingsPlansUtilizationCmd)
}
