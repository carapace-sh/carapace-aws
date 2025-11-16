package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_deleteChannelMembershipCmd = &cobra.Command{
	Use:   "delete-channel-membership",
	Short: "Removes a member from a channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_deleteChannelMembershipCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMessaging_deleteChannelMembershipCmd).Standalone()

		chimeSdkMessaging_deleteChannelMembershipCmd.Flags().String("channel-arn", "", "The ARN of the channel from which you want to remove the user.")
		chimeSdkMessaging_deleteChannelMembershipCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
		chimeSdkMessaging_deleteChannelMembershipCmd.Flags().String("member-arn", "", "The `AppInstanceUserArn` of the member that you're removing from the channel.")
		chimeSdkMessaging_deleteChannelMembershipCmd.Flags().String("sub-channel-id", "", "The ID of the SubChannel in the request.")
		chimeSdkMessaging_deleteChannelMembershipCmd.MarkFlagRequired("channel-arn")
		chimeSdkMessaging_deleteChannelMembershipCmd.MarkFlagRequired("chime-bearer")
		chimeSdkMessaging_deleteChannelMembershipCmd.MarkFlagRequired("member-arn")
	})
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_deleteChannelMembershipCmd)
}
