package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarNotifications_createNotificationRuleCmd = &cobra.Command{
	Use:   "create-notification-rule",
	Short: "Creates a notification rule for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarNotifications_createNotificationRuleCmd).Standalone()

	codestarNotifications_createNotificationRuleCmd.Flags().String("client-request-token", "", "A unique, client-generated idempotency token that, when provided in a request, ensures the request cannot be repeated with a changed parameter.")
	codestarNotifications_createNotificationRuleCmd.Flags().String("detail-type", "", "The level of detail to include in the notifications for this resource.")
	codestarNotifications_createNotificationRuleCmd.Flags().String("event-type-ids", "", "A list of event types associated with this notification rule.")
	codestarNotifications_createNotificationRuleCmd.Flags().String("name", "", "The name for the notification rule.")
	codestarNotifications_createNotificationRuleCmd.Flags().String("resource", "", "The Amazon Resource Name (ARN) of the resource to associate with the notification rule.")
	codestarNotifications_createNotificationRuleCmd.Flags().String("status", "", "The status of the notification rule.")
	codestarNotifications_createNotificationRuleCmd.Flags().String("tags", "", "A list of tags to apply to this notification rule.")
	codestarNotifications_createNotificationRuleCmd.Flags().String("targets", "", "A list of Amazon Resource Names (ARNs) of Amazon Simple Notification Service topics and Amazon Q Developer in chat applications clients to associate with the notification rule.")
	codestarNotifications_createNotificationRuleCmd.MarkFlagRequired("detail-type")
	codestarNotifications_createNotificationRuleCmd.MarkFlagRequired("event-type-ids")
	codestarNotifications_createNotificationRuleCmd.MarkFlagRequired("name")
	codestarNotifications_createNotificationRuleCmd.MarkFlagRequired("resource")
	codestarNotifications_createNotificationRuleCmd.MarkFlagRequired("targets")
	codestarNotificationsCmd.AddCommand(codestarNotifications_createNotificationRuleCmd)
}
