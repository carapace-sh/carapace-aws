package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_getAnomalySubscriptionsCmd = &cobra.Command{
	Use:   "get-anomaly-subscriptions",
	Short: "Retrieves the cost anomaly subscription objects for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_getAnomalySubscriptionsCmd).Standalone()

	ce_getAnomalySubscriptionsCmd.Flags().String("max-results", "", "The number of entries a paginated response contains.")
	ce_getAnomalySubscriptionsCmd.Flags().String("monitor-arn", "", "Cost anomaly monitor ARNs.")
	ce_getAnomalySubscriptionsCmd.Flags().String("next-page-token", "", "The token to retrieve the next set of results.")
	ce_getAnomalySubscriptionsCmd.Flags().String("subscription-arn-list", "", "A list of cost anomaly subscription ARNs.")
	ceCmd.AddCommand(ce_getAnomalySubscriptionsCmd)
}
