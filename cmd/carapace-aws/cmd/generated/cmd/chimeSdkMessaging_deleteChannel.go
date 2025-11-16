package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_deleteChannelCmd = &cobra.Command{
	Use:   "delete-channel",
	Short: "Immediately makes a channel and its memberships inaccessible and marks them for deletion.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_deleteChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMessaging_deleteChannelCmd).Standalone()

		chimeSdkMessaging_deleteChannelCmd.Flags().String("channel-arn", "", "The ARN of the channel being deleted.")
		chimeSdkMessaging_deleteChannelCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
		chimeSdkMessaging_deleteChannelCmd.MarkFlagRequired("channel-arn")
		chimeSdkMessaging_deleteChannelCmd.MarkFlagRequired("chime-bearer")
	})
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_deleteChannelCmd)
}
