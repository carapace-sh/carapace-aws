package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_createNotificationConfigurationCmd = &cobra.Command{
	Use:   "create-notification-configuration",
	Short: "Creates a new `NotificationConfiguration`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_createNotificationConfigurationCmd).Standalone()

	notifications_createNotificationConfigurationCmd.Flags().String("aggregation-duration", "", "The aggregation preference of the `NotificationConfiguration`.")
	notifications_createNotificationConfigurationCmd.Flags().String("description", "", "The description of the `NotificationConfiguration`.")
	notifications_createNotificationConfigurationCmd.Flags().String("name", "", "The name of the `NotificationConfiguration`.")
	notifications_createNotificationConfigurationCmd.Flags().String("tags", "", "A map of tags assigned to a resource.")
	notifications_createNotificationConfigurationCmd.MarkFlagRequired("description")
	notifications_createNotificationConfigurationCmd.MarkFlagRequired("name")
	notificationsCmd.AddCommand(notifications_createNotificationConfigurationCmd)
}
