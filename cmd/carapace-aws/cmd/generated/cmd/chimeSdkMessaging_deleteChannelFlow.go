package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_deleteChannelFlowCmd = &cobra.Command{
	Use:   "delete-channel-flow",
	Short: "Deletes a channel flow, an irreversible process.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_deleteChannelFlowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMessaging_deleteChannelFlowCmd).Standalone()

		chimeSdkMessaging_deleteChannelFlowCmd.Flags().String("channel-flow-arn", "", "The ARN of the channel flow.")
		chimeSdkMessaging_deleteChannelFlowCmd.MarkFlagRequired("channel-flow-arn")
	})
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_deleteChannelFlowCmd)
}
