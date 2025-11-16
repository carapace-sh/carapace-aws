package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_sendChannelMessageCmd = &cobra.Command{
	Use:   "send-channel-message",
	Short: "Sends a message to a particular channel that the member is a part of.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_sendChannelMessageCmd).Standalone()

	chimeSdkMessaging_sendChannelMessageCmd.Flags().String("channel-arn", "", "The ARN of the channel.")
	chimeSdkMessaging_sendChannelMessageCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
	chimeSdkMessaging_sendChannelMessageCmd.Flags().String("client-request-token", "", "The `Idempotency` token for each client request.")
	chimeSdkMessaging_sendChannelMessageCmd.Flags().String("content", "", "The content of the channel message.")
	chimeSdkMessaging_sendChannelMessageCmd.Flags().String("content-type", "", "The content type of the channel message.")
	chimeSdkMessaging_sendChannelMessageCmd.Flags().String("message-attributes", "", "The attributes for the message, used for message filtering along with a `FilterRule` defined in the `PushNotificationPreferences`.")
	chimeSdkMessaging_sendChannelMessageCmd.Flags().String("metadata", "", "The optional metadata for each message.")
	chimeSdkMessaging_sendChannelMessageCmd.Flags().String("persistence", "", "Boolean that controls whether the message is persisted on the back end.")
	chimeSdkMessaging_sendChannelMessageCmd.Flags().String("push-notification", "", "The push notification configuration of the message.")
	chimeSdkMessaging_sendChannelMessageCmd.Flags().String("sub-channel-id", "", "The ID of the SubChannel in the request.")
	chimeSdkMessaging_sendChannelMessageCmd.Flags().String("target", "", "The target of a message.")
	chimeSdkMessaging_sendChannelMessageCmd.Flags().String("type", "", "The type of message, `STANDARD` or `CONTROL`.")
	chimeSdkMessaging_sendChannelMessageCmd.MarkFlagRequired("channel-arn")
	chimeSdkMessaging_sendChannelMessageCmd.MarkFlagRequired("chime-bearer")
	chimeSdkMessaging_sendChannelMessageCmd.MarkFlagRequired("client-request-token")
	chimeSdkMessaging_sendChannelMessageCmd.MarkFlagRequired("content")
	chimeSdkMessaging_sendChannelMessageCmd.MarkFlagRequired("persistence")
	chimeSdkMessaging_sendChannelMessageCmd.MarkFlagRequired("type")
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_sendChannelMessageCmd)
}
