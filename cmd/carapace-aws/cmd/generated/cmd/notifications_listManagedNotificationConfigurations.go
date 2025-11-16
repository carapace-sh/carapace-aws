package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_listManagedNotificationConfigurationsCmd = &cobra.Command{
	Use:   "list-managed-notification-configurations",
	Short: "Returns a list of Managed Notification Configurations according to specified filters, ordered by creation time in reverse chronological order (newest first).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_listManagedNotificationConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(notifications_listManagedNotificationConfigurationsCmd).Standalone()

		notifications_listManagedNotificationConfigurationsCmd.Flags().String("channel-identifier", "", "The identifier or ARN of the notification channel to filter configurations by.")
		notifications_listManagedNotificationConfigurationsCmd.Flags().String("max-results", "", "The maximum number of results to be returned in this call.")
		notifications_listManagedNotificationConfigurationsCmd.Flags().String("next-token", "", "The start token for paginated calls.")
	})
	notificationsCmd.AddCommand(notifications_listManagedNotificationConfigurationsCmd)
}
