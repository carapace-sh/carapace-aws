package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_listNotificationHubsCmd = &cobra.Command{
	Use:   "list-notification-hubs",
	Short: "Returns a list of `NotificationHubs`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_listNotificationHubsCmd).Standalone()

	notifications_listNotificationHubsCmd.Flags().String("max-results", "", "The maximum number of records to list in a single response.")
	notifications_listNotificationHubsCmd.Flags().String("next-token", "", "A pagination token.")
	notificationsCmd.AddCommand(notifications_listNotificationHubsCmd)
}
