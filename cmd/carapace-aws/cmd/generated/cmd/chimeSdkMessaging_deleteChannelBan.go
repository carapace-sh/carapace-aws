package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_deleteChannelBanCmd = &cobra.Command{
	Use:   "delete-channel-ban",
	Short: "Removes a member from a channel's ban list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_deleteChannelBanCmd).Standalone()

	chimeSdkMessaging_deleteChannelBanCmd.Flags().String("channel-arn", "", "The ARN of the channel from which the `AppInstanceUser` was banned.")
	chimeSdkMessaging_deleteChannelBanCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
	chimeSdkMessaging_deleteChannelBanCmd.Flags().String("member-arn", "", "The ARN of the `AppInstanceUser` that you want to reinstate.")
	chimeSdkMessaging_deleteChannelBanCmd.MarkFlagRequired("channel-arn")
	chimeSdkMessaging_deleteChannelBanCmd.MarkFlagRequired("chime-bearer")
	chimeSdkMessaging_deleteChannelBanCmd.MarkFlagRequired("member-arn")
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_deleteChannelBanCmd)
}
