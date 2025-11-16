package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_updateChannelMessageCmd = &cobra.Command{
	Use:   "update-channel-message",
	Short: "Updates the content of a message.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_updateChannelMessageCmd).Standalone()

	chimeSdkMessaging_updateChannelMessageCmd.Flags().String("channel-arn", "", "The ARN of the channel.")
	chimeSdkMessaging_updateChannelMessageCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
	chimeSdkMessaging_updateChannelMessageCmd.Flags().String("content", "", "The content of the channel message.")
	chimeSdkMessaging_updateChannelMessageCmd.Flags().String("content-type", "", "The content type of the channel message.")
	chimeSdkMessaging_updateChannelMessageCmd.Flags().String("message-id", "", "The ID string of the message being updated.")
	chimeSdkMessaging_updateChannelMessageCmd.Flags().String("metadata", "", "The metadata of the message being updated.")
	chimeSdkMessaging_updateChannelMessageCmd.Flags().String("sub-channel-id", "", "The ID of the SubChannel in the request.")
	chimeSdkMessaging_updateChannelMessageCmd.MarkFlagRequired("channel-arn")
	chimeSdkMessaging_updateChannelMessageCmd.MarkFlagRequired("chime-bearer")
	chimeSdkMessaging_updateChannelMessageCmd.MarkFlagRequired("content")
	chimeSdkMessaging_updateChannelMessageCmd.MarkFlagRequired("message-id")
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_updateChannelMessageCmd)
}
