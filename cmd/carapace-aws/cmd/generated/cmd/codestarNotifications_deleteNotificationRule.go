package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarNotifications_deleteNotificationRuleCmd = &cobra.Command{
	Use:   "delete-notification-rule",
	Short: "Deletes a notification rule for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarNotifications_deleteNotificationRuleCmd).Standalone()

	codestarNotifications_deleteNotificationRuleCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the notification rule you want to delete.")
	codestarNotifications_deleteNotificationRuleCmd.MarkFlagRequired("arn")
	codestarNotificationsCmd.AddCommand(codestarNotifications_deleteNotificationRuleCmd)
}
