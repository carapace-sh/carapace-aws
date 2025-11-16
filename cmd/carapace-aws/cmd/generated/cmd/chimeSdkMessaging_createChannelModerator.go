package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_createChannelModeratorCmd = &cobra.Command{
	Use:   "create-channel-moderator",
	Short: "Creates a new `ChannelModerator`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_createChannelModeratorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMessaging_createChannelModeratorCmd).Standalone()

		chimeSdkMessaging_createChannelModeratorCmd.Flags().String("channel-arn", "", "The ARN of the channel.")
		chimeSdkMessaging_createChannelModeratorCmd.Flags().String("channel-moderator-arn", "", "The `AppInstanceUserArn` of the moderator.")
		chimeSdkMessaging_createChannelModeratorCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
		chimeSdkMessaging_createChannelModeratorCmd.MarkFlagRequired("channel-arn")
		chimeSdkMessaging_createChannelModeratorCmd.MarkFlagRequired("channel-moderator-arn")
		chimeSdkMessaging_createChannelModeratorCmd.MarkFlagRequired("chime-bearer")
	})
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_createChannelModeratorCmd)
}
