package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_deleteChannelModeratorCmd = &cobra.Command{
	Use:   "delete-channel-moderator",
	Short: "Deletes a channel moderator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_deleteChannelModeratorCmd).Standalone()

	chimeSdkMessaging_deleteChannelModeratorCmd.Flags().String("channel-arn", "", "The ARN of the channel.")
	chimeSdkMessaging_deleteChannelModeratorCmd.Flags().String("channel-moderator-arn", "", "The `AppInstanceUserArn` of the moderator being deleted.")
	chimeSdkMessaging_deleteChannelModeratorCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
	chimeSdkMessaging_deleteChannelModeratorCmd.MarkFlagRequired("channel-arn")
	chimeSdkMessaging_deleteChannelModeratorCmd.MarkFlagRequired("channel-moderator-arn")
	chimeSdkMessaging_deleteChannelModeratorCmd.MarkFlagRequired("chime-bearer")
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_deleteChannelModeratorCmd)
}
