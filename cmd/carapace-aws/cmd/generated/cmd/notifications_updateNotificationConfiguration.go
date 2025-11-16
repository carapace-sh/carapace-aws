package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_updateNotificationConfigurationCmd = &cobra.Command{
	Use:   "update-notification-configuration",
	Short: "Updates a `NotificationConfiguration`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_updateNotificationConfigurationCmd).Standalone()

	notifications_updateNotificationConfigurationCmd.Flags().String("aggregation-duration", "", "The aggregation preference of the `NotificationConfiguration`.")
	notifications_updateNotificationConfigurationCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) used to update the `NotificationConfiguration`.")
	notifications_updateNotificationConfigurationCmd.Flags().String("description", "", "The description of the `NotificationConfiguration`.")
	notifications_updateNotificationConfigurationCmd.Flags().String("name", "", "The name of the `NotificationConfiguration`.")
	notifications_updateNotificationConfigurationCmd.MarkFlagRequired("arn")
	notificationsCmd.AddCommand(notifications_updateNotificationConfigurationCmd)
}
