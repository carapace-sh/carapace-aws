package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_deleteNotificationSubscriptionCmd = &cobra.Command{
	Use:   "delete-notification-subscription",
	Short: "Deletes the specified subscription from the specified organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_deleteNotificationSubscriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workdocs_deleteNotificationSubscriptionCmd).Standalone()

		workdocs_deleteNotificationSubscriptionCmd.Flags().String("organization-id", "", "The ID of the organization.")
		workdocs_deleteNotificationSubscriptionCmd.Flags().String("subscription-id", "", "The ID of the subscription.")
		workdocs_deleteNotificationSubscriptionCmd.MarkFlagRequired("organization-id")
		workdocs_deleteNotificationSubscriptionCmd.MarkFlagRequired("subscription-id")
	})
	workdocsCmd.AddCommand(workdocs_deleteNotificationSubscriptionCmd)
}
