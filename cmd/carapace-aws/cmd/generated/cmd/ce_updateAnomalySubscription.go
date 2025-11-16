package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_updateAnomalySubscriptionCmd = &cobra.Command{
	Use:   "update-anomaly-subscription",
	Short: "Updates an existing cost anomaly subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_updateAnomalySubscriptionCmd).Standalone()

	ce_updateAnomalySubscriptionCmd.Flags().String("frequency", "", "The update to the frequency value that subscribers receive notifications.")
	ce_updateAnomalySubscriptionCmd.Flags().String("monitor-arn-list", "", "A list of cost anomaly monitor ARNs.")
	ce_updateAnomalySubscriptionCmd.Flags().String("subscribers", "", "The update to the subscriber list.")
	ce_updateAnomalySubscriptionCmd.Flags().String("subscription-arn", "", "A cost anomaly subscription Amazon Resource Name (ARN).")
	ce_updateAnomalySubscriptionCmd.Flags().String("subscription-name", "", "The new name of the subscription.")
	ce_updateAnomalySubscriptionCmd.Flags().String("threshold", "", "(deprecated)")
	ce_updateAnomalySubscriptionCmd.Flags().String("threshold-expression", "", "The update to the [Expression](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html) object used to specify the anomalies that you want to generate alerts for.")
	ce_updateAnomalySubscriptionCmd.MarkFlagRequired("subscription-arn")
	ceCmd.AddCommand(ce_updateAnomalySubscriptionCmd)
}
