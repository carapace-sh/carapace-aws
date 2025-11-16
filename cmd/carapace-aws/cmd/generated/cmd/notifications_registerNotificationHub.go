package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_registerNotificationHubCmd = &cobra.Command{
	Use:   "register-notification-hub",
	Short: "Registers a `NotificationConfiguration` in the specified Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_registerNotificationHubCmd).Standalone()

	notifications_registerNotificationHubCmd.Flags().String("notification-hub-region", "", "The Region of the `NotificationHub`.")
	notifications_registerNotificationHubCmd.MarkFlagRequired("notification-hub-region")
	notificationsCmd.AddCommand(notifications_registerNotificationHubCmd)
}
