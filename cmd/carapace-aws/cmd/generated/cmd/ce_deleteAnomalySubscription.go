package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_deleteAnomalySubscriptionCmd = &cobra.Command{
	Use:   "delete-anomaly-subscription",
	Short: "Deletes a cost anomaly subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_deleteAnomalySubscriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ce_deleteAnomalySubscriptionCmd).Standalone()

		ce_deleteAnomalySubscriptionCmd.Flags().String("subscription-arn", "", "The unique identifier of the cost anomaly subscription that you want to delete.")
		ce_deleteAnomalySubscriptionCmd.MarkFlagRequired("subscription-arn")
	})
	ceCmd.AddCommand(ce_deleteAnomalySubscriptionCmd)
}
