package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_putChannelExpirationSettingsCmd = &cobra.Command{
	Use:   "put-channel-expiration-settings",
	Short: "Sets the number of days before the channel is automatically deleted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_putChannelExpirationSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMessaging_putChannelExpirationSettingsCmd).Standalone()

		chimeSdkMessaging_putChannelExpirationSettingsCmd.Flags().String("channel-arn", "", "The ARN of the channel.")
		chimeSdkMessaging_putChannelExpirationSettingsCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
		chimeSdkMessaging_putChannelExpirationSettingsCmd.Flags().String("expiration-settings", "", "Settings that control the interval after which a channel is deleted.")
		chimeSdkMessaging_putChannelExpirationSettingsCmd.MarkFlagRequired("channel-arn")
	})
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_putChannelExpirationSettingsCmd)
}
