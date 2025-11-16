package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarNotifications_describeNotificationRuleCmd = &cobra.Command{
	Use:   "describe-notification-rule",
	Short: "Returns information about a specified notification rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarNotifications_describeNotificationRuleCmd).Standalone()

	codestarNotifications_describeNotificationRuleCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the notification rule.")
	codestarNotifications_describeNotificationRuleCmd.MarkFlagRequired("arn")
	codestarNotificationsCmd.AddCommand(codestarNotifications_describeNotificationRuleCmd)
}
