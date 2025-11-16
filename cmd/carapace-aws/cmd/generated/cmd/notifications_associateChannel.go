package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_associateChannelCmd = &cobra.Command{
	Use:   "associate-channel",
	Short: "Associates a delivery [Channel](https://docs.aws.amazon.com/notifications/latest/userguide/managing-delivery-channels.html) with a particular `NotificationConfiguration`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_associateChannelCmd).Standalone()

	notifications_associateChannelCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the Channel to associate with the `NotificationConfiguration`.")
	notifications_associateChannelCmd.Flags().String("notification-configuration-arn", "", "The ARN of the `NotificationConfiguration` to associate with the Channel.")
	notifications_associateChannelCmd.MarkFlagRequired("arn")
	notifications_associateChannelCmd.MarkFlagRequired("notification-configuration-arn")
	notificationsCmd.AddCommand(notifications_associateChannelCmd)
}
