package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_getChannelMessageCmd = &cobra.Command{
	Use:   "get-channel-message",
	Short: "Gets the full details of a channel message.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_getChannelMessageCmd).Standalone()

	chimeSdkMessaging_getChannelMessageCmd.Flags().String("channel-arn", "", "The ARN of the channel.")
	chimeSdkMessaging_getChannelMessageCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
	chimeSdkMessaging_getChannelMessageCmd.Flags().String("message-id", "", "The ID of the message.")
	chimeSdkMessaging_getChannelMessageCmd.Flags().String("sub-channel-id", "", "The ID of the SubChannel in the request.")
	chimeSdkMessaging_getChannelMessageCmd.MarkFlagRequired("channel-arn")
	chimeSdkMessaging_getChannelMessageCmd.MarkFlagRequired("chime-bearer")
	chimeSdkMessaging_getChannelMessageCmd.MarkFlagRequired("message-id")
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_getChannelMessageCmd)
}
