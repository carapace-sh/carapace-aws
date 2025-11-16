package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_searchChannelsCmd = &cobra.Command{
	Use:   "search-channels",
	Short: "Allows the `ChimeBearer` to search channels by channel members.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_searchChannelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMessaging_searchChannelsCmd).Standalone()

		chimeSdkMessaging_searchChannelsCmd.Flags().String("chime-bearer", "", "The `AppInstanceUserArn` of the user making the API call.")
		chimeSdkMessaging_searchChannelsCmd.Flags().String("fields", "", "A list of the `Field` objects in the channel being searched.")
		chimeSdkMessaging_searchChannelsCmd.Flags().String("max-results", "", "The maximum number of channels that you want returned.")
		chimeSdkMessaging_searchChannelsCmd.Flags().String("next-token", "", "The token returned from previous API requests until the number of channels is reached.")
		chimeSdkMessaging_searchChannelsCmd.MarkFlagRequired("fields")
	})
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_searchChannelsCmd)
}
