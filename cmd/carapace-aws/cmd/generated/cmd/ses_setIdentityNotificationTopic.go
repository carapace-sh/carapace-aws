package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_setIdentityNotificationTopicCmd = &cobra.Command{
	Use:   "set-identity-notification-topic",
	Short: "Sets an Amazon Simple Notification Service (Amazon SNS) topic to use when delivering notifications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_setIdentityNotificationTopicCmd).Standalone()

	ses_setIdentityNotificationTopicCmd.Flags().String("identity", "", "The identity (email address or domain) for the Amazon SNS topic.")
	ses_setIdentityNotificationTopicCmd.Flags().String("notification-type", "", "The type of notifications that are published to the specified Amazon SNS topic.")
	ses_setIdentityNotificationTopicCmd.Flags().String("sns-topic", "", "The Amazon Resource Name (ARN) of the Amazon SNS topic.")
	ses_setIdentityNotificationTopicCmd.MarkFlagRequired("identity")
	ses_setIdentityNotificationTopicCmd.MarkFlagRequired("notification-type")
	sesCmd.AddCommand(ses_setIdentityNotificationTopicCmd)
}
