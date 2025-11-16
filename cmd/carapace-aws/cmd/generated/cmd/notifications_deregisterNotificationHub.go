package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_deregisterNotificationHubCmd = &cobra.Command{
	Use:   "deregister-notification-hub",
	Short: "Deregisters a `NotificationConfiguration` in the specified Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_deregisterNotificationHubCmd).Standalone()

	notifications_deregisterNotificationHubCmd.Flags().String("notification-hub-region", "", "The `NotificationConfiguration` Region.")
	notifications_deregisterNotificationHubCmd.MarkFlagRequired("notification-hub-region")
	notificationsCmd.AddCommand(notifications_deregisterNotificationHubCmd)
}
