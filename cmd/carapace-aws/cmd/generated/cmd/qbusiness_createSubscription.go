package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_createSubscriptionCmd = &cobra.Command{
	Use:   "create-subscription",
	Short: "Subscribes an IAM Identity Center user or a group to a pricing tier for an Amazon Q Business application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_createSubscriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_createSubscriptionCmd).Standalone()

		qbusiness_createSubscriptionCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application the subscription should be added to.")
		qbusiness_createSubscriptionCmd.Flags().String("client-token", "", "A token that you provide to identify the request to create a subscription for your Amazon Q Business application.")
		qbusiness_createSubscriptionCmd.Flags().String("principal", "", "The IAM Identity Center `UserId` or `GroupId` of a user or group in the IAM Identity Center instance connected to the Amazon Q Business application.")
		qbusiness_createSubscriptionCmd.Flags().String("type", "", "The type of Amazon Q Business subscription you want to create.")
		qbusiness_createSubscriptionCmd.MarkFlagRequired("application-id")
		qbusiness_createSubscriptionCmd.MarkFlagRequired("principal")
		qbusiness_createSubscriptionCmd.MarkFlagRequired("type")
	})
	qbusinessCmd.AddCommand(qbusiness_createSubscriptionCmd)
}
