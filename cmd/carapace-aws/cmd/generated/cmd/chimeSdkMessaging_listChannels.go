package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_listChannelsCmd = &cobra.Command{
	Use:   "list-channels",
	Short: "Lists all Channels created under a single Chime App as a paginated list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_listChannelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMessaging_listChannelsCmd).Standalone()

		chimeSdkMessaging_listChannelsCmd.Flags().String("app-instance-arn", "", "The ARN of the `AppInstance`.")
		chimeSdkMessaging_listChannelsCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
		chimeSdkMessaging_listChannelsCmd.Flags().String("max-results", "", "The maximum number of channels that you want to return.")
		chimeSdkMessaging_listChannelsCmd.Flags().String("next-token", "", "The token passed by previous API calls until all requested channels are returned.")
		chimeSdkMessaging_listChannelsCmd.Flags().String("privacy", "", "The privacy setting.")
		chimeSdkMessaging_listChannelsCmd.MarkFlagRequired("app-instance-arn")
		chimeSdkMessaging_listChannelsCmd.MarkFlagRequired("chime-bearer")
	})
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_listChannelsCmd)
}
