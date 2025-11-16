package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_putNotificationChannelCmd = &cobra.Command{
	Use:   "put-notification-channel",
	Short: "Designates the IAM role and Amazon Simple Notification Service (SNS) topic that Firewall Manager uses to record SNS logs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_putNotificationChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fms_putNotificationChannelCmd).Standalone()

		fms_putNotificationChannelCmd.Flags().String("sns-role-name", "", "The Amazon Resource Name (ARN) of the IAM role that allows Amazon SNS to record Firewall Manager activity.")
		fms_putNotificationChannelCmd.Flags().String("sns-topic-arn", "", "The Amazon Resource Name (ARN) of the SNS topic that collects notifications from Firewall Manager.")
		fms_putNotificationChannelCmd.MarkFlagRequired("sns-role-name")
		fms_putNotificationChannelCmd.MarkFlagRequired("sns-topic-arn")
	})
	fmsCmd.AddCommand(fms_putNotificationChannelCmd)
}
