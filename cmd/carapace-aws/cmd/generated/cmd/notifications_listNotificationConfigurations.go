package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_listNotificationConfigurationsCmd = &cobra.Command{
	Use:   "list-notification-configurations",
	Short: "Returns a list of abbreviated `NotificationConfigurations` according to specified filters, in reverse chronological order (newest first).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_listNotificationConfigurationsCmd).Standalone()

	notifications_listNotificationConfigurationsCmd.Flags().String("channel-arn", "", "The Amazon Resource Name (ARN) of the Channel to match.")
	notifications_listNotificationConfigurationsCmd.Flags().String("event-rule-source", "", "The matched event source.")
	notifications_listNotificationConfigurationsCmd.Flags().String("max-results", "", "The maximum number of results to be returned in this call.")
	notifications_listNotificationConfigurationsCmd.Flags().String("next-token", "", "The start token for paginated calls.")
	notifications_listNotificationConfigurationsCmd.Flags().String("status", "", "The `NotificationConfiguration` status to match.")
	notifications_listNotificationConfigurationsCmd.Flags().String("subtype", "", "The subtype used to filter the notification configurations in the request.")
	notificationsCmd.AddCommand(notifications_listNotificationConfigurationsCmd)
}
