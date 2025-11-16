package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_deleteChannelMessageCmd = &cobra.Command{
	Use:   "delete-channel-message",
	Short: "Deletes a channel message.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_deleteChannelMessageCmd).Standalone()

	chimeSdkMessaging_deleteChannelMessageCmd.Flags().String("channel-arn", "", "The ARN of the channel.")
	chimeSdkMessaging_deleteChannelMessageCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
	chimeSdkMessaging_deleteChannelMessageCmd.Flags().String("message-id", "", "The ID of the message being deleted.")
	chimeSdkMessaging_deleteChannelMessageCmd.Flags().String("sub-channel-id", "", "The ID of the SubChannel in the request.")
	chimeSdkMessaging_deleteChannelMessageCmd.MarkFlagRequired("channel-arn")
	chimeSdkMessaging_deleteChannelMessageCmd.MarkFlagRequired("chime-bearer")
	chimeSdkMessaging_deleteChannelMessageCmd.MarkFlagRequired("message-id")
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_deleteChannelMessageCmd)
}
