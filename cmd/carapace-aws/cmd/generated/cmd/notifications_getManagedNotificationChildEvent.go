package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_getManagedNotificationChildEventCmd = &cobra.Command{
	Use:   "get-managed-notification-child-event",
	Short: "Returns the child event of a specific given `ManagedNotificationEvent`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_getManagedNotificationChildEventCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(notifications_getManagedNotificationChildEventCmd).Standalone()

		notifications_getManagedNotificationChildEventCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the `ManagedNotificationChildEvent` to return.")
		notifications_getManagedNotificationChildEventCmd.Flags().String("locale", "", "The locale code of the language used for the retrieved `ManagedNotificationChildEvent`.")
		notifications_getManagedNotificationChildEventCmd.MarkFlagRequired("arn")
	})
	notificationsCmd.AddCommand(notifications_getManagedNotificationChildEventCmd)
}
