package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_listSubChannelsCmd = &cobra.Command{
	Use:   "list-sub-channels",
	Short: "Lists all the SubChannels in an elastic channel when given a channel ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_listSubChannelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMessaging_listSubChannelsCmd).Standalone()

		chimeSdkMessaging_listSubChannelsCmd.Flags().String("channel-arn", "", "The ARN of elastic channel.")
		chimeSdkMessaging_listSubChannelsCmd.Flags().String("chime-bearer", "", "The `AppInstanceUserArn` of the user making the API call.")
		chimeSdkMessaging_listSubChannelsCmd.Flags().String("max-results", "", "The maximum number of sub-channels that you want to return.")
		chimeSdkMessaging_listSubChannelsCmd.Flags().String("next-token", "", "The token passed by previous API calls until all requested sub-channels are returned.")
		chimeSdkMessaging_listSubChannelsCmd.MarkFlagRequired("channel-arn")
		chimeSdkMessaging_listSubChannelsCmd.MarkFlagRequired("chime-bearer")
	})
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_listSubChannelsCmd)
}
