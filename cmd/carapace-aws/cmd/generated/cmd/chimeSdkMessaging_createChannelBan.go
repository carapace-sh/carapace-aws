package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_createChannelBanCmd = &cobra.Command{
	Use:   "create-channel-ban",
	Short: "Permanently bans a member from a channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_createChannelBanCmd).Standalone()

	chimeSdkMessaging_createChannelBanCmd.Flags().String("channel-arn", "", "The ARN of the ban request.")
	chimeSdkMessaging_createChannelBanCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
	chimeSdkMessaging_createChannelBanCmd.Flags().String("member-arn", "", "The `AppInstanceUserArn` of the member being banned.")
	chimeSdkMessaging_createChannelBanCmd.MarkFlagRequired("channel-arn")
	chimeSdkMessaging_createChannelBanCmd.MarkFlagRequired("chime-bearer")
	chimeSdkMessaging_createChannelBanCmd.MarkFlagRequired("member-arn")
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_createChannelBanCmd)
}
