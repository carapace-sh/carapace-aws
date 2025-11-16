package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_getSavingsPlansCoverageCmd = &cobra.Command{
	Use:   "get-savings-plans-coverage",
	Short: "Retrieves the Savings Plans covered for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_getSavingsPlansCoverageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ce_getSavingsPlansCoverageCmd).Standalone()

		ce_getSavingsPlansCoverageCmd.Flags().String("filter", "", "Filters Savings Plans coverage data by dimensions.")
		ce_getSavingsPlansCoverageCmd.Flags().String("granularity", "", "The granularity of the Amazon Web Services cost data for your Savings Plans.")
		ce_getSavingsPlansCoverageCmd.Flags().String("group-by", "", "You can group the data using the attributes `INSTANCE_FAMILY`, `REGION`, or `SERVICE`.")
		ce_getSavingsPlansCoverageCmd.Flags().String("max-results", "", "The number of items to be returned in a response.")
		ce_getSavingsPlansCoverageCmd.Flags().String("metrics", "", "The measurement that you want your Savings Plans coverage reported in.")
		ce_getSavingsPlansCoverageCmd.Flags().String("next-token", "", "The token to retrieve the next set of results.")
		ce_getSavingsPlansCoverageCmd.Flags().String("sort-by", "", "The value that you want to sort the data by.")
		ce_getSavingsPlansCoverageCmd.Flags().String("time-period", "", "The time period that you want the usage and costs for.")
		ce_getSavingsPlansCoverageCmd.MarkFlagRequired("time-period")
	})
	ceCmd.AddCommand(ce_getSavingsPlansCoverageCmd)
}
