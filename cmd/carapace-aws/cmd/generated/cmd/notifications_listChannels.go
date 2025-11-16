package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_listChannelsCmd = &cobra.Command{
	Use:   "list-channels",
	Short: "Returns a list of Channels for a `NotificationConfiguration`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_listChannelsCmd).Standalone()

	notifications_listChannelsCmd.Flags().String("max-results", "", "The maximum number of results to be returned in this call.")
	notifications_listChannelsCmd.Flags().String("next-token", "", "The start token for paginated calls.")
	notifications_listChannelsCmd.Flags().String("notification-configuration-arn", "", "The Amazon Resource Name (ARN) of the `NotificationConfiguration`.")
	notifications_listChannelsCmd.MarkFlagRequired("notification-configuration-arn")
	notificationsCmd.AddCommand(notifications_listChannelsCmd)
}
