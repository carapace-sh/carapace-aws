package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_listChannelsAssociatedWithChannelFlowCmd = &cobra.Command{
	Use:   "list-channels-associated-with-channel-flow",
	Short: "Lists all channels associated with a specified channel flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_listChannelsAssociatedWithChannelFlowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMessaging_listChannelsAssociatedWithChannelFlowCmd).Standalone()

		chimeSdkMessaging_listChannelsAssociatedWithChannelFlowCmd.Flags().String("channel-flow-arn", "", "The ARN of the channel flow.")
		chimeSdkMessaging_listChannelsAssociatedWithChannelFlowCmd.Flags().String("max-results", "", "The maximum number of channels that you want to return.")
		chimeSdkMessaging_listChannelsAssociatedWithChannelFlowCmd.Flags().String("next-token", "", "The token passed by previous API calls until all requested channels are returned.")
		chimeSdkMessaging_listChannelsAssociatedWithChannelFlowCmd.MarkFlagRequired("channel-flow-arn")
	})
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_listChannelsAssociatedWithChannelFlowCmd)
}
