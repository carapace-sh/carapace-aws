package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_createAnomalySubscriptionCmd = &cobra.Command{
	Use:   "create-anomaly-subscription",
	Short: "Adds an alert subscription to a cost anomaly detection monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_createAnomalySubscriptionCmd).Standalone()

	ce_createAnomalySubscriptionCmd.Flags().String("anomaly-subscription", "", "The cost anomaly subscription object that you want to create.")
	ce_createAnomalySubscriptionCmd.Flags().String("resource-tags", "", "An optional list of tags to associate with the specified [`AnomalySubscription`](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_AnomalySubscription.html) .")
	ce_createAnomalySubscriptionCmd.MarkFlagRequired("anomaly-subscription")
	ceCmd.AddCommand(ce_createAnomalySubscriptionCmd)
}
