package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_describeChannelCmd = &cobra.Command{
	Use:   "describe-channel",
	Short: "Returns the full details of a channel in an Amazon Chime `AppInstance`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_describeChannelCmd).Standalone()

	chimeSdkMessaging_describeChannelCmd.Flags().String("channel-arn", "", "The ARN of the channel.")
	chimeSdkMessaging_describeChannelCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
	chimeSdkMessaging_describeChannelCmd.MarkFlagRequired("channel-arn")
	chimeSdkMessaging_describeChannelCmd.MarkFlagRequired("chime-bearer")
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_describeChannelCmd)
}
