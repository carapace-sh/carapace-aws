package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_createChannelMembershipCmd = &cobra.Command{
	Use:   "create-channel-membership",
	Short: "Adds a member to a channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_createChannelMembershipCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMessaging_createChannelMembershipCmd).Standalone()

		chimeSdkMessaging_createChannelMembershipCmd.Flags().String("channel-arn", "", "The ARN of the channel to which you're adding users.")
		chimeSdkMessaging_createChannelMembershipCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
		chimeSdkMessaging_createChannelMembershipCmd.Flags().String("member-arn", "", "The `AppInstanceUserArn` of the member you want to add to the channel.")
		chimeSdkMessaging_createChannelMembershipCmd.Flags().String("sub-channel-id", "", "The ID of the SubChannel in the request.")
		chimeSdkMessaging_createChannelMembershipCmd.Flags().String("type", "", "The membership type of a user, `DEFAULT` or `HIDDEN`.")
		chimeSdkMessaging_createChannelMembershipCmd.MarkFlagRequired("channel-arn")
		chimeSdkMessaging_createChannelMembershipCmd.MarkFlagRequired("chime-bearer")
		chimeSdkMessaging_createChannelMembershipCmd.MarkFlagRequired("member-arn")
		chimeSdkMessaging_createChannelMembershipCmd.MarkFlagRequired("type")
	})
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_createChannelMembershipCmd)
}
