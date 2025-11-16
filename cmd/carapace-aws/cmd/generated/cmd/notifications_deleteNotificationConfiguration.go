package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_deleteNotificationConfigurationCmd = &cobra.Command{
	Use:   "delete-notification-configuration",
	Short: "Deletes a `NotificationConfiguration`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_deleteNotificationConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(notifications_deleteNotificationConfigurationCmd).Standalone()

		notifications_deleteNotificationConfigurationCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the `NotificationConfiguration` to delete.")
		notifications_deleteNotificationConfigurationCmd.MarkFlagRequired("arn")
	})
	notificationsCmd.AddCommand(notifications_deleteNotificationConfigurationCmd)
}
