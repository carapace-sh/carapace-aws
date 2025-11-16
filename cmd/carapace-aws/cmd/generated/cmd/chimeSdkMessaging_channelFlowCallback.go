package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_channelFlowCallbackCmd = &cobra.Command{
	Use:   "channel-flow-callback",
	Short: "Calls back Amazon Chime SDK messaging with a processing response message.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_channelFlowCallbackCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMessaging_channelFlowCallbackCmd).Standalone()

		chimeSdkMessaging_channelFlowCallbackCmd.Flags().String("callback-id", "", "The identifier passed to the processor by the service when invoked.")
		chimeSdkMessaging_channelFlowCallbackCmd.Flags().String("channel-arn", "", "The ARN of the channel.")
		chimeSdkMessaging_channelFlowCallbackCmd.Flags().String("channel-message", "", "Stores information about the processed message.")
		chimeSdkMessaging_channelFlowCallbackCmd.Flags().String("delete-resource", "", "When a processor determines that a message needs to be `DENIED`, pass this parameter with a value of true.")
		chimeSdkMessaging_channelFlowCallbackCmd.MarkFlagRequired("callback-id")
		chimeSdkMessaging_channelFlowCallbackCmd.MarkFlagRequired("channel-arn")
		chimeSdkMessaging_channelFlowCallbackCmd.MarkFlagRequired("channel-message")
	})
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_channelFlowCallbackCmd)
}
