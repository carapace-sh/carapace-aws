package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_disassociateManagedNotificationAdditionalChannelCmd = &cobra.Command{
	Use:   "disassociate-managed-notification-additional-channel",
	Short: "Disassociates an additional Channel from a particular `ManagedNotificationConfiguration`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_disassociateManagedNotificationAdditionalChannelCmd).Standalone()

	notifications_disassociateManagedNotificationAdditionalChannelCmd.Flags().String("channel-arn", "", "The Amazon Resource Name (ARN) of the Channel to associate with the `ManagedNotificationConfiguration`.")
	notifications_disassociateManagedNotificationAdditionalChannelCmd.Flags().String("managed-notification-configuration-arn", "", "The Amazon Resource Name (ARN) of the Managed Notification Configuration to associate with the additional Channel.")
	notifications_disassociateManagedNotificationAdditionalChannelCmd.MarkFlagRequired("channel-arn")
	notifications_disassociateManagedNotificationAdditionalChannelCmd.MarkFlagRequired("managed-notification-configuration-arn")
	notificationsCmd.AddCommand(notifications_disassociateManagedNotificationAdditionalChannelCmd)
}
