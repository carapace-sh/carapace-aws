package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_batchCreateChannelMembershipCmd = &cobra.Command{
	Use:   "batch-create-channel-membership",
	Short: "Adds a specified number of users and bots to a channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_batchCreateChannelMembershipCmd).Standalone()

	chimeSdkMessaging_batchCreateChannelMembershipCmd.Flags().String("channel-arn", "", "The ARN of the channel to which you're adding users or bots.")
	chimeSdkMessaging_batchCreateChannelMembershipCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
	chimeSdkMessaging_batchCreateChannelMembershipCmd.Flags().String("member-arns", "", "The ARNs of the members you want to add to the channel.")
	chimeSdkMessaging_batchCreateChannelMembershipCmd.Flags().String("sub-channel-id", "", "The ID of the SubChannel in the request.")
	chimeSdkMessaging_batchCreateChannelMembershipCmd.Flags().String("type", "", "The membership type of a user, `DEFAULT` or `HIDDEN`.")
	chimeSdkMessaging_batchCreateChannelMembershipCmd.MarkFlagRequired("channel-arn")
	chimeSdkMessaging_batchCreateChannelMembershipCmd.MarkFlagRequired("chime-bearer")
	chimeSdkMessaging_batchCreateChannelMembershipCmd.MarkFlagRequired("member-arns")
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_batchCreateChannelMembershipCmd)
}
