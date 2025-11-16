package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_listNotificationEventsCmd = &cobra.Command{
	Use:   "list-notification-events",
	Short: "Returns a list of `NotificationEvents` according to specified filters, in reverse chronological order (newest first).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_listNotificationEventsCmd).Standalone()

	notifications_listNotificationEventsCmd.Flags().String("aggregate-notification-event-arn", "", "The Amazon Resource Name (ARN) of the `aggregatedNotificationEventArn` to match.")
	notifications_listNotificationEventsCmd.Flags().String("end-time", "", "Latest time of events to return from this call.")
	notifications_listNotificationEventsCmd.Flags().Bool("include-child-events", false, "Include aggregated child events in the result.")
	notifications_listNotificationEventsCmd.Flags().String("locale", "", "The locale code of the language used for the retrieved `NotificationEvent`.")
	notifications_listNotificationEventsCmd.Flags().String("max-results", "", "The maximum number of results to be returned in this call.")
	notifications_listNotificationEventsCmd.Flags().String("next-token", "", "The start token for paginated calls.")
	notifications_listNotificationEventsCmd.Flags().Bool("no-include-child-events", false, "Include aggregated child events in the result.")
	notifications_listNotificationEventsCmd.Flags().String("organizational-unit-id", "", "The unique identifier of the organizational unit used to filter notification events.")
	notifications_listNotificationEventsCmd.Flags().String("source", "", "The matched event source.")
	notifications_listNotificationEventsCmd.Flags().String("start-time", "", "The earliest time of events to return from this call.")
	notifications_listNotificationEventsCmd.Flag("no-include-child-events").Hidden = true
	notificationsCmd.AddCommand(notifications_listNotificationEventsCmd)
}
