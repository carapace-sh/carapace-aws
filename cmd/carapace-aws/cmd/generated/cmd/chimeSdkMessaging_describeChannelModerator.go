package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_describeChannelModeratorCmd = &cobra.Command{
	Use:   "describe-channel-moderator",
	Short: "Returns the full details of a single ChannelModerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_describeChannelModeratorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMessaging_describeChannelModeratorCmd).Standalone()

		chimeSdkMessaging_describeChannelModeratorCmd.Flags().String("channel-arn", "", "The ARN of the channel.")
		chimeSdkMessaging_describeChannelModeratorCmd.Flags().String("channel-moderator-arn", "", "The `AppInstanceUserArn` of the channel moderator.")
		chimeSdkMessaging_describeChannelModeratorCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
		chimeSdkMessaging_describeChannelModeratorCmd.MarkFlagRequired("channel-arn")
		chimeSdkMessaging_describeChannelModeratorCmd.MarkFlagRequired("channel-moderator-arn")
		chimeSdkMessaging_describeChannelModeratorCmd.MarkFlagRequired("chime-bearer")
	})
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_describeChannelModeratorCmd)
}
