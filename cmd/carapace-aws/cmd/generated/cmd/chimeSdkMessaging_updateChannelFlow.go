package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_updateChannelFlowCmd = &cobra.Command{
	Use:   "update-channel-flow",
	Short: "Updates channel flow attributes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_updateChannelFlowCmd).Standalone()

	chimeSdkMessaging_updateChannelFlowCmd.Flags().String("channel-flow-arn", "", "The ARN of the channel flow.")
	chimeSdkMessaging_updateChannelFlowCmd.Flags().String("name", "", "The name of the channel flow.")
	chimeSdkMessaging_updateChannelFlowCmd.Flags().String("processors", "", "Information about the processor Lambda functions")
	chimeSdkMessaging_updateChannelFlowCmd.MarkFlagRequired("channel-flow-arn")
	chimeSdkMessaging_updateChannelFlowCmd.MarkFlagRequired("name")
	chimeSdkMessaging_updateChannelFlowCmd.MarkFlagRequired("processors")
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_updateChannelFlowCmd)
}
