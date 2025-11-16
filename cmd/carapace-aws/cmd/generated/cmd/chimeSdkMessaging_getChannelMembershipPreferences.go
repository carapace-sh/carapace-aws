package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_getChannelMembershipPreferencesCmd = &cobra.Command{
	Use:   "get-channel-membership-preferences",
	Short: "Gets the membership preferences of an `AppInstanceUser` or `AppInstanceBot` for the specified channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_getChannelMembershipPreferencesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMessaging_getChannelMembershipPreferencesCmd).Standalone()

		chimeSdkMessaging_getChannelMembershipPreferencesCmd.Flags().String("channel-arn", "", "The ARN of the channel.")
		chimeSdkMessaging_getChannelMembershipPreferencesCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
		chimeSdkMessaging_getChannelMembershipPreferencesCmd.Flags().String("member-arn", "", "The `AppInstanceUserArn` of the member retrieving the preferences.")
		chimeSdkMessaging_getChannelMembershipPreferencesCmd.MarkFlagRequired("channel-arn")
		chimeSdkMessaging_getChannelMembershipPreferencesCmd.MarkFlagRequired("chime-bearer")
		chimeSdkMessaging_getChannelMembershipPreferencesCmd.MarkFlagRequired("member-arn")
	})
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_getChannelMembershipPreferencesCmd)
}
