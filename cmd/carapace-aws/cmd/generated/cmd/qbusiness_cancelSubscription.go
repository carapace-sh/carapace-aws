package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_cancelSubscriptionCmd = &cobra.Command{
	Use:   "cancel-subscription",
	Short: "Unsubscribes a user or a group from their pricing tier in an Amazon Q Business application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_cancelSubscriptionCmd).Standalone()

	qbusiness_cancelSubscriptionCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application for which the subscription is being cancelled.")
	qbusiness_cancelSubscriptionCmd.Flags().String("subscription-id", "", "The identifier of the Amazon Q Business subscription being cancelled.")
	qbusiness_cancelSubscriptionCmd.MarkFlagRequired("application-id")
	qbusiness_cancelSubscriptionCmd.MarkFlagRequired("subscription-id")
	qbusinessCmd.AddCommand(qbusiness_cancelSubscriptionCmd)
}
