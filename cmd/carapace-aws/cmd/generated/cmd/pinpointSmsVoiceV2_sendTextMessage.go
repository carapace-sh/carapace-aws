package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_sendTextMessageCmd = &cobra.Command{
	Use:   "send-text-message",
	Short: "Creates a new text message and sends it to a recipient's phone number.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_sendTextMessageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_sendTextMessageCmd).Standalone()

		pinpointSmsVoiceV2_sendTextMessageCmd.Flags().String("configuration-set-name", "", "The name of the configuration set to use.")
		pinpointSmsVoiceV2_sendTextMessageCmd.Flags().String("context", "", "You can specify custom data in this field.")
		pinpointSmsVoiceV2_sendTextMessageCmd.Flags().String("destination-country-parameters", "", "This field is used for any country-specific registration requirements.")
		pinpointSmsVoiceV2_sendTextMessageCmd.Flags().String("destination-phone-number", "", "The destination phone number in E.164 format.")
		pinpointSmsVoiceV2_sendTextMessageCmd.Flags().String("dry-run", "", "When set to true, the message is checked and validated, but isn't sent to the end recipient.")
		pinpointSmsVoiceV2_sendTextMessageCmd.Flags().String("keyword", "", "When you register a short code in the US, you must specify a program name.")
		pinpointSmsVoiceV2_sendTextMessageCmd.Flags().String("max-price", "", "The maximum amount that you want to spend, in US dollars, per each text message.")
		pinpointSmsVoiceV2_sendTextMessageCmd.Flags().String("message-body", "", "The body of the text message.")
		pinpointSmsVoiceV2_sendTextMessageCmd.Flags().Bool("message-feedback-enabled", false, "Set to true to enable message feedback for the message.")
		pinpointSmsVoiceV2_sendTextMessageCmd.Flags().String("message-type", "", "The type of message.")
		pinpointSmsVoiceV2_sendTextMessageCmd.Flags().Bool("no-message-feedback-enabled", false, "Set to true to enable message feedback for the message.")
		pinpointSmsVoiceV2_sendTextMessageCmd.Flags().String("origination-identity", "", "The origination identity of the message.")
		pinpointSmsVoiceV2_sendTextMessageCmd.Flags().String("protect-configuration-id", "", "The unique identifier for the protect configuration.")
		pinpointSmsVoiceV2_sendTextMessageCmd.Flags().String("time-to-live", "", "How long the text message is valid for, in seconds.")
		pinpointSmsVoiceV2_sendTextMessageCmd.MarkFlagRequired("destination-phone-number")
		pinpointSmsVoiceV2_sendTextMessageCmd.Flag("no-message-feedback-enabled").Hidden = true
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_sendTextMessageCmd)
}
