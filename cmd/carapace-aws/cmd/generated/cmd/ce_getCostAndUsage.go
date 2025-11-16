package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_getCostAndUsageCmd = &cobra.Command{
	Use:   "get-cost-and-usage",
	Short: "Retrieves cost and usage metrics for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_getCostAndUsageCmd).Standalone()

	ce_getCostAndUsageCmd.Flags().String("billing-view-arn", "", "The Amazon Resource Name (ARN) that uniquely identifies a specific billing view.")
	ce_getCostAndUsageCmd.Flags().String("filter", "", "Filters Amazon Web Services costs by different dimensions.")
	ce_getCostAndUsageCmd.Flags().String("granularity", "", "Sets the Amazon Web Services cost granularity to `MONTHLY` or `DAILY`, or `HOURLY`.")
	ce_getCostAndUsageCmd.Flags().String("group-by", "", "You can group Amazon Web Services costs using up to two different groups, either dimensions, tag keys, cost categories, or any two group by types.")
	ce_getCostAndUsageCmd.Flags().String("metrics", "", "Which metrics are returned in the query.")
	ce_getCostAndUsageCmd.Flags().String("next-page-token", "", "The token to retrieve the next set of results.")
	ce_getCostAndUsageCmd.Flags().String("time-period", "", "Sets the start date and end date for retrieving Amazon Web Services costs.")
	ce_getCostAndUsageCmd.MarkFlagRequired("granularity")
	ce_getCostAndUsageCmd.MarkFlagRequired("metrics")
	ce_getCostAndUsageCmd.MarkFlagRequired("time-period")
	ceCmd.AddCommand(ce_getCostAndUsageCmd)
}
