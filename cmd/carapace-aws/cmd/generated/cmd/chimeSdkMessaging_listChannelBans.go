package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_listChannelBansCmd = &cobra.Command{
	Use:   "list-channel-bans",
	Short: "Lists all the users and bots banned from a particular channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_listChannelBansCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMessaging_listChannelBansCmd).Standalone()

		chimeSdkMessaging_listChannelBansCmd.Flags().String("channel-arn", "", "The ARN of the channel.")
		chimeSdkMessaging_listChannelBansCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
		chimeSdkMessaging_listChannelBansCmd.Flags().String("max-results", "", "The maximum number of bans that you want returned.")
		chimeSdkMessaging_listChannelBansCmd.Flags().String("next-token", "", "The token passed by previous API calls until all requested bans are returned.")
		chimeSdkMessaging_listChannelBansCmd.MarkFlagRequired("channel-arn")
		chimeSdkMessaging_listChannelBansCmd.MarkFlagRequired("chime-bearer")
	})
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_listChannelBansCmd)
}
