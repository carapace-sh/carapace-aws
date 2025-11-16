package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_listManagedNotificationChildEventsCmd = &cobra.Command{
	Use:   "list-managed-notification-child-events",
	Short: "Returns a list of `ManagedNotificationChildEvents` for a specified aggregate `ManagedNotificationEvent`, ordered by creation time in reverse chronological order (newest first).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_listManagedNotificationChildEventsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(notifications_listManagedNotificationChildEventsCmd).Standalone()

		notifications_listManagedNotificationChildEventsCmd.Flags().String("aggregate-managed-notification-event-arn", "", "The Amazon Resource Name (ARN) of the `ManagedNotificationEvent`.")
		notifications_listManagedNotificationChildEventsCmd.Flags().String("end-time", "", "Latest time of events to return from this call.")
		notifications_listManagedNotificationChildEventsCmd.Flags().String("locale", "", "The locale code of the language used for the retrieved `NotificationEvent`.")
		notifications_listManagedNotificationChildEventsCmd.Flags().String("max-results", "", "The maximum number of results to be returned in this call.")
		notifications_listManagedNotificationChildEventsCmd.Flags().String("next-token", "", "The start token for paginated calls.")
		notifications_listManagedNotificationChildEventsCmd.Flags().String("organizational-unit-id", "", "The identifier of the Amazon Web Services Organizations organizational unit (OU) associated with the Managed Notification Child Events.")
		notifications_listManagedNotificationChildEventsCmd.Flags().String("related-account", "", "The Amazon Web Services account ID associated with the Managed Notification Child Events.")
		notifications_listManagedNotificationChildEventsCmd.Flags().String("start-time", "", "The earliest time of events to return from this call.")
		notifications_listManagedNotificationChildEventsCmd.MarkFlagRequired("aggregate-managed-notification-event-arn")
	})
	notificationsCmd.AddCommand(notifications_listManagedNotificationChildEventsCmd)
}
