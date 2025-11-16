package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_redactChannelMessageCmd = &cobra.Command{
	Use:   "redact-channel-message",
	Short: "Redacts message content and metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_redactChannelMessageCmd).Standalone()

	chimeSdkMessaging_redactChannelMessageCmd.Flags().String("channel-arn", "", "The ARN of the channel containing the messages that you want to redact.")
	chimeSdkMessaging_redactChannelMessageCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
	chimeSdkMessaging_redactChannelMessageCmd.Flags().String("message-id", "", "The ID of the message being redacted.")
	chimeSdkMessaging_redactChannelMessageCmd.Flags().String("sub-channel-id", "", "The ID of the SubChannel in the request.")
	chimeSdkMessaging_redactChannelMessageCmd.MarkFlagRequired("channel-arn")
	chimeSdkMessaging_redactChannelMessageCmd.MarkFlagRequired("chime-bearer")
	chimeSdkMessaging_redactChannelMessageCmd.MarkFlagRequired("message-id")
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_redactChannelMessageCmd)
}
