package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_createChannelCmd = &cobra.Command{
	Use:   "create-channel",
	Short: "Creates a channel to which you can add users and send messages.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_createChannelCmd).Standalone()

	chimeSdkMessaging_createChannelCmd.Flags().String("app-instance-arn", "", "The ARN of the channel request.")
	chimeSdkMessaging_createChannelCmd.Flags().String("channel-id", "", "An ID for the channel being created.")
	chimeSdkMessaging_createChannelCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
	chimeSdkMessaging_createChannelCmd.Flags().String("client-request-token", "", "The client token for the request.")
	chimeSdkMessaging_createChannelCmd.Flags().String("elastic-channel-configuration", "", "The attributes required to configure and create an elastic channel.")
	chimeSdkMessaging_createChannelCmd.Flags().String("expiration-settings", "", "Settings that control the interval after which the channel is automatically deleted.")
	chimeSdkMessaging_createChannelCmd.Flags().String("member-arns", "", "The ARNs of the channel members in the request.")
	chimeSdkMessaging_createChannelCmd.Flags().String("metadata", "", "The metadata of the creation request.")
	chimeSdkMessaging_createChannelCmd.Flags().String("mode", "", "The channel mode: `UNRESTRICTED` or `RESTRICTED`.")
	chimeSdkMessaging_createChannelCmd.Flags().String("moderator-arns", "", "The ARNs of the channel moderators in the request.")
	chimeSdkMessaging_createChannelCmd.Flags().String("name", "", "The name of the channel.")
	chimeSdkMessaging_createChannelCmd.Flags().String("privacy", "", "The channel's privacy level: `PUBLIC` or `PRIVATE`.")
	chimeSdkMessaging_createChannelCmd.Flags().String("tags", "", "The tags for the creation request.")
	chimeSdkMessaging_createChannelCmd.MarkFlagRequired("app-instance-arn")
	chimeSdkMessaging_createChannelCmd.MarkFlagRequired("chime-bearer")
	chimeSdkMessaging_createChannelCmd.MarkFlagRequired("client-request-token")
	chimeSdkMessaging_createChannelCmd.MarkFlagRequired("name")
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_createChannelCmd)
}
