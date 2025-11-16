package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_getCostAndUsageComparisonsCmd = &cobra.Command{
	Use:   "get-cost-and-usage-comparisons",
	Short: "Retrieves cost and usage comparisons for your account between two periods within the last 13 months.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_getCostAndUsageComparisonsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ce_getCostAndUsageComparisonsCmd).Standalone()

		ce_getCostAndUsageComparisonsCmd.Flags().String("baseline-time-period", "", "The reference time period for comparison.")
		ce_getCostAndUsageComparisonsCmd.Flags().String("billing-view-arn", "", "The Amazon Resource Name (ARN) that uniquely identifies a specific billing view.")
		ce_getCostAndUsageComparisonsCmd.Flags().String("comparison-time-period", "", "The comparison time period for analysis.")
		ce_getCostAndUsageComparisonsCmd.Flags().String("filter", "", "")
		ce_getCostAndUsageComparisonsCmd.Flags().String("group-by", "", "You can group results using the attributes `DIMENSION`, `TAG`, and `COST_CATEGORY`.")
		ce_getCostAndUsageComparisonsCmd.Flags().String("max-results", "", "The maximum number of results that are returned for the request.")
		ce_getCostAndUsageComparisonsCmd.Flags().String("metric-for-comparison", "", "The cost and usage metric to compare.")
		ce_getCostAndUsageComparisonsCmd.Flags().String("next-page-token", "", "The token to retrieve the next set of paginated results.")
		ce_getCostAndUsageComparisonsCmd.MarkFlagRequired("baseline-time-period")
		ce_getCostAndUsageComparisonsCmd.MarkFlagRequired("comparison-time-period")
		ce_getCostAndUsageComparisonsCmd.MarkFlagRequired("metric-for-comparison")
	})
	ceCmd.AddCommand(ce_getCostAndUsageComparisonsCmd)
}
