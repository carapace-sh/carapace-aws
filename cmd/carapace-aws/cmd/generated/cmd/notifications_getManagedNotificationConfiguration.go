package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_getManagedNotificationConfigurationCmd = &cobra.Command{
	Use:   "get-managed-notification-configuration",
	Short: "Returns a specified `ManagedNotificationConfiguration`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_getManagedNotificationConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(notifications_getManagedNotificationConfigurationCmd).Standalone()

		notifications_getManagedNotificationConfigurationCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the `ManagedNotificationConfiguration` to return.")
		notifications_getManagedNotificationConfigurationCmd.MarkFlagRequired("arn")
	})
	notificationsCmd.AddCommand(notifications_getManagedNotificationConfigurationCmd)
}
