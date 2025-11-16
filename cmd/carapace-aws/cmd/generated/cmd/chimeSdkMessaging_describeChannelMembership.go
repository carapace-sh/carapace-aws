package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_describeChannelMembershipCmd = &cobra.Command{
	Use:   "describe-channel-membership",
	Short: "Returns the full details of a user's channel membership.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_describeChannelMembershipCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMessaging_describeChannelMembershipCmd).Standalone()

		chimeSdkMessaging_describeChannelMembershipCmd.Flags().String("channel-arn", "", "The ARN of the channel.")
		chimeSdkMessaging_describeChannelMembershipCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
		chimeSdkMessaging_describeChannelMembershipCmd.Flags().String("member-arn", "", "The `AppInstanceUserArn` of the member.")
		chimeSdkMessaging_describeChannelMembershipCmd.Flags().String("sub-channel-id", "", "The ID of the SubChannel in the request.")
		chimeSdkMessaging_describeChannelMembershipCmd.MarkFlagRequired("channel-arn")
		chimeSdkMessaging_describeChannelMembershipCmd.MarkFlagRequired("chime-bearer")
		chimeSdkMessaging_describeChannelMembershipCmd.MarkFlagRequired("member-arn")
	})
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_describeChannelMembershipCmd)
}
