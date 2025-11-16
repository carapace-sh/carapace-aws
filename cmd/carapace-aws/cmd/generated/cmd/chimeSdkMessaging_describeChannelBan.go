package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_describeChannelBanCmd = &cobra.Command{
	Use:   "describe-channel-ban",
	Short: "Returns the full details of a channel ban.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_describeChannelBanCmd).Standalone()

	chimeSdkMessaging_describeChannelBanCmd.Flags().String("channel-arn", "", "The ARN of the channel from which the user is banned.")
	chimeSdkMessaging_describeChannelBanCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
	chimeSdkMessaging_describeChannelBanCmd.Flags().String("member-arn", "", "The `AppInstanceUserArn` of the member being banned.")
	chimeSdkMessaging_describeChannelBanCmd.MarkFlagRequired("channel-arn")
	chimeSdkMessaging_describeChannelBanCmd.MarkFlagRequired("chime-bearer")
	chimeSdkMessaging_describeChannelBanCmd.MarkFlagRequired("member-arn")
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_describeChannelBanCmd)
}
