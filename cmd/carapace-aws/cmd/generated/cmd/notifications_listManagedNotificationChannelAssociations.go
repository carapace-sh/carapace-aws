package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_listManagedNotificationChannelAssociationsCmd = &cobra.Command{
	Use:   "list-managed-notification-channel-associations",
	Short: "Returns a list of Account contacts and Channels associated with a `ManagedNotificationConfiguration`, in paginated format.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_listManagedNotificationChannelAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(notifications_listManagedNotificationChannelAssociationsCmd).Standalone()

		notifications_listManagedNotificationChannelAssociationsCmd.Flags().String("managed-notification-configuration-arn", "", "The Amazon Resource Name (ARN) of the `ManagedNotificationConfiguration` to match.")
		notifications_listManagedNotificationChannelAssociationsCmd.Flags().String("max-results", "", "The maximum number of results to be returned in this call.")
		notifications_listManagedNotificationChannelAssociationsCmd.Flags().String("next-token", "", "The start token for paginated calls.")
		notifications_listManagedNotificationChannelAssociationsCmd.MarkFlagRequired("managed-notification-configuration-arn")
	})
	notificationsCmd.AddCommand(notifications_listManagedNotificationChannelAssociationsCmd)
}
