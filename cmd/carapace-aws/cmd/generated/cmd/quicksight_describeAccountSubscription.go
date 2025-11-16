package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeAccountSubscriptionCmd = &cobra.Command{
	Use:   "describe-account-subscription",
	Short: "Use the DescribeAccountSubscription operation to receive a description of an Quick Sight account's subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeAccountSubscriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_describeAccountSubscriptionCmd).Standalone()

		quicksight_describeAccountSubscriptionCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID associated with your Quick Sight account.")
		quicksight_describeAccountSubscriptionCmd.MarkFlagRequired("aws-account-id")
	})
	quicksightCmd.AddCommand(quicksight_describeAccountSubscriptionCmd)
}
