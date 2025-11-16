package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_listChannelMembershipsCmd = &cobra.Command{
	Use:   "list-channel-memberships",
	Short: "Lists all channel memberships in a channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_listChannelMembershipsCmd).Standalone()

	chimeSdkMessaging_listChannelMembershipsCmd.Flags().String("channel-arn", "", "The maximum number of channel memberships that you want returned.")
	chimeSdkMessaging_listChannelMembershipsCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
	chimeSdkMessaging_listChannelMembershipsCmd.Flags().String("max-results", "", "The maximum number of channel memberships that you want returned.")
	chimeSdkMessaging_listChannelMembershipsCmd.Flags().String("next-token", "", "The token passed by previous API calls until all requested channel memberships are returned.")
	chimeSdkMessaging_listChannelMembershipsCmd.Flags().String("sub-channel-id", "", "The ID of the SubChannel in the request.")
	chimeSdkMessaging_listChannelMembershipsCmd.Flags().String("type", "", "The membership type of a user, `DEFAULT` or `HIDDEN`.")
	chimeSdkMessaging_listChannelMembershipsCmd.MarkFlagRequired("channel-arn")
	chimeSdkMessaging_listChannelMembershipsCmd.MarkFlagRequired("chime-bearer")
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_listChannelMembershipsCmd)
}
