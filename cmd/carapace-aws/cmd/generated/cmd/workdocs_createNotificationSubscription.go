package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_createNotificationSubscriptionCmd = &cobra.Command{
	Use:   "create-notification-subscription",
	Short: "Configure Amazon WorkDocs to use Amazon SNS notifications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_createNotificationSubscriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workdocs_createNotificationSubscriptionCmd).Standalone()

		workdocs_createNotificationSubscriptionCmd.Flags().String("endpoint", "", "The endpoint to receive the notifications.")
		workdocs_createNotificationSubscriptionCmd.Flags().String("organization-id", "", "The ID of the organization.")
		workdocs_createNotificationSubscriptionCmd.Flags().String("protocol", "", "The protocol to use.")
		workdocs_createNotificationSubscriptionCmd.Flags().String("subscription-type", "", "The notification type.")
		workdocs_createNotificationSubscriptionCmd.MarkFlagRequired("endpoint")
		workdocs_createNotificationSubscriptionCmd.MarkFlagRequired("organization-id")
		workdocs_createNotificationSubscriptionCmd.MarkFlagRequired("protocol")
		workdocs_createNotificationSubscriptionCmd.MarkFlagRequired("subscription-type")
	})
	workdocsCmd.AddCommand(workdocs_createNotificationSubscriptionCmd)
}
