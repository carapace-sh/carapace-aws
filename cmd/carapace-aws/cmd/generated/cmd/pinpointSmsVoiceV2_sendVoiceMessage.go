package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_sendVoiceMessageCmd = &cobra.Command{
	Use:   "send-voice-message",
	Short: "Allows you to send a request that sends a voice message.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_sendVoiceMessageCmd).Standalone()

	pinpointSmsVoiceV2_sendVoiceMessageCmd.Flags().String("configuration-set-name", "", "The name of the configuration set to use.")
	pinpointSmsVoiceV2_sendVoiceMessageCmd.Flags().String("context", "", "You can specify custom data in this field.")
	pinpointSmsVoiceV2_sendVoiceMessageCmd.Flags().String("destination-phone-number", "", "The destination phone number in E.164 format.")
	pinpointSmsVoiceV2_sendVoiceMessageCmd.Flags().String("dry-run", "", "When set to true, the message is checked and validated, but isn't sent to the end recipient.")
	pinpointSmsVoiceV2_sendVoiceMessageCmd.Flags().String("max-price-per-minute", "", "The maximum amount to spend per voice message, in US dollars.")
	pinpointSmsVoiceV2_sendVoiceMessageCmd.Flags().String("message-body", "", "The text to convert to a voice message.")
	pinpointSmsVoiceV2_sendVoiceMessageCmd.Flags().String("message-body-text-type", "", "Specifies if the MessageBody field contains text or [speech synthesis markup language (SSML)](https://docs.aws.amazon.com/polly/latest/dg/what-is.html).")
	pinpointSmsVoiceV2_sendVoiceMessageCmd.Flags().Bool("message-feedback-enabled", false, "Set to true to enable message feedback for the message.")
	pinpointSmsVoiceV2_sendVoiceMessageCmd.Flags().Bool("no-message-feedback-enabled", false, "Set to true to enable message feedback for the message.")
	pinpointSmsVoiceV2_sendVoiceMessageCmd.Flags().String("origination-identity", "", "The origination identity to use for the voice call.")
	pinpointSmsVoiceV2_sendVoiceMessageCmd.Flags().String("protect-configuration-id", "", "The unique identifier for the protect configuration.")
	pinpointSmsVoiceV2_sendVoiceMessageCmd.Flags().String("time-to-live", "", "How long the voice message is valid for.")
	pinpointSmsVoiceV2_sendVoiceMessageCmd.Flags().String("voice-id", "", "The voice for the [Amazon Polly](https://docs.aws.amazon.com/polly/latest/dg/what-is.html) service to use.")
	pinpointSmsVoiceV2_sendVoiceMessageCmd.MarkFlagRequired("destination-phone-number")
	pinpointSmsVoiceV2_sendVoiceMessageCmd.Flag("no-message-feedback-enabled").Hidden = true
	pinpointSmsVoiceV2_sendVoiceMessageCmd.MarkFlagRequired("origination-identity")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_sendVoiceMessageCmd)
}
