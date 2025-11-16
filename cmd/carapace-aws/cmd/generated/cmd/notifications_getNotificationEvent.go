package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_getNotificationEventCmd = &cobra.Command{
	Use:   "get-notification-event",
	Short: "Returns a specified `NotificationEvent`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_getNotificationEventCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(notifications_getNotificationEventCmd).Standalone()

		notifications_getNotificationEventCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the `NotificationEvent` to return.")
		notifications_getNotificationEventCmd.Flags().String("locale", "", "The locale code of the language used for the retrieved `NotificationEvent`.")
		notifications_getNotificationEventCmd.MarkFlagRequired("arn")
	})
	notificationsCmd.AddCommand(notifications_getNotificationEventCmd)
}
