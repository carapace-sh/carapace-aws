package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_listChannelMessagesCmd = &cobra.Command{
	Use:   "list-channel-messages",
	Short: "List all the messages in a channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_listChannelMessagesCmd).Standalone()

	chimeSdkMessaging_listChannelMessagesCmd.Flags().String("channel-arn", "", "The ARN of the channel.")
	chimeSdkMessaging_listChannelMessagesCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
	chimeSdkMessaging_listChannelMessagesCmd.Flags().String("max-results", "", "The maximum number of messages that you want returned.")
	chimeSdkMessaging_listChannelMessagesCmd.Flags().String("next-token", "", "The token passed by previous API calls until all requested messages are returned.")
	chimeSdkMessaging_listChannelMessagesCmd.Flags().String("not-after", "", "The final or ending time stamp for your requested messages.")
	chimeSdkMessaging_listChannelMessagesCmd.Flags().String("not-before", "", "The initial or starting time stamp for your requested messages.")
	chimeSdkMessaging_listChannelMessagesCmd.Flags().String("sort-order", "", "The order in which you want messages sorted.")
	chimeSdkMessaging_listChannelMessagesCmd.Flags().String("sub-channel-id", "", "The ID of the SubChannel in the request.")
	chimeSdkMessaging_listChannelMessagesCmd.MarkFlagRequired("channel-arn")
	chimeSdkMessaging_listChannelMessagesCmd.MarkFlagRequired("chime-bearer")
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_listChannelMessagesCmd)
}
