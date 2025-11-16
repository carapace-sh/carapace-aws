package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_getCostComparisonDriversCmd = &cobra.Command{
	Use:   "get-cost-comparison-drivers",
	Short: "Retrieves key factors driving cost changes between two time periods within the last 13 months, such as usage changes, discount changes, and commitment-based savings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_getCostComparisonDriversCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ce_getCostComparisonDriversCmd).Standalone()

		ce_getCostComparisonDriversCmd.Flags().String("baseline-time-period", "", "The reference time period for comparison.")
		ce_getCostComparisonDriversCmd.Flags().String("billing-view-arn", "", "The Amazon Resource Name (ARN) that uniquely identifies a specific billing view.")
		ce_getCostComparisonDriversCmd.Flags().String("comparison-time-period", "", "The comparison time period for analysis.")
		ce_getCostComparisonDriversCmd.Flags().String("filter", "", "")
		ce_getCostComparisonDriversCmd.Flags().String("group-by", "", "You can group results using the attributes `DIMENSION`, `TAG`, and `COST_CATEGORY`.")
		ce_getCostComparisonDriversCmd.Flags().String("max-results", "", "The maximum number of results that are returned for the request.")
		ce_getCostComparisonDriversCmd.Flags().String("metric-for-comparison", "", "The cost and usage metric to compare.")
		ce_getCostComparisonDriversCmd.Flags().String("next-page-token", "", "The token to retrieve the next set of paginated results.")
		ce_getCostComparisonDriversCmd.MarkFlagRequired("baseline-time-period")
		ce_getCostComparisonDriversCmd.MarkFlagRequired("comparison-time-period")
		ce_getCostComparisonDriversCmd.MarkFlagRequired("metric-for-comparison")
	})
	ceCmd.AddCommand(ce_getCostComparisonDriversCmd)
}
