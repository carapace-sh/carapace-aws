package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_updateSubscriptionCmd = &cobra.Command{
	Use:   "update-subscription",
	Short: "Updates the pricing tier for an Amazon Q Business subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_updateSubscriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_updateSubscriptionCmd).Standalone()

		qbusiness_updateSubscriptionCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application where the subscription update should take effect.")
		qbusiness_updateSubscriptionCmd.Flags().String("subscription-id", "", "The identifier of the Amazon Q Business subscription to be updated.")
		qbusiness_updateSubscriptionCmd.Flags().String("type", "", "The type of the Amazon Q Business subscription to be updated.")
		qbusiness_updateSubscriptionCmd.MarkFlagRequired("application-id")
		qbusiness_updateSubscriptionCmd.MarkFlagRequired("subscription-id")
		qbusiness_updateSubscriptionCmd.MarkFlagRequired("type")
	})
	qbusinessCmd.AddCommand(qbusiness_updateSubscriptionCmd)
}
