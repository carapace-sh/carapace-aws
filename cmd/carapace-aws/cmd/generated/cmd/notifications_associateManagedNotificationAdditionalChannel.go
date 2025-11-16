package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_associateManagedNotificationAdditionalChannelCmd = &cobra.Command{
	Use:   "associate-managed-notification-additional-channel",
	Short: "Associates an additional Channel with a particular `ManagedNotificationConfiguration`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_associateManagedNotificationAdditionalChannelCmd).Standalone()

	notifications_associateManagedNotificationAdditionalChannelCmd.Flags().String("channel-arn", "", "The Amazon Resource Name (ARN) of the Channel to associate with the `ManagedNotificationConfiguration`.")
	notifications_associateManagedNotificationAdditionalChannelCmd.Flags().String("managed-notification-configuration-arn", "", "The Amazon Resource Name (ARN) of the `ManagedNotificationConfiguration` to associate with the additional Channel.")
	notifications_associateManagedNotificationAdditionalChannelCmd.MarkFlagRequired("channel-arn")
	notifications_associateManagedNotificationAdditionalChannelCmd.MarkFlagRequired("managed-notification-configuration-arn")
	notificationsCmd.AddCommand(notifications_associateManagedNotificationAdditionalChannelCmd)
}
