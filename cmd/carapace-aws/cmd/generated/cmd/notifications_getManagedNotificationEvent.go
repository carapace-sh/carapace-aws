package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_getManagedNotificationEventCmd = &cobra.Command{
	Use:   "get-managed-notification-event",
	Short: "Returns a specified `ManagedNotificationEvent`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_getManagedNotificationEventCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(notifications_getManagedNotificationEventCmd).Standalone()

		notifications_getManagedNotificationEventCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the `ManagedNotificationEvent` to return.")
		notifications_getManagedNotificationEventCmd.Flags().String("locale", "", "The locale code of the language used for the retrieved `ManagedNotificationEvent`.")
		notifications_getManagedNotificationEventCmd.MarkFlagRequired("arn")
	})
	notificationsCmd.AddCommand(notifications_getManagedNotificationEventCmd)
}
