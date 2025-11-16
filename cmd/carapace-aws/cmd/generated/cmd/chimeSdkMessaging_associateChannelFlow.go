package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_associateChannelFlowCmd = &cobra.Command{
	Use:   "associate-channel-flow",
	Short: "Associates a channel flow with a channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_associateChannelFlowCmd).Standalone()

	chimeSdkMessaging_associateChannelFlowCmd.Flags().String("channel-arn", "", "The ARN of the channel.")
	chimeSdkMessaging_associateChannelFlowCmd.Flags().String("channel-flow-arn", "", "The ARN of the channel flow.")
	chimeSdkMessaging_associateChannelFlowCmd.Flags().String("chime-bearer", "", "The `AppInstanceUserArn` of the user making the API call.")
	chimeSdkMessaging_associateChannelFlowCmd.MarkFlagRequired("channel-arn")
	chimeSdkMessaging_associateChannelFlowCmd.MarkFlagRequired("channel-flow-arn")
	chimeSdkMessaging_associateChannelFlowCmd.MarkFlagRequired("chime-bearer")
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_associateChannelFlowCmd)
}
