package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_disassociateChannelCmd = &cobra.Command{
	Use:   "disassociate-channel",
	Short: "Disassociates a Channel from a specified `NotificationConfiguration`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_disassociateChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(notifications_disassociateChannelCmd).Standalone()

		notifications_disassociateChannelCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the Channel to disassociate.")
		notifications_disassociateChannelCmd.Flags().String("notification-configuration-arn", "", "The ARN of the `NotificationConfiguration` to disassociate.")
		notifications_disassociateChannelCmd.MarkFlagRequired("arn")
		notifications_disassociateChannelCmd.MarkFlagRequired("notification-configuration-arn")
	})
	notificationsCmd.AddCommand(notifications_disassociateChannelCmd)
}
