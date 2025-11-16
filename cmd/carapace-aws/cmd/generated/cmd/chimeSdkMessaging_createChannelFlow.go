package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_createChannelFlowCmd = &cobra.Command{
	Use:   "create-channel-flow",
	Short: "Creates a channel flow, a container for processors.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_createChannelFlowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMessaging_createChannelFlowCmd).Standalone()

		chimeSdkMessaging_createChannelFlowCmd.Flags().String("app-instance-arn", "", "The ARN of the channel flow request.")
		chimeSdkMessaging_createChannelFlowCmd.Flags().String("client-request-token", "", "The client token for the request.")
		chimeSdkMessaging_createChannelFlowCmd.Flags().String("name", "", "The name of the channel flow.")
		chimeSdkMessaging_createChannelFlowCmd.Flags().String("processors", "", "Information about the processor Lambda functions.")
		chimeSdkMessaging_createChannelFlowCmd.Flags().String("tags", "", "The tags for the creation request.")
		chimeSdkMessaging_createChannelFlowCmd.MarkFlagRequired("app-instance-arn")
		chimeSdkMessaging_createChannelFlowCmd.MarkFlagRequired("client-request-token")
		chimeSdkMessaging_createChannelFlowCmd.MarkFlagRequired("name")
		chimeSdkMessaging_createChannelFlowCmd.MarkFlagRequired("processors")
	})
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_createChannelFlowCmd)
}
