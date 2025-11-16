package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_getNotificationConfigurationCmd = &cobra.Command{
	Use:   "get-notification-configuration",
	Short: "Returns a specified `NotificationConfiguration`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_getNotificationConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(notifications_getNotificationConfigurationCmd).Standalone()

		notifications_getNotificationConfigurationCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the `NotificationConfiguration` to return.")
		notifications_getNotificationConfigurationCmd.MarkFlagRequired("arn")
	})
	notificationsCmd.AddCommand(notifications_getNotificationConfigurationCmd)
}
