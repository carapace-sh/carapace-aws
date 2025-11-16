package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_listChannelModeratorsCmd = &cobra.Command{
	Use:   "list-channel-moderators",
	Short: "Lists all the moderators for a channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_listChannelModeratorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMessaging_listChannelModeratorsCmd).Standalone()

		chimeSdkMessaging_listChannelModeratorsCmd.Flags().String("channel-arn", "", "The ARN of the channel.")
		chimeSdkMessaging_listChannelModeratorsCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
		chimeSdkMessaging_listChannelModeratorsCmd.Flags().String("max-results", "", "The maximum number of moderators that you want returned.")
		chimeSdkMessaging_listChannelModeratorsCmd.Flags().String("next-token", "", "The token passed by previous API calls until all requested moderators are returned.")
		chimeSdkMessaging_listChannelModeratorsCmd.MarkFlagRequired("channel-arn")
		chimeSdkMessaging_listChannelModeratorsCmd.MarkFlagRequired("chime-bearer")
	})
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_listChannelModeratorsCmd)
}
