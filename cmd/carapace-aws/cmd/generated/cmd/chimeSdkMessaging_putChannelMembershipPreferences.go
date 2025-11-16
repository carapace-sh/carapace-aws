package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_putChannelMembershipPreferencesCmd = &cobra.Command{
	Use:   "put-channel-membership-preferences",
	Short: "Sets the membership preferences of an `AppInstanceUser` or `AppInstanceBot` for the specified channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_putChannelMembershipPreferencesCmd).Standalone()

	chimeSdkMessaging_putChannelMembershipPreferencesCmd.Flags().String("channel-arn", "", "The ARN of the channel.")
	chimeSdkMessaging_putChannelMembershipPreferencesCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
	chimeSdkMessaging_putChannelMembershipPreferencesCmd.Flags().String("member-arn", "", "The ARN of the member setting the preferences.")
	chimeSdkMessaging_putChannelMembershipPreferencesCmd.Flags().String("preferences", "", "The channel membership preferences of an `AppInstanceUser` .")
	chimeSdkMessaging_putChannelMembershipPreferencesCmd.MarkFlagRequired("channel-arn")
	chimeSdkMessaging_putChannelMembershipPreferencesCmd.MarkFlagRequired("chime-bearer")
	chimeSdkMessaging_putChannelMembershipPreferencesCmd.MarkFlagRequired("member-arn")
	chimeSdkMessaging_putChannelMembershipPreferencesCmd.MarkFlagRequired("preferences")
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_putChannelMembershipPreferencesCmd)
}
