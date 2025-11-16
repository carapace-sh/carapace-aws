package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_getNotificationChannelCmd = &cobra.Command{
	Use:   "get-notification-channel",
	Short: "Information about the Amazon Simple Notification Service (SNS) topic that is used to record Firewall Manager SNS logs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_getNotificationChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fms_getNotificationChannelCmd).Standalone()

	})
	fmsCmd.AddCommand(fms_getNotificationChannelCmd)
}
