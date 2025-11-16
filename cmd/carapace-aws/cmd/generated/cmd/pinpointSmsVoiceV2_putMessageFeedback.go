package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_putMessageFeedbackCmd = &cobra.Command{
	Use:   "put-message-feedback",
	Short: "Set the MessageFeedbackStatus as `RECEIVED` or `FAILED` for the passed in MessageId.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_putMessageFeedbackCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_putMessageFeedbackCmd).Standalone()

		pinpointSmsVoiceV2_putMessageFeedbackCmd.Flags().String("message-feedback-status", "", "Set the message feedback to be either `RECEIVED` or `FAILED`.")
		pinpointSmsVoiceV2_putMessageFeedbackCmd.Flags().String("message-id", "", "The unique identifier for the message.")
		pinpointSmsVoiceV2_putMessageFeedbackCmd.MarkFlagRequired("message-feedback-status")
		pinpointSmsVoiceV2_putMessageFeedbackCmd.MarkFlagRequired("message-id")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_putMessageFeedbackCmd)
}
