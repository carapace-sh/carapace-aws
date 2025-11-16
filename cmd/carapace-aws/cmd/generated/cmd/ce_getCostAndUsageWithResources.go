package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_getCostAndUsageWithResourcesCmd = &cobra.Command{
	Use:   "get-cost-and-usage-with-resources",
	Short: "Retrieves cost and usage metrics with resources for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_getCostAndUsageWithResourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ce_getCostAndUsageWithResourcesCmd).Standalone()

		ce_getCostAndUsageWithResourcesCmd.Flags().String("billing-view-arn", "", "The Amazon Resource Name (ARN) that uniquely identifies a specific billing view.")
		ce_getCostAndUsageWithResourcesCmd.Flags().String("filter", "", "Filters Amazon Web Services costs by different dimensions.")
		ce_getCostAndUsageWithResourcesCmd.Flags().String("granularity", "", "Sets the Amazon Web Services cost granularity to `MONTHLY`, `DAILY`, or `HOURLY`.")
		ce_getCostAndUsageWithResourcesCmd.Flags().String("group-by", "", "You can group Amazon Web Services costs using up to two different groups: `DIMENSION`, `TAG`, `COST_CATEGORY`.")
		ce_getCostAndUsageWithResourcesCmd.Flags().String("metrics", "", "Which metrics are returned in the query.")
		ce_getCostAndUsageWithResourcesCmd.Flags().String("next-page-token", "", "The token to retrieve the next set of results.")
		ce_getCostAndUsageWithResourcesCmd.Flags().String("time-period", "", "Sets the start and end dates for retrieving Amazon Web Services costs.")
		ce_getCostAndUsageWithResourcesCmd.MarkFlagRequired("filter")
		ce_getCostAndUsageWithResourcesCmd.MarkFlagRequired("granularity")
		ce_getCostAndUsageWithResourcesCmd.MarkFlagRequired("time-period")
	})
	ceCmd.AddCommand(ce_getCostAndUsageWithResourcesCmd)
}
