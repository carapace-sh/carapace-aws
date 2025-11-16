package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_listManagedNotificationEventsCmd = &cobra.Command{
	Use:   "list-managed-notification-events",
	Short: "Returns a list of Managed Notification Events according to specified filters, ordered by creation time in reverse chronological order (newest first).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_listManagedNotificationEventsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(notifications_listManagedNotificationEventsCmd).Standalone()

		notifications_listManagedNotificationEventsCmd.Flags().String("end-time", "", "Latest time of events to return from this call.")
		notifications_listManagedNotificationEventsCmd.Flags().String("locale", "", "The locale code of the language used for the retrieved NotificationEvent.")
		notifications_listManagedNotificationEventsCmd.Flags().String("max-results", "", "The maximum number of results to be returned in this call.")
		notifications_listManagedNotificationEventsCmd.Flags().String("next-token", "", "The start token for paginated calls.")
		notifications_listManagedNotificationEventsCmd.Flags().String("organizational-unit-id", "", "The Organizational Unit Id that an Amazon Web Services account belongs to.")
		notifications_listManagedNotificationEventsCmd.Flags().String("related-account", "", "The Amazon Web Services account ID associated with the Managed Notification Events.")
		notifications_listManagedNotificationEventsCmd.Flags().String("source", "", "The Amazon Web Services service the event originates from.")
		notifications_listManagedNotificationEventsCmd.Flags().String("start-time", "", "The earliest time of events to return from this call.")
	})
	notificationsCmd.AddCommand(notifications_listManagedNotificationEventsCmd)
}
