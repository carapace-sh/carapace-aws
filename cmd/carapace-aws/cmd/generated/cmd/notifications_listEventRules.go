package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_listEventRulesCmd = &cobra.Command{
	Use:   "list-event-rules",
	Short: "Returns a list of `EventRules` according to specified filters, in reverse chronological order (newest first).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_listEventRulesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(notifications_listEventRulesCmd).Standalone()

		notifications_listEventRulesCmd.Flags().String("max-results", "", "The maximum number of results to be returned in this call.")
		notifications_listEventRulesCmd.Flags().String("next-token", "", "The start token for paginated calls.")
		notifications_listEventRulesCmd.Flags().String("notification-configuration-arn", "", "The Amazon Resource Name (ARN) of the `NotificationConfiguration`.")
		notifications_listEventRulesCmd.MarkFlagRequired("notification-configuration-arn")
	})
	notificationsCmd.AddCommand(notifications_listEventRulesCmd)
}
