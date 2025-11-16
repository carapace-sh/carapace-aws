package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_deleteNotificationChannelCmd = &cobra.Command{
	Use:   "delete-notification-channel",
	Short: "Deletes an Firewall Manager association with the IAM role and the Amazon Simple Notification Service (SNS) topic that is used to record Firewall Manager SNS logs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_deleteNotificationChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fms_deleteNotificationChannelCmd).Standalone()

	})
	fmsCmd.AddCommand(fms_deleteNotificationChannelCmd)
}
