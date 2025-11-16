package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_listChannelFlowsCmd = &cobra.Command{
	Use:   "list-channel-flows",
	Short: "Returns a paginated lists of all the channel flows created under a single Chime.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_listChannelFlowsCmd).Standalone()

	chimeSdkMessaging_listChannelFlowsCmd.Flags().String("app-instance-arn", "", "The ARN of the app instance.")
	chimeSdkMessaging_listChannelFlowsCmd.Flags().String("max-results", "", "The maximum number of channel flows that you want to return.")
	chimeSdkMessaging_listChannelFlowsCmd.Flags().String("next-token", "", "The token passed by previous API calls until all requested channel flows are returned.")
	chimeSdkMessaging_listChannelFlowsCmd.MarkFlagRequired("app-instance-arn")
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_listChannelFlowsCmd)
}
