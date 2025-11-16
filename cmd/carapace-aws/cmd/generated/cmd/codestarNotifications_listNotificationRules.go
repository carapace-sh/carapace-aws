package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarNotifications_listNotificationRulesCmd = &cobra.Command{
	Use:   "list-notification-rules",
	Short: "Returns a list of the notification rules for an Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarNotifications_listNotificationRulesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codestarNotifications_listNotificationRulesCmd).Standalone()

		codestarNotifications_listNotificationRulesCmd.Flags().String("filters", "", "The filters to use to return information by service or resource type.")
		codestarNotifications_listNotificationRulesCmd.Flags().String("max-results", "", "A non-negative integer used to limit the number of returned results.")
		codestarNotifications_listNotificationRulesCmd.Flags().String("next-token", "", "An enumeration token that, when provided in a request, returns the next batch of the results.")
	})
	codestarNotificationsCmd.AddCommand(codestarNotifications_listNotificationRulesCmd)
}
