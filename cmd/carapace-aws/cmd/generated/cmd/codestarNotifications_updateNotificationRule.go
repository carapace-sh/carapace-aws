package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarNotifications_updateNotificationRuleCmd = &cobra.Command{
	Use:   "update-notification-rule",
	Short: "Updates a notification rule for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarNotifications_updateNotificationRuleCmd).Standalone()

	codestarNotifications_updateNotificationRuleCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the notification rule.")
	codestarNotifications_updateNotificationRuleCmd.Flags().String("detail-type", "", "The level of detail to include in the notifications for this resource.")
	codestarNotifications_updateNotificationRuleCmd.Flags().String("event-type-ids", "", "A list of event types associated with this notification rule.")
	codestarNotifications_updateNotificationRuleCmd.Flags().String("name", "", "The name of the notification rule.")
	codestarNotifications_updateNotificationRuleCmd.Flags().String("status", "", "The status of the notification rule.")
	codestarNotifications_updateNotificationRuleCmd.Flags().String("targets", "", "The address and type of the targets to receive notifications from this notification rule.")
	codestarNotifications_updateNotificationRuleCmd.MarkFlagRequired("arn")
	codestarNotificationsCmd.AddCommand(codestarNotifications_updateNotificationRuleCmd)
}
