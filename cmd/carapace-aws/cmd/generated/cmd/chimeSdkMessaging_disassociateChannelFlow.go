package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_disassociateChannelFlowCmd = &cobra.Command{
	Use:   "disassociate-channel-flow",
	Short: "Disassociates a channel flow from all its channels.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_disassociateChannelFlowCmd).Standalone()

	chimeSdkMessaging_disassociateChannelFlowCmd.Flags().String("channel-arn", "", "The ARN of the channel.")
	chimeSdkMessaging_disassociateChannelFlowCmd.Flags().String("channel-flow-arn", "", "The ARN of the channel flow.")
	chimeSdkMessaging_disassociateChannelFlowCmd.Flags().String("chime-bearer", "", "The `AppInstanceUserArn` of the user making the API call.")
	chimeSdkMessaging_disassociateChannelFlowCmd.MarkFlagRequired("channel-arn")
	chimeSdkMessaging_disassociateChannelFlowCmd.MarkFlagRequired("channel-flow-arn")
	chimeSdkMessaging_disassociateChannelFlowCmd.MarkFlagRequired("chime-bearer")
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_disassociateChannelFlowCmd)
}
