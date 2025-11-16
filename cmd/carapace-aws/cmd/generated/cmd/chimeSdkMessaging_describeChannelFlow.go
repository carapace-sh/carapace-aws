package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_describeChannelFlowCmd = &cobra.Command{
	Use:   "describe-channel-flow",
	Short: "Returns the full details of a channel flow in an Amazon Chime `AppInstance`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_describeChannelFlowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMessaging_describeChannelFlowCmd).Standalone()

		chimeSdkMessaging_describeChannelFlowCmd.Flags().String("channel-flow-arn", "", "The ARN of the channel flow.")
		chimeSdkMessaging_describeChannelFlowCmd.MarkFlagRequired("channel-flow-arn")
	})
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_describeChannelFlowCmd)
}
