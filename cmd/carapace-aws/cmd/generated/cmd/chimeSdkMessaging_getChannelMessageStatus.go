package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_getChannelMessageStatusCmd = &cobra.Command{
	Use:   "get-channel-message-status",
	Short: "Gets message status for a specified `messageId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_getChannelMessageStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMessaging_getChannelMessageStatusCmd).Standalone()

		chimeSdkMessaging_getChannelMessageStatusCmd.Flags().String("channel-arn", "", "The ARN of the channel")
		chimeSdkMessaging_getChannelMessageStatusCmd.Flags().String("chime-bearer", "", "The `AppInstanceUserArn` of the user making the API call.")
		chimeSdkMessaging_getChannelMessageStatusCmd.Flags().String("message-id", "", "The ID of the message.")
		chimeSdkMessaging_getChannelMessageStatusCmd.Flags().String("sub-channel-id", "", "The ID of the SubChannel in the request.")
		chimeSdkMessaging_getChannelMessageStatusCmd.MarkFlagRequired("channel-arn")
		chimeSdkMessaging_getChannelMessageStatusCmd.MarkFlagRequired("chime-bearer")
		chimeSdkMessaging_getChannelMessageStatusCmd.MarkFlagRequired("message-id")
	})
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_getChannelMessageStatusCmd)
}
