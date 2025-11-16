package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_updateChannelReadMarkerCmd = &cobra.Command{
	Use:   "update-channel-read-marker",
	Short: "The details of the time when a user last read messages in a channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_updateChannelReadMarkerCmd).Standalone()

	chimeSdkMessaging_updateChannelReadMarkerCmd.Flags().String("channel-arn", "", "The ARN of the channel.")
	chimeSdkMessaging_updateChannelReadMarkerCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
	chimeSdkMessaging_updateChannelReadMarkerCmd.MarkFlagRequired("channel-arn")
	chimeSdkMessaging_updateChannelReadMarkerCmd.MarkFlagRequired("chime-bearer")
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_updateChannelReadMarkerCmd)
}
