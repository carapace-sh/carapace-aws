package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_updateChannelCmd = &cobra.Command{
	Use:   "update-channel",
	Short: "Update a channel's attributes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_updateChannelCmd).Standalone()

	chimeSdkMessaging_updateChannelCmd.Flags().String("channel-arn", "", "The ARN of the channel.")
	chimeSdkMessaging_updateChannelCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
	chimeSdkMessaging_updateChannelCmd.Flags().String("metadata", "", "The metadata for the update request.")
	chimeSdkMessaging_updateChannelCmd.Flags().String("mode", "", "The mode of the update request.")
	chimeSdkMessaging_updateChannelCmd.Flags().String("name", "", "The name of the channel.")
	chimeSdkMessaging_updateChannelCmd.MarkFlagRequired("channel-arn")
	chimeSdkMessaging_updateChannelCmd.MarkFlagRequired("chime-bearer")
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_updateChannelCmd)
}
