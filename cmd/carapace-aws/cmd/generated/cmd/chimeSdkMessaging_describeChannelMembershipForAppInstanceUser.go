package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_describeChannelMembershipForAppInstanceUserCmd = &cobra.Command{
	Use:   "describe-channel-membership-for-app-instance-user",
	Short: "Returns the details of a channel based on the membership of the specified `AppInstanceUser` or `AppInstanceBot`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_describeChannelMembershipForAppInstanceUserCmd).Standalone()

	chimeSdkMessaging_describeChannelMembershipForAppInstanceUserCmd.Flags().String("app-instance-user-arn", "", "The ARN of the user or bot in a channel.")
	chimeSdkMessaging_describeChannelMembershipForAppInstanceUserCmd.Flags().String("channel-arn", "", "The ARN of the channel to which the user belongs.")
	chimeSdkMessaging_describeChannelMembershipForAppInstanceUserCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
	chimeSdkMessaging_describeChannelMembershipForAppInstanceUserCmd.MarkFlagRequired("app-instance-user-arn")
	chimeSdkMessaging_describeChannelMembershipForAppInstanceUserCmd.MarkFlagRequired("channel-arn")
	chimeSdkMessaging_describeChannelMembershipForAppInstanceUserCmd.MarkFlagRequired("chime-bearer")
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_describeChannelMembershipForAppInstanceUserCmd)
}
