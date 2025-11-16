package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_describeChannelModeratedByAppInstanceUserCmd = &cobra.Command{
	Use:   "describe-channel-moderated-by-app-instance-user",
	Short: "Returns the full details of a channel moderated by the specified `AppInstanceUser` or `AppInstanceBot`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_describeChannelModeratedByAppInstanceUserCmd).Standalone()

	chimeSdkMessaging_describeChannelModeratedByAppInstanceUserCmd.Flags().String("app-instance-user-arn", "", "The ARN of the user or bot in the moderated channel.")
	chimeSdkMessaging_describeChannelModeratedByAppInstanceUserCmd.Flags().String("channel-arn", "", "The ARN of the moderated channel.")
	chimeSdkMessaging_describeChannelModeratedByAppInstanceUserCmd.Flags().String("chime-bearer", "", "The ARN of the `AppInstanceUser` or `AppInstanceBot` that makes the API call.")
	chimeSdkMessaging_describeChannelModeratedByAppInstanceUserCmd.MarkFlagRequired("app-instance-user-arn")
	chimeSdkMessaging_describeChannelModeratedByAppInstanceUserCmd.MarkFlagRequired("channel-arn")
	chimeSdkMessaging_describeChannelModeratedByAppInstanceUserCmd.MarkFlagRequired("chime-bearer")
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_describeChannelModeratedByAppInstanceUserCmd)
}
